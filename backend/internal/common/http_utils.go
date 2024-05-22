package common

import (
	"github.com/gorilla/mux"
	"net/http"
)

func GetRouteParameter(r *http.Request, parameter string) string {
	return mux.Vars(r)[parameter]
}
