package loadbalancer

import (
	"backend/internal/common"
	"backend/internal/database"
	"backend/internal/logger"
	"backend/pkg/models"
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
)

type LoadBalancerRuleHandler struct {
	Service LoadBalancerService
	Logger  logger.HttpLogger
}

func NewLBHandler() LoadBalancerRuleHandler {
	handler := LoadBalancerRuleHandler{
		Service: LoadBalancerService{
			Repo: LoadBalancerMongoRepository{
				Db: *database.Db,
			}},
		Logger: logger.HttpLogger{},
	}
	err := handler.createMockLBs()
	if err != nil {
		log.Error().Msg(fmt.Sprintf("Mock LB's cannot be created, due to: %s", err.Error()))
	}
	return handler
}

func (h *LoadBalancerRuleHandler) createMockLBs() error {
	lbs := []models.LoadBalancer{{Name: "round robin"}, {Name: "least connections"}}

	for _, lb := range lbs {
		_, err := h.Service.CreateLB(lb)
		if err != nil {
			log.Error().Msg(fmt.Sprintf(fmt.Sprintf("Failed creating mock LB: %s", lb.Name)))
			return err
		}
	}

	return nil
}

func (h *LoadBalancerRuleHandler) GetLBHandler(w http.ResponseWriter, r *http.Request) {
	logger := h.Logger.Request(r)

	requests, error := h.Service.Repo.FindAll()

	if error != nil {
		common.LogBadRequestError(w, logger, error)
		return
	}

	if len(requests) == 0 {
		requests = make([]models.LoadBalancer, 0)
	}

	serialized, error := json.Marshal(requests)
	if error != nil {
		common.LogInternalServerError(w, logger, error)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(serialized)
}

func (h *LoadBalancerRuleHandler) CreateLBHandler(w http.ResponseWriter, r *http.Request) {
	logger := h.Logger.Request(r)
	var lb models.LoadBalancer
	err := json.NewDecoder(r.Body).Decode(&lb)

	if err != nil {
		common.LogBadRequestError(w, logger, err)
		return
	}

	proxy, err := h.Service.CreateLB(lb)

	if err != nil {
		common.LogInternalServerError(w, logger, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(proxy)
}

func (h *LoadBalancerRuleHandler) UpdateLBHandler(w http.ResponseWriter, r *http.Request) {
	logger := h.Logger.Request(r)

	decoder := json.NewDecoder(r.Body)
	encoder := json.NewEncoder(w)

	var lb models.LoadBalancer
	err := decoder.Decode(&lb)
	lbId := common.GetRouteParameter(r, "id")

	if err != nil {
		common.LogBadRequestError(w, logger, err)
		return
	}

	updatedProxyRule, err := h.Service.UpdateRule(lbId, lb)

	if err != nil {
		common.LogInternalServerError(w, logger, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	encoder.Encode(updatedProxyRule)
}

func (h *LoadBalancerRuleHandler) DeleteLBHandler(w http.ResponseWriter, r *http.Request) {
	logger := h.Logger.Request(r)

	lbId := common.GetRouteParameter(r, "id")

	if !h.Service.isLBPresent(lbId) {
		logger.Warn("LoadBalancer not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err := h.Service.DeleteLB(lbId)
	if err != nil {
		common.LogInternalServerError(w, logger, err)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
