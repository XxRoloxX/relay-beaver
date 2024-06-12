package proxyevent

import (
	"backend/internal/database"
	"backend/internal/logger"
	"backend/internal/request"
	connectionpool "backend/pkg/connection_pool"
	"backend/pkg/models"
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
)

type ProxyEventsHandler struct {
	Logger        logger.HttpLogger
	Service       ProxyEventsService
	ConnectionHub *connectionpool.Hub
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func NewProxyEventsHandler(connectionHub *connectionpool.Hub) ProxyEventsHandler {
	return ProxyEventsHandler{
		Logger: logger.HttpLogger{},
		Service: ProxyEventsService{
			Repo: request.RequestMongoRepository{
				Db: *database.Db,
			},
		},
		ConnectionHub: connectionHub,
	}
}

type EventMessage struct {
	Type           string                `json:"type"`
	ProxiedRequest models.ProxiedRequest `json: "proxiedRequest"`
}

func (handler *ProxyEventsHandler) WebsocketRequestsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	logger := handler.Logger.Request(r)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			logger.Error(err.Error())
			return
		}
		var eventMessage EventMessage

		err = json.
			NewDecoder(bytes.NewReader(message)).
			Decode(&eventMessage)

		if err != nil {
			logger.Error(err.Error())
			return
		}

		proxiedRequest := eventMessage.ProxiedRequest
		proxiedRequest, err = handler.Service.HandleProxiedRequest(proxiedRequest)
		encodedProxiedRequest, err := json.Marshal(proxiedRequest)

		if err != nil {
			logger.Error(err.Error())
			return
		}

		handler.ConnectionHub.Broadcast(encodedProxiedRequest)
	}
}
