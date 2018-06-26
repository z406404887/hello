package gateway

import (
	"log"
	"starter-kit/internal/pkg/network"
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

	for msg := range agent.client.RecvChan {
		agent.handleMsg(msg)
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
