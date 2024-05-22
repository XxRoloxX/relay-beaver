package proxyrule

import (
	"backend/internal/logger"
	"backend/pkg/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

type ProxyRuleHandler struct {
	Service ProxyRuleService
	Logger  logger.HttpLogger
}

func (h *ProxyRuleHandler) GetProxyRuleHandler(w http.ResponseWriter, r *http.Request) {

	logger := h.Logger.Request(r)
	logger.LogRequest()

	requests, error := h.Service.Repo.FindAll()

	if error != nil {
		logger.Error(error.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	serialized, error := json.Marshal(requests)

	w.WriteHeader(http.StatusOK)
	w.Write(serialized)
}

func (h *ProxyRuleHandler) CreateProxyRuleHandler(w http.ResponseWriter, r *http.Request) {

	logger := h.Logger.Request(r)
	logger.LogRequest()

	decoder := json.NewDecoder(r.Body)
	encoder := json.NewEncoder(w)
	var proxyrule models.ProxyRule

	err := decoder.Decode(&proxyrule)

	if err != nil {
		logger.Warn(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
		return
	}

	proxy, err := h.Service.CreateRule(proxyrule)

	if err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	encoder.Encode(proxy)
}

func (h *ProxyRuleHandler) UpdateProxyRuleHandler(w http.ResponseWriter, r *http.Request) {

	logger := h.Logger.Request(r)
	logger.LogRequest()

	decoder := json.NewDecoder(r.Body)
	encoder := json.NewEncoder(w)

	var proxyrule models.ProxyRule
	routeVars := mux.Vars(r)
	err := decoder.Decode(&proxyrule)

	if err != nil {
		logger.Warn("Error decoding request body")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
		return
	}

	updatedProxyRule, err := h.Service.UpdateRule(routeVars["id"], proxyrule)

	if err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	encoder.Encode(updatedProxyRule)
}

func (h *ProxyRuleHandler) DeleteProxyRuleHandler(w http.ResponseWriter, r *http.Request) {
	logger := h.Logger.Request(r)
	logger.LogRequest()

	routeVars := mux.Vars(r)
	ruleId := routeVars["id"]

	if !h.Service.isRulePresent(ruleId) {
		logger.Warn("Rule not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err := h.Service.DeleteRule(routeVars["id"])

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
	}

	w.WriteHeader(http.StatusNoContent)
}
