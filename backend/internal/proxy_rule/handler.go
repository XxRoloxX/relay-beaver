package proxyrule

import (
	"backend/internal/common"
	"backend/internal/database"
	"backend/internal/logger"
	"backend/pkg/models"
	"encoding/json"
	"io"
	"net/http"
)

type ProxyRuleHandler struct {
	Service ProxyRuleService
	Logger  logger.HttpLogger
}

func NewProxyRuleHandler() ProxyRuleHandler {
	return ProxyRuleHandler{
		Service: ProxyRuleService{
			Repo: ProxyRuleMongoRepository{
				Db: *database.Db,
			}},
		Logger: logger.HttpLogger{},
	}
}

func (h *ProxyRuleHandler) GetProxyRuleHandler(w http.ResponseWriter, r *http.Request) {
	logger := h.Logger.Request(r)

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

	decoder := json.NewDecoder(r.Body)
	encoder := json.NewEncoder(w)

	var proxyrule models.ProxyRule
	err := decoder.Decode(&proxyrule)
	ruleId := common.GetRouteParameter(r, "id")

	if err != nil {
		logger.Warn("Error decoding request body")
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
		return
	}

	updatedProxyRule, err := h.Service.UpdateRule(ruleId, proxyrule)

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

	ruleId := common.GetRouteParameter(r, "id")

	if !h.Service.isRulePresent(ruleId) {
		logger.Warn("Rule not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err := h.Service.DeleteRule(ruleId)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
	}

	w.WriteHeader(http.StatusNoContent)
}
