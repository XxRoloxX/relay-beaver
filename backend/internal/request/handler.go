package request

import (
	"encoding/json"
	"fmt"
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

func PostRequestHandler(w http.ResponseWriter, r *http.Request) {

}
