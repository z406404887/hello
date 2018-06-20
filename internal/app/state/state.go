package state

import (
	"net/http"
	"github.com/gorilla/websocket"
	"log"
	"hello/internal/pkg/network"
	"fmt"
	"runtime"
	"sync"
)

type StateServer struct {
	cfg *Configuration

	recvConnMsg chan *network.Message
	closeChan   chan struct{}
	errChan     chan error

	cid       uint32
	muId      sync.Mutex
	connMap   map[uint32]*network.WsConn
	regConn   chan *network.WsConn
	unRegConn chan *network.WsConn
}

func NewStateServer(cfgPath string) (*StateServer,error){
	cfg , err := NewConfiguration(cfgPath)
	if err != nil {
		return  nil , err
	}
	log.Printf("cfg %v",cfg)
	state := &StateServer{
		cfg:cfg,
		recvConnMsg: make(chan *network.Message,1000),
		closeChan:   make(chan struct{}),
		errChan:     make(chan error),
		cid:         0,
		connMap:     make(map[uint32]*network.WsConn),
		regConn:     make(chan *network.WsConn),
		unRegConn:   make(chan *network.WsConn),
		muId:        sync.Mutex{},
	}
	return state,nil
}

func (srv *StateServer) GetId() uint32  {
	srv.muId.Lock()
	defer srv.muId.Unlock()
	srv.cid++
	return  srv.cid
}

func (srv *StateServer) onNewWsConn(conn *network.WsConn){
	log.Printf("client connected. id=%d",conn.Id)
	srv.connMap[conn.Id] = conn
}

func (srv *StateServer) onCloseConn(conn *network.WsConn){
	delete(srv.connMap,conn.Id)
}

func (srv *StateServer) handleConnMsg(msg *network.Message){
	fmt.Printf("recv msg %v",msg)
	if ws, ok := srv.connMap[msg.Head.ClientId] ; ok {
		ws.SendChan <- append(msg.Head.Encode(),msg.Data...)
	}
}


func (srv *StateServer) onClose(){

}

func (srv *StateServer) Run() {
	go srv.ServeWs()
	srv.doWork()
}


func (srv *StateServer) doWork(){
	for {
		select {
		case conn := <-srv.regConn:
			srv.onNewWsConn(conn)
		case conn := <-srv.unRegConn:
			srv.onCloseConn(conn)
			close(conn.SendChan)
		case msg := <- srv.recvConnMsg:
			srv.handleConnMsg(msg)
		case <-srv.closeChan:
			srv.onClose()
		case err := <- srv.errChan:
			log.Printf("server error %v \n",err)
			break
		}
	}
}

var upgrader = websocket.Upgrader{}
func (srv *StateServer) ServeWs(){
	http.HandleFunc("/", srv.HandleHttpMsg)
	err := http.ListenAndServe(srv.cfg.Addr,nil)
	log.Println("goroutine num" ,runtime.NumGoroutine())
	if err != nil {
		srv.errChan <- err
	}
}


func (srv *StateServer) HandleHttpMsg(w http.ResponseWriter, r *http.Request){
	conn, err := upgrader.Upgrade(w,r,nil)
	if err != nil {
		log.Println(err)
		return
	}

	ws := network.NewWsConn(srv.GetId(),conn,1000)

	agent := NewConnAgent(srv,ws)

	agent.Run()
}