package clientevent

import (
	"backend/pkg/connection_pool"
	"github.com/gorilla/websocket"
	"net/http"
)

type ClientEventsHandler struct {
	connectionHub *connectionpool.Hub
}

func NewClientEventsHandler(connectionHub *connectionpool.Hub) *ClientEventsHandler {
	return &ClientEventsHandler{
		connectionHub: connectionHub,
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (handler *ClientEventsHandler) HandleClientEvent(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		println(err.Error())
	}

	client := connectionpool.
		NewClient().
		SetHub(handler.connectionHub).
		SetConn(conn)

	handler.connectionHub.Register(client)

	go client.ReadPump()
	go client.WritePump()

	// for {
	// 	_, message, err := conn.ReadMessage()
	// 	if err != nil {
	// 	}
	// 	println(string(message))
	// }
}
