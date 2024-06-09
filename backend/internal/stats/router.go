package stats

import "github.com/gorilla/mux"

func GetStatsRouter(router *mux.Router) *mux.Router {

	statsHandler := NewStatsHandler()

	router.HandleFunc("", statsHandler.GetStats).Methods("GET")

	return router
}
