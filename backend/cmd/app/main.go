package main

import (
	"backend/internal/auth"
	clientevent "backend/internal/client_event"
	"backend/internal/common"
	"backend/internal/database"
	"backend/internal/logger"
	proxyevent "backend/internal/proxy_event"
	proxyrule "backend/internal/proxy_rule"
	connectionpool "backend/pkg/connection_pool"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env.example")
	if err != nil {
		panic(err)
	}

	database.InitializeDatabase()
	router := mux.NewRouter()

	connectionHub := connectionpool.NewHub()
	go connectionHub.Run()

	authMiddleware := auth.NewAuthMiddleware()
	loggerMiddleware := logger.NewLoggerMiddleware()
	router.Use(loggerMiddleware.Handler)

	_ = auth.GetAuthRouter(router.PathPrefix("/auth").Subrouter())

	proxyRuleRouter := proxyrule.GetProxyRuleRouter(router.PathPrefix("/proxy-rules").Subrouter())
	proxyRuleRouter.Use(authMiddleware.Handler)

	proxyEventsRouter := proxyevent.GetProxyEventsRouter(connectionHub, router.PathPrefix("/proxy-events").Subrouter())
	proxyEventsRouter.Use(authMiddleware.Handler)

	clientEventsRouter := clientevent.GetClientEventsRouter(connectionHub, router.PathPrefix("/client-events").Subrouter())
	clientEventsRouter.Use(authMiddleware.Handler)

	http.ListenAndServe(":8080", common.HandleCors(router))
}
