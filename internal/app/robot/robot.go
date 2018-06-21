package robot

import (
	"hello/internal/pkg/network"
	"log"

	"github.com/golang/protobuf/proto"
)

type Robot struct {
	account  string
	password string
	srvAddr  string
	Id       uint32
	Name     string
	Money    int32
	ws       *WsClient
}

func NewRobot(acc string, passwd string, srv string) *Robot {
	r := &Robot{
		account:  acc,
		password: passwd,
		srvAddr:  srv,
		Id:       0,
	}
	return r
}

func (robot *Robot) Run() {
	if ok := robot.doRun(); ok {
		return
	} else {
		robot.Run()
	}
}

func (robot *Robot) doRun() bool {
	robot.ws = NewWsClient(robot.srvAddr)
	Login(robot)
	for msg := range robot.ws.RecvChan {
		handleMsg(robot, msg)
	}
	log.Printf("connection read closed, %s exit. msg=%v", robot.account, msg)
	close(robot.ws.SendChan)
	return false
}

func (robot *Robot) SendMsg(main uint8, sub uint8, msg proto.Message) {
	send, _ := proto.Marshal(msg)
	header := &network.CommonHeader{
		MainType: main,
		SubType:  sub,
		ClientId: robot.Id,
		Result:   0,
		Len:      uint16(len(send)),
	}

	data := append(header.Encode(), send...)
	robot.ws.SendChan <- data
}
