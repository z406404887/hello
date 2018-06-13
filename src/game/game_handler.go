package game

import (
	"network"
	"pb/pbgame"
	"github.com/golang/protobuf/proto"
	"log"
	"context"
	"time"
	"math/rand"
)

func handleMsg(agent *ConnAgent, header *network.CommonHeader,data []byte)  {
	log.Printf("handleMsg header=%+v",header)
	switch header.MainType {
	case pbgame.MainGame:
		handleGameMsg(agent,header,data)
	}
}

func handleGameMsg(agent *ConnAgent, header *network.CommonHeader,data []byte)  {
	switch header.SubType {
	case pbgame.SubEnterGameReq:
		handleEnterGame(agent,header,data)
	case pbgame.SubRollReq:
		handleRoll(agent,header,data)
	}
}

func handleRoll(agent *ConnAgent, header *network.CommonHeader,data []byte)  {
	player,ok := agent.room.playerMap[header.ClientId]
	if !ok {
		log.Printf("player not found. id=%d",header.ClientId)
		return
	}

	rand.Seed(time.Now().UnixNano())
	result := rand.Int31n(100)
	win := result - 50
	rsp := &pbgame.RollResponse{
		Win:int32(win),
	}

	player.money += win

	svcReq := &pbgame.SaveRequest{
		Uid:player.id,
		Money:player.money,
	}
	ctx,cancel := context.WithTimeout(context.Background(),time.Second)

	svcRsp, err := agent.dbClient.SavePlayer(ctx,svcReq)
	cancel()

	if err != nil{
		if svcRsp == nil {
			header.Result = uint16(pbgame.ErrorCode_MYSQL_ERROR)
		}else {
			header.Result = uint16(svcRsp.Result)
		}
		log.Printf("save error. %v",err)
		rsp.Win = 0
	}
	header.SubType = pbgame.SubRollRsp
	agent.sendMsgBack(header,rsp)
}

func handleEnterGame(agent *ConnAgent, header *network.CommonHeader,data []byte)  {
	log.Printf("enter game header %+v",header)
	header.SubType = pbgame.SubLoginRsp
	req := &pbgame.EnterGameRequest{}
	rsp := &pbgame.EnterGameResponse{}

	err := proto.Unmarshal(data,req)
	if err != nil {
		log.Printf("unmarshal failed. %v",err)
		return
	}


	header.Result = uint16(doEnterGame(agent,req,rsp))
	if header.Result == 0{
		player := NewPlayer(req.Account,rsp.Name,rsp.Uid,rsp.Money)
		agent.room.AddPlayer(player)
	}
	agent.sendMsgBack(header,rsp)
}

func doEnterGame(agent *ConnAgent,req *pbgame.EnterGameRequest,rsp *pbgame.EnterGameResponse) pbgame.ErrorCode  {
	ctx,cancel := context.WithTimeout(context.Background(),time.Second)

	svcReq := &pbgame.LoadRequest{
		Account:req.Account,
	}

	svcRsp, err := agent.dbClient.LoadPlayer(ctx,svcReq)
	cancel()
	if err != nil {
		log.Printf("load player failed. %v",err)
		return pbgame.ErrorCode_MYSQL_ERROR
	}

	if svcRsp.Result != pbgame.ErrorCode_ACCOUNT_NOT_EXISTS {
		rsp.Uid = svcRsp.Uid
		rsp.Name = svcRsp.Name
		rsp.Money = svcRsp.Money
		return pbgame.ErrorCode_SUCCESS
	}

	//if not exists, create new player
	newReq := &pbgame.CreatePlayerRequest{
		Account:req.Account,
		Name:req.Account,
		Money:10000,
	}
	ctx, cancel = context.WithTimeout(context.Background(),time.Second)
	newRsp, err := agent.dbClient.CreatePlayer(ctx,newReq)
	if err != nil {
		return newRsp.Result
	}

	rsp.Uid = newRsp.Uid
	rsp.Name = newReq.Account
	rsp.Money = newReq.Money
	player := NewPlayer(req.Account,rsp.Name,rsp.Uid,rsp.Money)
	agent.room.AddPlayer(player)
	return pbgame.ErrorCode_SUCCESS
}