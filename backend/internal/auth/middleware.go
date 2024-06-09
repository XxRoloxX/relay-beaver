package auth

import (
	"backend/internal/logger"
	"net/http"
	"os"
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

func GetIdTokenFromCookie(r *http.Request) (string, error) {
	idToken, err := r.Cookie("id_token")

	if err != nil {
		return "", err
	}

	return idToken.Value, nil
}

func (authMiddleware AuthMiddleware) isM2MAuthenticated(r *http.Request) bool {
	authHeader := os.Getenv("PROXY_AUTH_HEADER")
	authSecret := os.Getenv("PROXY_AUTH_SECRET")
	return r.Header.Get(authHeader) == authSecret
}

func (authMiddleware AuthMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		logger := authMiddleware.Logger.Request(r)

		if authMiddleware.isM2MAuthenticated(r) {
			logger.Info("Authenticated with M2M")
			next.ServeHTTP(w, r)
			return
		}

		idToken, err := GetIdTokenFromCookie(r)

		if err != nil {
			logger.Error("Error getting id_token cookie")
			logger.Error(err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		authErr := authMiddleware.AuthService.ValidateToken(idToken)

		if authErr != nil {
			logger.Error("Error validating token")
			logger.Error(authErr.Error())
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
