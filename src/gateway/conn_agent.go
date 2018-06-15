package gateway

import (
	"github.com/gorilla/websocket"
	"log"
	"network"
	"pb/pbgame"
)

type ConnAgent struct {
	gate *Gateway
	ws   *network.WsConn
}

func NewWsAgent(gate *Gateway, ws *network.WsConn) *ConnAgent {
	a := &ConnAgent{
		gate: gate,
		ws:   ws,
	}
	return a
}

func (agent *ConnAgent) Run() {
	agent.gate.regConn <- agent.ws
	defer func() {
		agent.gate.unRegConn <- agent.ws
	}()

	for {
		mt, msg, err := agent.ws.ReadMsg()
		if err != nil {
			log.Printf("read msg error, %v", err)
			break
		}

		if mt == websocket.TextMessage {
			log.Printf("invalid msg type, TextMessage")
		}

		agent.DispatchMsg(msg)
	}
}

func (agent *ConnAgent) DispatchMsg(data []byte) {
	header := &network.CommonHeader{}
	header.Decode(data)
	header.ClientId = uint32(agent.ws.Id)
	//log.Printf("receive client msg, header=%+v",header)
	msgData := data[network.COMMON_HEADER_LENGTH:]
	if ok := agent.handleMsg(header, msgData); ok {
		return
	}

	msg := &network.Message{
		SrcId: agent.ws.Id,
		Head:  header,
		Data:  msgData,
	}

	agent.gate.recvConnMsg <- msg
}

func (agent *ConnAgent) handleMsg(header *network.CommonHeader, data []byte) bool {
	switch header.MainType {
	case pbgame.MainAccount:
		return handleAccountMsg(agent, header, data)
	default:
		return false
	}
	return true
}

func (agent *ConnAgent) sendMsgBack(header *network.CommonHeader, data []byte) {
	header.Len = uint16(len(data))
	msg := append(header.Encode(), data...)
	//log.Printf("send back %v",msg)
	agent.ws.SendChan <- msg
}
