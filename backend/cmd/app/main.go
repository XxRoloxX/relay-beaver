package main

import (
	"backend/internal/auth"
	"backend/internal/client_events"
	"backend/internal/proxy_events"
	"backend/internal/proxy_rule"
	"backend/pkg/connection_pool"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	godotenv.Load()

	connectionHub := connectionpool.NewHub()
	go connectionHub.Run()

	authMiddleware := auth.NewAuthMiddleware()

	_ = auth.GetAuthRouter(router.PathPrefix("/auth").Subrouter())

	proxyRuleRouter := proxyrule.GetProxyRuleRouter(router.PathPrefix("/proxy-rules").Subrouter())
	proxyRuleRouter.Use(authMiddleware.Handler)

	_ = proxyevents.GetProxyEventsRouter(connectionHub, router.PathPrefix("/ws").Subrouter())

	_ = clientevents.GetClientEventsRouter(connectionHub, router.PathPrefix("/client-events").Subrouter())

	http.ListenAndServe(":8080", router)
}
