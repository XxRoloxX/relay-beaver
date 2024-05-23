package auth

import "github.com/gorilla/mux"

func GetAuthRouter(router *mux.Router) *mux.Router {

	authHandler := NewAuthHandler()
	router.HandleFunc("/login", authHandler.LoginHandler).Methods("POST")
	router.HandleFunc("/profile", authHandler.TokenInfoHandler).Methods("GET")

	return router
}
