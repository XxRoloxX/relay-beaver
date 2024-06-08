package request

import (
	"backend/pkg/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RequestHandler struct {
	Service RequestService
}

func (h *RequestHandler) GetRequestsHandler(w http.ResponseWriter, r *http.Request) {
	requests, error := h.Service.Repo.FindAll()

	if error != nil {
		fmt.Println(error)
		w.WriteHeader(http.StatusInternalServerError)
	}

	serialized, error := json.Marshal(requests)

	w.WriteHeader(http.StatusOK)
	w.Write(serialized)
}

func (h *RequestHandler) PostRequestHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	encoder := json.NewEncoder(w)
	var request models.ProxiedRequest

	err := decoder.Decode(&request)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
		return
	}

	newRequest, err := h.Service.CreateRequest(request)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	encoder.Encode(newRequest)
}
