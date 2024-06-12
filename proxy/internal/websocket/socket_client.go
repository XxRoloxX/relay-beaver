package websocket

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"proxy/internal/env"
	"proxy/internal/request"

	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
)

type WebsocketClient struct {
	upgrader websocket.Upgrader
	channel  chan request.ProxiedRequest
	host     string
	port     int
	address  string
	endpoint string
}

type EventMessage struct {
	Type           string                 `json:"type"`
	ProxiedRequest request.ProxiedRequest `json:"proxiedRequest"`
}

func NewWebsocketServer(channel chan request.ProxiedRequest, host string, port int, endpoint string) *WebsocketClient {
	return &WebsocketClient{
		upgrader: websocket.Upgrader{},
		channel:  channel,
		host:     host,
		port:     port,
		address:  fmt.Sprintf("%s:%d", host, port),
		endpoint: endpoint,
	}
}

func (u *WebsocketClient) Connect() error {
	wsUrl := url.URL{Scheme: "ws", Host: u.address, Path: u.endpoint}

	log.Info().Msg(fmt.Sprintf("Connecting to websocket server at: %s", wsUrl.String()))
	headers := make(map[string][]string)
	headers[env.GetProxyBackendAuthHeader()] = []string{env.GetProxyBackendAuthSecret()}

	conn, _, err := websocket.DefaultDialer.Dial(wsUrl.String(), headers)
	if err != nil {
		log.Error().Msg(fmt.Sprintf("Error connecting to websocket server: %s", err))
		return err
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	for {
		select {
		case req := <-u.channel:

			message := EventMessage{
				Type:           "proxiedRequest",
				ProxiedRequest: req,
			}
			res, err := json.Marshal(&message)
			if err != nil {
				log.Error().Msg("Error parsing request as JSON")
			}

			err = conn.WriteMessage(1, res)

			if err != nil {
				log.Error().Err(err).Msg("Error writing request message")
				log.Error().Msg("Error writing request message")
			}
		case <-interrupt:
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Error().Msg("Error writing close message")
				return err
			}
		}
	}
}
