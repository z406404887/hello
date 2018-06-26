package gateway

import (
	"starter-kit/internal/pkg/network"
	"starter-kit/internal/pkg/pb/pbgame"
	"log"
	"time"
)

//handle conn msg
func handleConnMsg(gate *Gateway, msg *network.Message) {
	//log.Printf("handleConnMsg %+v",msg.Head)
	switch msg.Head.MainType {
	case pbgame.MainGame:
		HandleConnGameMsg(gate, msg)
	}
}

func HandleConnGameMsg(gate *Gateway, msg *network.Message) {
	switch msg.Head.SubType {
	case pbgame.SubEnterGameReq:
		HandleEnterGameReq(gate, msg)
	default:
		sendToGame(gate, msg)
	}
}

func sendToGame(gate *Gateway, msg *network.Message) {
	if msg.Head.MainType == pbgame.MainGame && msg.Head.SubType == pbgame.SubRollReq {
		gate.traceTime[msg.Head.ClientId] = time.Now().UnixNano()
	}
	//log.Printf("route msg %+v",msg.Head)
	game := gate.GetGameClientById(msg.Head.ClientId)
	if game == nil {
		log.Printf("game server not found.")
		return
	}

	data := append(msg.Head.Encode(), msg.Data...)
	game.SendChan <- data
}

func HandleEnterGameReq(gate *Gateway, msg *network.Message) {
	data := append(msg.Head.Encode(), msg.Data...)
	game := gate.GetGameClient()
	if game == nil {
		log.Printf("game is not connected.")
		return
	}
	game.SendChan <- data
}
