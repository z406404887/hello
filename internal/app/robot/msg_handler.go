package robot

import (
	"log"
	"starter-kit/internal/pkg/network"
	"starter-kit/internal/pkg/pb/pbgame"
	"time"

	"github.com/golang/protobuf/proto"
)

func handleMsg(robot *Robot, msg []byte) {
	header := &network.CommonHeader{}
	header.Decode(msg)
	//log.Printf("handle msg header %+v",header)
	log.Printf("msg: %v", msg)
	data := msg[network.COMMON_HEADER_LENGTH:]
	switch header.MainType {
	case pbgame.MainAccount:
		handleLoginMsg(robot, header, data)
	case pbgame.MainGame:
		handleGameMsg(robot, header, data)
	default:
		log.Printf("no handler found %+v", header)
	}

}

//login
func handleLoginMsg(robot *Robot, header *network.CommonHeader, data []byte) {
	switch header.SubType {
	case pbgame.SubLoginRsp:
		handleLoginRsp(robot, header, data)
	}
}

func handleLoginRsp(robot *Robot, header *network.CommonHeader, data []byte) {
	log.Printf("header %+v", header)
	rsp := &pbgame.LoginResponse{}

	err := proto.Unmarshal(data, rsp)

	if err != nil {
		log.Printf("unmarshal login response failed. %v", err)
	}
}

//game
func handleGameMsg(robot *Robot, header *network.CommonHeader, data []byte) {
	switch header.SubType {
	case pbgame.SubEnterGameRsp:
		handleEnterGameMsg(robot, header, data)
	case pbgame.SubRollRsp:
		handleRollResponse(robot, header, data)
	default:
		log.Printf("no handler found. %+v", header)
	}
}

func handleRollResponse(robot *Robot, heaer *network.CommonHeader, data []byte) {
	rsp := &pbgame.RollResponse{}
	if err := proto.Unmarshal(data, rsp); err != nil {
		log.Printf("unmarshal failed. %v", err)
		return
	}
	if rsp.Win > 0 {
		log.Printf("win:%d", rsp.Win)
		log.Printf("previous:%d", robot.Money)
		log.Printf("balance:%d", robot.Money+rsp.Win)
		log.Printf("woooow! you are the winner. ")
	} else {
		log.Printf("lose:%d", rsp.Win*-1)
		log.Printf("previous:%d", robot.Money)
		log.Printf("balance:%d", robot.Money+rsp.Win)
		log.Printf("come on! big win is waiting for you. ")
	}
	robot.Money += rsp.Win

	ticker := time.After(5 * time.Second)
	<-ticker
	log.Printf("let's fight again!")
	Rolll(robot)
}

func handleEnterGameMsg(robot *Robot, header *network.CommonHeader, data []byte) {
	rsp := &pbgame.EnterGameResponse{}
	err := proto.Unmarshal(data, rsp)

	if err != nil {
		log.Printf("unmarshal enter game msg failed. %v", err)
		return
	}
	log.Printf("game msg header:%v, rsp:%v", header, rsp)
	robot.Id = rsp.Uid
	robot.Name = rsp.Name
	robot.Money = rsp.Money

	Rolll(robot)
}
