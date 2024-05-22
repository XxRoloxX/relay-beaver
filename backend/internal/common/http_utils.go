package common

import (
	"net/http"

	"github.com/gorilla/mux"
)

func GetRouteParameter(r *http.Request, parameter string) string {
	return mux.Vars(r)[parameter]
}
