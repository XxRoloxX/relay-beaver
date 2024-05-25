package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
	"net/http"
	"proxy/internal/request"
)

type WebsocketServer struct {
	upgrader websocket.Upgrader
	channel  chan request.Request
	port     int
}

func NewWebsocketServer(channel chan request.Request, port int) *WebsocketServer {
	return &WebsocketServer{
		upgrader: websocket.Upgrader{},
		channel:  channel,
		port:     port,
	}
}

func (u *WebsocketServer) sendRequests(w http.ResponseWriter, r *http.Request) {
	conn, err := u.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error().Msg(fmt.Sprintf("Connection upgrade error: %s", err))
		return
	}

	defer conn.Close()

	for {
		select {
		case req := <-u.channel:
			res, err := json.Marshal(&req)
			if err != nil {
				log.Error().Msg("Error parsing request as JSON")
			}

			err = conn.WriteMessage(1, res)
			if err != nil {
				log.Error().Msg("Error writing request message")
			}
		}
	}
}

func (u *WebsocketServer) Start() error {
	http.HandleFunc("/requests", u.sendRequests)

	log.Info().Msg(fmt.Sprintf("Starting websocket proxy on port: %s", u.port))

	err := http.ListenAndServe(fmt.Sprintf("localhost:%d", u.port), nil)
	if err != nil {
		log.Error().Msg("Failed to start websocket proxy")
		return err
	}

	return nil
}

func (u *WebsocketServer) GetChannel() chan request.Request {
	return u.channel
}
