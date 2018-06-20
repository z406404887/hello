package gateway

import (
	"context"
	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"network"
	"pb/pbgame"
	"runtime"
	"sync"
	"time"
	"util"
)

type Gateway struct {
	cfg *Configuration

	recvConnMsg chan *network.Message
	recvSrvMsg  chan *network.Message
	closeChan   chan struct{}
	errChan     chan error

	cid       uint32
	muId      sync.Mutex
	connMap   map[uint32]*network.WsConn
	regConn   chan *network.WsConn
	unRegConn chan *network.WsConn

	gameMap  map[uint32]*network.WsClient
	regSrv   chan *network.WsClient
	unRegSrv chan *network.WsClient

	//loign
	loginConn *grpc.ClientConn

	curGameIdx uint32

	uidSrvMap map[uint32]uint32

	stateSrv *network.WsClient

	//track time consume
	traceTime map[uint32]int64
}

func NewGateway(cfgPath string) (*Gateway, error) {
	cfg, err := NewConfiguration(cfgPath)
	if err != nil {
		return nil, err
	}
	log.Printf("cfg %v", cfg)
	gate := &Gateway{
		cfg:         cfg,
		recvConnMsg: make(chan *network.Message, 1000),
		closeChan:   make(chan struct{}),
		errChan:     make(chan error),
		cid:         0,
		connMap:     make(map[uint32]*network.WsConn),
		regConn:     make(chan *network.WsConn),
		unRegConn:   make(chan *network.WsConn),
		muId:        sync.Mutex{},
		gameMap:     make(map[uint32]*network.WsClient),
		curGameIdx:  0,
		uidSrvMap:   make(map[uint32]uint32),
		regSrv:      make(chan *network.WsClient),
		unRegSrv:    make(chan *network.WsClient),
		recvSrvMsg:  make(chan *network.Message, 1000),
		traceTime:   make(map[uint32]int64),
	}

	return gate, nil
}

func (gate *Gateway) GetGameClientById(id uint32) *network.WsClient {
	sid, ok := gate.uidSrvMap[id]
	if !ok {
		log.Printf("client is not assigned to game server. %d", id)
		return nil
	}

	if c, ok := gate.gameMap[sid]; ok {
		return c
	}
	return nil
}

func (gate *Gateway) GetGameClient() *network.WsClient {
	len := len(gate.gameMap)
	if len == 0 {
		return nil
	}

	mark := gate.curGameIdx % uint32(len)

	var i uint32 = 0
	for _, cli := range gate.gameMap {
		if i == mark {
			gate.curGameIdx++
			return cli
		}
		i++
	}
	return nil
}

func (gate *Gateway) GetId() uint32 {
	gate.muId.Lock()
	defer gate.muId.Unlock()
	gate.cid++
	return gate.cid
}

func (gate *Gateway) getLoginConn() *grpc.ClientConn {
	return gate.loginConn
}

func (gate *Gateway) connectLogin() {
	conn, err := grpc.Dial(gate.cfg.LoginAddr, grpc.WithInsecure(), grpc.WithWaitForHandshake())
	if err != nil {
		log.Printf("connect to login failed. %v", err)
	}
	gate.loginConn = conn

}

func (gate *Gateway) onNewWsConn(conn *network.WsConn) {
	gate.addConn(conn)
}

func (gate *Gateway) getConn(id uint32) *network.WsConn {
	if conn, ok := gate.connMap[id]; ok {
		return conn
	}
	return nil
}

func (gate *Gateway) delConn(id uint32) {
	delete(gate.connMap, id)
}

func (gate *Gateway) addConn(conn *network.WsConn) {
	log.Printf("add connection %d", conn.Id)
	gate.connMap[conn.Id] = conn
}

func (gate *Gateway) onCloseConn(conn *network.WsConn) {
	log.Printf("connection closed. id=%d", conn.Id)
	gate.delConn(conn.Id)
}

func (gate *Gateway) handleConnMsg(msg *network.Message) {
	handleConnMsg(gate, msg)
}

func (gate *Gateway) onNewServerConnected(srv *network.WsClient) {
	id := util.GetIdentity(srv.Id, srv.Type)
	log.Printf("server connected.id = %d", srv.Id)
	switch srv.Type {
	case util.ServerTypeGame:
		gate.gameMap[id] = srv
	case util.ServerTypeState:
		if gate.stateSrv != nil {
			close(gate.stateSrv.SendChan)
			gate.stateSrv = srv
		} else {
			gate.stateSrv = srv
		}
	}
}

func (gate *Gateway) onServerDisconnected(srv *network.WsClient) {
	id := util.GetIdentity(srv.Id, srv.Type)
	delete(gate.gameMap, id)

	//重新连接
	go gate.createClientAgent(srv.Id, srv.Type, srv.SrvAddr)
}

func (gate *Gateway) handleServerMsg(msg *network.Message) {
	//log.Printf("recv server msg %v", msg)
	HandleSrvGameMsg(gate, msg)
}

func (gate *Gateway) onClose() {

}

func (gate *Gateway) Run() {
	go gate.ServeWs()
	gate.connectLogin()
	gate.initServerConnections()
	gate.doWork()
}

func (gate *Gateway) initServerConnections() error {
	rsp, err := gate.getServerList()
	if err != nil {
		log.Printf("get server list failed. %v", err)
		return err
	}
	log.Printf("get server list success. %v", rsp)
	for _, s := range rsp.Server {
		gate.createClientAgent(uint16(s.Id), uint16(s.Type), s.Addr)
	}
	return nil
}

func (gate *Gateway) createClientAgent(sid uint16, stype uint16, addr string) {
	cli := network.NewWsClient(sid, stype, addr)
	agent := NewClientAgent(gate, cli)
	go agent.Run()
}

func (gate *Gateway) getServerList() (*pbgame.ServerListRsp, error) {
	conn, err := grpc.Dial(gate.cfg.MgrAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	log.Printf("connect to manager %s", gate.cfg.MgrAddr)
	defer conn.Close()
	c := pbgame.NewManagerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req := &pbgame.ServerListReq{
		Token: gate.cfg.Token,
		Server: &pbgame.Server{
			Id:   uint32(gate.cfg.Id),
			Type: uint32(gate.cfg.Type),
			Addr: gate.cfg.Addr,
		},
	}

	return c.GetServerList(ctx, req)
}

func (gate *Gateway) doWork() {
	for {
		select {
		case conn := <-gate.regConn:
			gate.onNewWsConn(conn)
		case conn := <-gate.unRegConn:
			gate.onCloseConn(conn)
			close(conn.SendChan)
		case msg := <-gate.recvConnMsg:
			gate.handleConnMsg(msg)
		case <-gate.closeChan:
			gate.onClose()
		case srv := <-gate.regSrv:
			gate.onNewServerConnected(srv)
		case srv := <-gate.unRegSrv:
			gate.onServerDisconnected(srv)
		case msg := <-gate.recvSrvMsg:
			gate.handleServerMsg(msg)
		case err := <-gate.errChan:
			log.Printf("server error %v \n", err)
			break
		}
	}
}

var upgrader = websocket.Upgrader{}

func (gate *Gateway) ServeWs() {
	http.HandleFunc("/", gate.HandleHttpMsg)
	err := http.ListenAndServe(gate.cfg.Addr, nil)
	log.Println("goroutine num", runtime.NumGoroutine())
	if err != nil {
		gate.errChan <- err
	}
}

func (gate *Gateway) HandleHttpMsg(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	ws := network.NewWsConn(gate.GetId(), conn, 1000)

	agent := ConnAgent{
		gate: gate,
		ws:   ws,
	}

	agent.Run()
}
