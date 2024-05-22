package proxyevents

import (
	connectionpool "backend/pkg/connection_pool"

	"github.com/gorilla/mux"
)

func GetProxyEventsRouter(connectionHub *connectionpool.Hub, router *mux.Router) *mux.Router {

	authHandler := NewProxyEventsHandler(connectionHub)
	router.HandleFunc("", authHandler.WebsocketRequestsHandler)

	return router
}
