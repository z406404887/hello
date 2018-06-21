package robot

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
	"hello/internal/pkg/util"
)

type WsClient struct {
	SrvAddr  string
	conn     *websocket.Conn
	SendChan chan []byte
	RecvChan chan []byte
}

func NewWsClient(srv string) *WsClient {
	c := &WsClient{
		SrvAddr:  srv,
		SendChan: make(chan []byte, 5000),
		RecvChan: make(chan []byte, 5000),
	}
	conn, err := c.dial()
	if err != nil {
		log.Printf("dial error, %v", err)
		return nil
	}

	log.Printf("server connected. %s", c.SrvAddr)
	c.conn = conn
	go c.readPump()
	go c.writePump()

	return c
}

func (client *WsClient) dial() (*websocket.Conn, error) {
	for {
		conn, _, err := websocket.DefaultDialer.Dial(client.SrvAddr, nil)
		if err == nil {
			return conn, nil
		}

		log.Printf("establish connection failed. %v", err)
		time.Sleep(3 * time.Second)
	}
}

func (c *WsClient) readPump() {
	defer func() {
		util.Close(c.conn)
		close(c.RecvChan)
	}()

	//c.ws.SetReadLimit(maxMessageSize)
	if err := c.conn.SetReadDeadline(time.Now().Add(PongWait)); err != nil {
		log.Fatalf("connection SetReadline failed. %v",err)
	}

	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(PongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()

		if err != nil {
			log.Println(err)
			break
		}
		c.RecvChan <- message
	}
}

func (c *WsClient) writePump() {
	ticker := time.NewTicker(PingPeriod)
	defer func() {
		ticker.Stop()
	}()
	for {
		select {
		case message, ok := <-c.SendChan:
			//log.Printf("pending send msg %v",message)
			if err := c.conn.SetWriteDeadline(time.Now().Add(WriteWait)); err!= nil {
				log.Fatalf("SetWriteDeadline failed. %v",err)
			}

			if !ok {
				if err := c.conn.SetWriteDeadline(time.Now().Add(WriteWait)); err != nil {
					log.Fatalf("connection SetWriteDeadline failed. %v",err)
				}
				// The hub closed the channel.
				if err := c.conn.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
					log.Fatalf("connection WriteMessage failed.%v",err)
				}
				return
			}

			w, err := c.conn.NextWriter(websocket.BinaryMessage)
			if err != nil {
				return
			}

			if _, err := w.Write(message); err != nil {
				log.Fatalf("write message failed. %v",err)
			}

			// Add queued chat messages to the current websocket message.
			n := len(c.SendChan)
			for i := 0; i < n; i++ {
				if _, err := w.Write(<-c.SendChan); err != nil {
					log.Fatalf("write message failed. %v",err)
				}
			}

			if err := w.Close(); err != nil {
				log.Println("Write error", err)
				return
			}
		case <-ticker.C:
			if err := c.conn.SetWriteDeadline(time.Now().Add(WriteWait)); err != nil {
				log.Fatalf("connection SetWriteDeadline failed. %v",err)
			}
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
