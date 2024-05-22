package clientevents

import (
	connectionpool "backend/pkg/connection_pool"

	"github.com/gorilla/mux"
)

func GetClientEventsRouter(connectionHub *connectionpool.Hub, router *mux.Router) *mux.Router {

	clientEventsHandler := NewClientEventsHandler(connectionHub)
	router.HandleFunc("", clientEventsHandler.HandleClientEvent)

	return router
}
