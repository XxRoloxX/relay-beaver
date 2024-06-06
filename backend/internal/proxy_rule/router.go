package proxyrule

import "github.com/gorilla/mux"

func GetProxyRuleRouter(router *mux.Router) *mux.Router {

	proxyRuleHandler := NewProxyRuleHandler()

	router.HandleFunc("", proxyRuleHandler.CreateProxyRuleHandler).Methods("POST")
	router.HandleFunc("", proxyRuleHandler.GetProxyRuleHandler).Methods("GET")
	router.HandleFunc("/{id}", proxyRuleHandler.UpdateProxyRuleHandler).Methods("PUT")
	router.HandleFunc("/{id}", proxyRuleHandler.DeleteProxyRuleHandler).Methods("DELETE")

	return router
}
