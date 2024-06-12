package common

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleCors(
	router *mux.Router,
) http.Handler {

	origins := handlers.AllowedOrigins([]string{os.Getenv("FRONTEND_URL")})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	headers := handlers.AllowedHeaders([]string{
		"X-Requested-With",
		"Content-Type",
		"Authorization",
	})

	return handlers.CORS(origins, headers, methods, handlers.AllowCredentials())(router)
}
