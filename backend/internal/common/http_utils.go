package common

import (
	"backend/internal/logger"
	"github.com/gorilla/mux"
	"net/http"
)

func GetRouteParameter(r *http.Request, parameter string) string {
	return mux.Vars(r)[parameter]
}

func LogBadRequestError(w http.ResponseWriter, logger logger.RequestLogger, err error) {
	w.WriteHeader(http.StatusBadRequest)
	logger.Warn(err.Error())
	w.Write([]byte(err.Error()))
}

func LogInternalServerError(w http.ResponseWriter, logger logger.RequestLogger, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	logger.Error(err.Error())
}
