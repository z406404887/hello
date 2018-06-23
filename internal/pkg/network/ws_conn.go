package network

import (
	"log"
	"net"
	"time"

	"hello/internal/pkg/util"

	"github.com/gorilla/websocket"
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

func (c *WsConn) setReadDeadline(d time.Time) {
	if err := c.conn.SetReadDeadline(d); err != nil {
		log.Printf("connection SetReadDeadline failed. %v", err)
	}
}

func (c *WsConn) setWriteDeadline(d time.Time) {
	if err := c.conn.SetReadDeadline(d); err != nil {
		log.Printf("connection SetWriteDeadline failed. %v", err)
	}
}

func (c *WsConn) writeMessage(messageType int, data []byte) {
	if err := c.conn.WriteMessage(messageType, data); err != nil {
		log.Printf("write msg failed. messageType=%d, %v", messageType, err)
	}
}

func (ws *WsConn) writePump() {
	ticker := time.NewTicker(PingPeriod)
	defer func() {
		ticker.Stop()
		util.Close(ws.conn)
	}()

	for {
		select {
		case message, ok := <-ws.SendChan:
			if !ok {
				ws.setWriteDeadline(time.Now().Add(WriteWait))
				//channel has been closed, close the connection
				ws.writeMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := ws.conn.NextWriter(websocket.BinaryMessage)
			if err != nil {
				return
			}
			if _, err := w.Write(message); err != nil {
				log.Printf("WsConn write msg failed. %v", err)
			}

			n := len(ws.SendChan)
			for i := 0; i < n; i++ {
				if _, err := w.Write(<-ws.SendChan); err != nil {
					log.Printf("WsConn write msg failed. %v", err)
				}
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

	ws := &WsConn{
		Id:       id,
		SendChan: make(chan []byte, sendChanSize),
		conn:     conn,
	}

	ws.setReadDeadline(time.Now().Add(PongWait))
	conn.SetPongHandler(func(string) error { ws.setReadDeadline(time.Now().Add(PongWait)); return nil })
	go ws.writePump()
	return ws
}
