package proxyrule

import (
	"backend/internal/common"
	"backend/internal/database"
	"backend/internal/logger"
	"backend/pkg/models"
	"encoding/json"
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
		common.LogBadRequestError(w, logger, error)
		return
	}

	_, error = json.Marshal(requests)

	if error != nil {
		common.LogInternalServerError(w, logger, error)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("0011311003344##323333HHggHelloee2233444, Wworld! 234443333322444"))
}

func (h *ProxyRuleHandler) CreateProxyRuleHandler(w http.ResponseWriter, r *http.Request) {
	logger := h.Logger.Request(r)
	var proxyrule models.ProxyRule
	err := json.NewDecoder(r.Body).Decode(&proxyrule)

	if err != nil {
		common.LogBadRequestError(w, logger, err)
		return
	}

	proxy, err := h.Service.CreateRule(proxyrule)

	if err != nil {
		common.LogInternalServerError(w, logger, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(proxy)
}

func (h *ProxyRuleHandler) UpdateProxyRuleHandler(w http.ResponseWriter, r *http.Request) {
	logger := h.Logger.Request(r)

	decoder := json.NewDecoder(r.Body)
	encoder := json.NewEncoder(w)

	var proxyrule models.ProxyRule
	err := decoder.Decode(&proxyrule)
	ruleId := common.GetRouteParameter(r, "id")

	if err != nil {
		common.LogBadRequestError(w, logger, err)
		return
	}

	updatedProxyRule, err := h.Service.UpdateRule(ruleId, proxyrule)

	if err != nil {
		common.LogInternalServerError(w, logger, err)
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
		common.LogInternalServerError(w, logger, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
