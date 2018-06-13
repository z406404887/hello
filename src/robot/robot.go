package robot

import (
	"log"
	"network"
	"github.com/golang/protobuf/proto"
)

type Robot struct {
	account  string
	password string
	srvAddr  string
	Id       uint32
	Name 	 string
	Money int32
	ws       *WsClient
}

func NewRobot(acc string, passwd string, srv string) *Robot{
	r := &Robot{
		account:  acc,
		password: passwd,
		srvAddr:  srv,
		Id:       0,
	}

	r.ws = NewWsClient(srv)
	return r
}

func (robot *Robot) Run()  {
	Login(robot)
	for {
		select {
		case msg ,ok := <- robot.ws.RecvChan:
			if !ok {
				log.Printf("connection closed, exit.")
				close(robot.ws.SendChan)
				break
			}
			handleMsg(robot,msg)
		}
	}
}

func (robot *Robot) SendMsg(main uint8,sub uint8, msg proto.Message){
	send, _ := proto.Marshal(msg)
	header := &network.CommonHeader{
		MainType:main,
		SubType:sub,
		ClientId:robot.Id,
		Result:0,
		Len:uint16(len(send)),
	}

	data := append(header.Encode(), send...)
	robot.ws.SendChan <- data
}