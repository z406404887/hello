package robot
import (
	"github.com/gorilla/websocket"
	"time"
	"log"
)

type WsClient struct {
	SrvAddr  string
	conn     *websocket.Conn
	SendChan chan []byte
	RecvChan chan []byte
}

func NewWsClient(srv string) *WsClient  {
	c := &WsClient{
		SrvAddr:   srv,
		SendChan:make(chan []byte,5000),
		RecvChan:make(chan []byte, 5000),
	}
	conn, err := c.dial()
	if err != nil {
		log.Printf("dial error, %v",err)
		return nil;
	}

	log.Printf("server connected. %s",c.SrvAddr)
	c.conn = conn
	go c.readPump()
	go c.writePump()

	return c
}

func (client *WsClient) dial() (*websocket.Conn,error)  {
	for {
		conn, _, err := websocket.DefaultDialer.Dial(client.SrvAddr,nil)
		if err == nil {
			return conn,nil
		}

		log.Printf("establish connection failed. %v",err)
		time.Sleep(3 * time.Second)
	}
}

func (c *WsClient) readPump() {
	defer func() {
		c.conn.Close()
		close(c.RecvChan)
	}()

	//c.ws.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(PongWait))
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
			c.conn.SetWriteDeadline(time.Now().Add(WriteWait))
			if !ok {
				c.conn.SetWriteDeadline(time.Now().Add(WriteWait))
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.BinaryMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.SendChan)
			for i := 0; i < n; i++ {
				w.Write(<-c.SendChan)
			}

			if err := w.Close(); err != nil {
				log.Println("Write error", err)
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(WriteWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}