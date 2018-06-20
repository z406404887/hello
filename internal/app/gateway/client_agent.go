package gateway

import (
	"hello/network"
	"log"
)

type ClientAgent struct {
	gate   *Gateway
	client *network.WsClient
}

func NewClientAgent(gate *Gateway, client *network.WsClient) *ClientAgent {
	return &ClientAgent{
		gate:   gate,
		client: client,
	}
}

func (agent *ClientAgent) Run() {
	log.Printf("register server client ")
	agent.gate.regSrv <- agent.client
	defer func() {
		log.Printf("unregister server client")
		agent.gate.unRegSrv <- agent.client
	}()

	for {
		select {
		case msg, ok := <-agent.client.RecvChan:
			//client has been closed
			if !ok {
				return
			}
			//log.Printf("recv server msg %v",msg)
			agent.handleMsg(msg)
		}
	}
}

func (agent *ClientAgent) handleMsg(data []byte) {
	header := &network.CommonHeader{}
	header.Decode(data)
	msg := &network.Message{
		SrcId: agent.client.Sid,
		Head:  header,
		Data:  data,
	}

	agent.gate.recvSrvMsg <- msg
}
