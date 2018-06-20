package gateway

import (
	"github.com/golang/protobuf/proto"
	"hello/internal/pkg/network"
	"hello/internal/pkg/pb/pbgame"
	"log"
	"time"
)

//handle conn msg
func handleServerMsg(gate *Gateway, msg *network.Message) {
	switch msg.Head.MainType {
	case pbgame.MainGame:
		HandleSrvGameMsg(gate, msg)
	}
}

func HandleSrvGameMsg(gate *Gateway, msg *network.Message) {
	switch msg.Head.SubType {
	case pbgame.SubEnterGameRsp:
		HandleEnterGameRsp(gate, msg)
	default:
		sendToClient(gate, msg)
	}
}

func sendToClient(gate *Gateway, msg *network.Message) {
	conn := gate.getConn(msg.Head.ClientId)

	if msg.Head.MainType == pbgame.MainGame && msg.Head.SubType == pbgame.SubRollRsp {
		if recv, ok := gate.traceTime[msg.Head.ClientId]; ok {
			now := time.Now().UnixNano()
			diff := (now - recv) / 1e6
			log.Printf("track %d %d %d %d", msg.Head.ClientId, recv/1e6, now/1e6, diff)
		}
	}

	if conn == nil {
		log.Printf("conn not found. drop msg. uid=%d", msg.Head.ClientId)
		return
	}

	conn.SendChan <- msg.Data
}

func HandleEnterGameRsp(gate *Gateway, msg *network.Message) {
	//enter game success
	//log.Printf("receive game response %+v",msg.Head)
	conn := gate.getConn(msg.Head.ClientId)
	if conn == nil {
		log.Printf("conn not found, id=%d", msg.Head.ClientId)
		return
	}

	rsp := &pbgame.EnterGameResponse{}
	err := proto.Unmarshal(msg.Data[network.COMMON_HEADER_LENGTH:], rsp)
	if err != nil {
		log.Printf("unmarshal EnterGameResponse failed. %v", err)
		return
	}

	if msg.Head.Result == 0 {
		gate.delConn(msg.Head.ClientId)
		conn.Id = rsp.Uid
		gate.addConn(conn)
	}
	gate.uidSrvMap[rsp.Uid] = msg.SrcId
	msg.Head.ClientId = rsp.Uid

	data := append(msg.Head.Encode(), msg.Data[network.COMMON_HEADER_LENGTH:]...)
	conn.SendChan <- data
}
