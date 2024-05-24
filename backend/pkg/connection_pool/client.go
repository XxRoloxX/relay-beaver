package connectionpool

import "github.com/gorilla/websocket"

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
}

func NewClient() *Client {
	return &Client{
		send: make(chan []byte, 256),
	}
}
func (c *Client) SetHub(hub *Hub) *Client {
	c.hub = hub
	return c
}

func (c *Client) SetConn(conn *websocket.Conn) *Client {
	c.conn = conn
	return c
}

func (c *Client) ReadPump() {
	defer func() {
		c.hub.Unregister(c)
		c.conn.Close()
	}()
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			break
		}
		print(message)
		// c.hub.Broadcast(message)
	}
}

func (c *Client) WritePump() {
	defer func() {
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.conn.WriteMessage(websocket.TextMessage, message)
		}
	}
}
