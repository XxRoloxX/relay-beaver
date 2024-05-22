package main

import (
	"backend/internal/auth"
	"backend/internal/proxy_rule"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	godotenv.Load()

	// requestRepository := request.RequestMongoRepository{Db: *database}
	// requestService := request.RequestService{Repo: requestRepository}
	// requestHandler := request.RequestHandler{Service: requestService}

	authMiddleware := auth.NewAuthMiddleware()
	_ = auth.GetAuthRouter(router.PathPrefix("/auth").Subrouter())
	proxyRuleRouter := proxyrule.GetProxyRuleRouter(router.PathPrefix("/proxy-rules").Subrouter())
	proxyRuleRouter.Use(authMiddleware.Handler)

	// router.HandleFunc("/", handler).Methods("GET")
	// router.HandleFunc("/requests", requestHandler.GetRequestsHandler).Methods("GET")

	http.ListenAndServe(":8080", router)
}
