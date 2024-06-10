package stats

import "github.com/gorilla/mux"

func GetStatsRouter(router *mux.Router) *mux.Router {

	statsHandler := NewStatsHandler()

	router.HandleFunc("", statsHandler.GetStatsHandler).Methods("GET")
	router.HandleFunc("/hosts", statsHandler.GetHostsHandler).Methods("GET")

	return router
}
