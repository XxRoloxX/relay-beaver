package auth

import "github.com/gorilla/mux"

func GetAuthRouter(router *mux.Router) *mux.Router {

	authHandler := NewAuthHandler()
	router.HandleFunc("/login", authHandler.LoginHandler).Methods("POST")
	router.HandleFunc("/code", authHandler.AuthCodeHandler).Methods("GET")
	router.HandleFunc("/validate", authHandler.ValidateTokenHandler).Methods("POST")

	return router
}
