package auth

import (
	"backend/internal/logger"
	"net/http"
)

type AuthMiddleware struct {
	AuthService AuthService
	Logger      logger.HttpLogger
}

func NewAuthMiddleware() AuthMiddleware {
	return AuthMiddleware{
		AuthService: AuthService{},
		Logger:      logger.HttpLogger{},
	}
}

func (authMiddleware AuthMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		logger := authMiddleware.Logger.Request(r)

		// println("Checking for id_token cookie")

		// for name, header := range r.Header {
		// 	for _, value := range header {
		// 		println(name)
		// 		println(value)
		// 	}
		// }
		//
		// for _, cookie := range r.Cookies() {
		// 	println(cookie.Name)
		// }

		idToken, err := r.Cookie("id_token")

		if err != nil {
			logger.Error("Error getting id_token cookie")
			logger.Error(err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		authErr := authMiddleware.AuthService.ValidateToken(idToken.Value)

		if authErr != nil {
			logger.Error("Error validating token")
			logger.Error(authErr.Error())
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		println("Token is valid")

		next.ServeHTTP(w, r)
	})
}
