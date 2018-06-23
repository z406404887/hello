package game

import (
	"fmt"
	"hello/internal/pkg/network"
	"log"
	"net/http"
	"runtime"
	"sync"

	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
)

type Game struct {
	cfg *Configuration

	recvConnMsg chan *network.Message
	closeChan   chan struct{}
	errChan     chan error

	cid       uint32
	muId      sync.Mutex
	connMap   map[uint32]*network.WsConn
	regConn   chan *network.WsConn
	unRegConn chan *network.WsConn

	//loign
	dbConn *grpc.ClientConn
}

func NewGame(cfgPath string) (*Game, error) {
	cfg, err := NewConfiguration(cfgPath)
	if err != nil {
		return nil, err
	}
	log.Printf("cfg %v", cfg)
	game := &Game{
		cfg:         cfg,
		recvConnMsg: make(chan *network.Message, 1000),
		closeChan:   make(chan struct{}),
		errChan:     make(chan error),
		cid:         0,
		connMap:     make(map[uint32]*network.WsConn),
		regConn:     make(chan *network.WsConn),
		unRegConn:   make(chan *network.WsConn),
		muId:        sync.Mutex{},
	}
	return game, nil
}

func (game *Game) GetId() uint32 {
	game.muId.Lock()
	defer game.muId.Unlock()
	game.cid++
	return game.cid
}

func (game *Game) getDbConn() *grpc.ClientConn {
	return game.dbConn
}

func (game *Game) connectDb() {
	conn, err := grpc.Dial(game.cfg.DbAddr, grpc.WithInsecure(), grpc.WithWaitForHandshake())
	if err != nil {
		log.Printf("connect to db server failed. %v", err)
	}
	game.dbConn = conn

}

func (game *Game) onNewWsConn(conn *network.WsConn) {
	log.Printf("client connected. id=%d", conn.Id)
	game.connMap[conn.Id] = conn
}

func (game *Game) onCloseConn(conn *network.WsConn) {
	delete(game.connMap, conn.Id)
}

func (game *Game) handleConnMsg(msg *network.Message) {
	fmt.Printf("recv msg %v", msg)
	if ws, ok := game.connMap[msg.Head.ClientId]; ok {
		ws.SendChan <- append(msg.Head.Encode(), msg.Data...)
	}
}

func (game *Game) onClose() {

}

func (game *Game) Run() {
	go game.ServeWs()
	game.connectDb()
	game.doWork()
}

func (gate *Game) doWork() {
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
		case err := <-gate.errChan:
			log.Printf("server error %v \n", err)
			return
		}
	}
}

var upgrader = websocket.Upgrader{}

func (game *Game) ServeWs() {
	http.HandleFunc("/", game.HandleHttpMsg)
	log.Printf("game listen on %s", game.cfg.Addr)
	err := http.ListenAndServe(game.cfg.Addr, nil)
	log.Println("goroutine num", runtime.NumGoroutine())
	if err != nil {
		game.errChan <- err
	}
}

func (gate *Game) HandleHttpMsg(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	ws := network.NewWsConn(gate.GetId(), conn, 1000)

	agent := NewConnAgent(gate, ws)

	agent.Run()
}
