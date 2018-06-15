package network

import (
	"github.com/gorilla/websocket"
	"log"
	"net"
	"time"
)

type WsConn struct {
	Id       uint32
	SendChan chan []byte
	conn     *websocket.Conn
}

func (ws *WsConn) RemoteAddr() net.Addr {
	return ws.conn.RemoteAddr()
}

func (ws *WsConn) ReadMsg() (int, []byte, error) {
	return ws.conn.ReadMessage()
}

func (ws *WsConn) writePump() {
	ticker := time.NewTicker(PingPeriod)
	defer func() {
		ticker.Stop()
		ws.conn.Close()
	}()

	for {
		select {
		case message, ok := <-ws.SendChan:
			if !ok {
				ws.conn.SetWriteDeadline(time.Now().Add(WriteWait))
				//channel has been closed, close the connection
				ws.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := ws.conn.NextWriter(websocket.BinaryMessage)
			if err != nil {
				return
			}
			w.Write(message)

			n := len(ws.SendChan)
			for i := 0; i < n; i++ {
				w.Write(<-ws.SendChan)
			}

			if err := w.Close(); err != nil {
				log.Println("Write error", err)
				return
			}
		case <-ticker.C:
			if err := ws.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func NewWsConn(id uint32, conn *websocket.Conn, sendChanSize int) *WsConn {
	if sendChanSize == 0 {
		sendChanSize = DefaultSendChanSize
	}

	conn.SetReadDeadline(time.Now().Add(PongWait))
	conn.SetPongHandler(func(string) error { conn.SetReadDeadline(time.Now().Add(PongWait)); return nil })
	ws := &WsConn{
		Id:       id,
		SendChan: make(chan []byte, sendChanSize),
		conn:     conn,
	}

	go ws.writePump()
	return ws
}
