package gateway

import (
	"context"
	"hello/internal/pkg/network"
	"hello/internal/pkg/pb/pbgame"
	"log"
	"time"

	"github.com/golang/protobuf/proto"
)

func handleAccountMsg(agent *ConnAgent, header *network.CommonHeader, data []byte) bool {
	//log.Printf("handle account msg, header %+v",header)
	switch header.SubType {
	case pbgame.SubLoginReq:
		handleLoginRequest(agent, header, data)
	default:
		return false
	}
	return true
}

func handleLoginRequest(agent *ConnAgent, header *network.CommonHeader, data []byte) {
	req := &pbgame.LoginRequest{}

	err := proto.Unmarshal(data, req)
	if err != nil {
		log.Printf("unmarshal login request failed. %v", err)
		return
	}

	c := pbgame.NewLoginClient(agent.gate.loginConn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	rsp, err := c.Login(ctx, req)

	if err != nil {
		log.Printf("call service failed. %v", err)
		rsp = &pbgame.LoginResponse{
			ErrorCode: pbgame.ErrorCode_CALL_SERVICE_FAILD,
		}
	}

	header.SubType = pbgame.SubLoginRsp
	rsp_data, _ := proto.Marshal(rsp)
	agent.sendMsgBack(header, rsp_data)

	//login success, send enter game
	enterReq := &pbgame.EnterGameRequest{
		Account: req.Account,
	}

	enterData, err := proto.Marshal(enterReq)
	if err != nil {
		log.Printf("marshal enter game request failed. %v", err)
		return
	}

	newHeader := &network.CommonHeader{
		MainType: pbgame.MainGame,
		SubType:  pbgame.SubEnterGameReq,
		ClientId: header.ClientId,
		Result:   0,
		Len:      uint16(len(enterData)),
	}

	//log.Printf("send to game header %+v",newHeader)
	msg := &network.Message{
		SrcId: agent.ws.Id,
		Head:  newHeader,
		Data:  enterData,
	}
	agent.gate.recvConnMsg <- msg
}
