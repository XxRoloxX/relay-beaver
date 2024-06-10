package main

import (
	"backend/internal/auth"
	clientevent "backend/internal/client_event"
	"backend/internal/common"
	"backend/internal/logger"
	proxyevent "backend/internal/proxy_event"
	proxyrule "backend/internal/proxy_rule"
	"backend/internal/stats"
	connectionpool "backend/pkg/connection_pool"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	connectionHub := connectionpool.NewHub()
	go connectionHub.Run()

	authMiddleware := auth.NewAuthMiddleware()
	loggerMiddleware := logger.NewLoggerMiddleware()
	router.Use(loggerMiddleware.Handler)

	_ = auth.GetAuthRouter(router.PathPrefix("/auth").Subrouter())

	_ = proxyrule.GetProxyRuleRouter(router.PathPrefix("/proxy-rules").Subrouter())
	// proxyRuleRouter.Use(authMiddleware.Handler)

	proxyEventsRouter := proxyevent.GetProxyEventsRouter(connectionHub, router.PathPrefix("/proxy-events").Subrouter())
	proxyEventsRouter.Use(authMiddleware.Handler)

	clientEventsRouter := clientevent.GetClientEventsRouter(connectionHub, router.PathPrefix("/client-events").Subrouter())
	clientEventsRouter.Use(authMiddleware.Handler)

	_ = stats.GetStatsRouter(router.PathPrefix("/stats").Subrouter())
	// statsRouter.Use(authMiddleware.Handler)

	http.ListenAndServe(":8080", common.HandleCors(router))
}
