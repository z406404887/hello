package state

import (
	"network"
	"log"
	"github.com/gorilla/websocket"
	"github.com/golang/protobuf/proto"
)

type ConnAgent struct {
	gate *StateServer
	ws *network.WsConn
}

func NewConnAgent(gate*StateServer, ws *network.WsConn) *ConnAgent {
	a := &ConnAgent{
		gate:gate,
		ws:ws,
	}
	return a
}

func (agent *ConnAgent) Run(){
	agent.gate.regConn <- agent.ws
	defer func() {
		agent.gate.unRegConn <- agent.ws
	}()

	for {
		mt,msg,err := agent.ws.ReadMsg()
		if err != nil{
			log.Printf("read msg error, %v",err)
			break
		}

		if mt == websocket.TextMessage {
			log.Printf("invalid msg type, TextMessage")
		}

		agent.handleMsg(msg)
	}
}

func (agent *ConnAgent) handleMsg(data []byte){
	header := &network.CommonHeader{}
	header.Decode(data)
	log.Printf("receive msg. header=%+v",header)

	handleMsg(agent,header,data[network.COMMON_HEADER_LENGTH:])
}

func (agent *ConnAgent) sendMsgBack(header *network.CommonHeader, pb proto.Message)  {
	data,err := proto.Marshal(pb)
	if err != nil {
		log.Printf("marshal failed. %v",err)
		return
	}

	header.Len = uint16(len(data))
	msg := append(header.Encode(),data...)
	log.Printf("send back %v",msg)
	agent.ws.SendChan <- msg
}