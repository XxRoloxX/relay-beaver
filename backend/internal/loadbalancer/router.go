package loadbalancer

import "github.com/gorilla/mux"

func GetLBRouter(router *mux.Router) *mux.Router {

	lbHandler := NewLBHandler()

	router.HandleFunc("", lbHandler.CreateLBHandler).Methods("POST")
	router.HandleFunc("", lbHandler.GetLBHandler).Methods("GET")
	router.HandleFunc("/{id}", lbHandler.UpdateLBHandler).Methods("PUT")
	router.HandleFunc("/{id}", lbHandler.DeleteLBHandler).Methods("DELETE")

	return router
}
