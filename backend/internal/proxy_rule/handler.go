package proxyrule

import (
	"backend/pkg/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ProxyRuleHandler struct {
	Service ProxyRuleService
}

func (h *ProxyRuleHandler) GetProxyRuleHandler(w http.ResponseWriter, r *http.Request) {
	requests, error := h.Service.Repo.FindAll()

	if error != nil {
		fmt.Println(error)
		w.WriteHeader(http.StatusInternalServerError)
	}

	serialized, error := json.Marshal(requests)

	w.WriteHeader(http.StatusOK)
	w.Write(serialized)
}

func (h *ProxyRuleHandler) CreateProxyRuleHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	encoder := json.NewEncoder(w)
	var proxyrule models.ProxyRule

	err := decoder.Decode(&proxyrule)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
		return
	}

	proxy, err := h.Service.CreateRule(proxyrule)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	encoder.Encode(proxy)
}

func (h *ProxyRuleHandler) UpdateProxyRuleHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	encoder := json.NewEncoder(w)
	var proxyrule models.ProxyRule

	err := decoder.Decode(&proxyrule)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
		return
	}

	updatedProxyRule, err := h.Service.UpdateRule(proxyrule)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	encoder.Encode(updatedProxyRule)
}
