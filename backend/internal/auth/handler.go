package auth

import (
	"backend/internal/auth/dto"
	"backend/internal/logger"
	"encoding/json"
	"io"
	"net/http"
)

type AuthHandler struct {
	AuthService AuthService
	logger      logger.HttpLogger
}

func NewAuthHandler() AuthHandler {
	return AuthHandler{
		AuthService: AuthService{},
		logger:      logger.HttpLogger{},
	}
}

//
// func (handler AuthHandler) AuthCodeHandler(w http.ResponseWriter, request *http.Request) {
// 	logger := handler.logger.Request(request)
// 	logger.LogRequest()
//
// 	io.WriteString(w, handler.AuthService.GetAuthUrl(""))
// }

func (handler AuthHandler) LoginHandler(w http.ResponseWriter, request *http.Request) {
	logger := handler.logger.Request(request)
	logger.LogRequest()

	decoder := json.NewDecoder(request.Body)
	var loginRequestDto dto.LoginRequestDto
	err := decoder.Decode(&loginRequestDto)

	if err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Invalid request")
		return
	}

	response, err := handler.AuthService.Authenticate(loginRequestDto.Code)

	if err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, "Error authenticating")
		return
	}

	SetAuthCookies(&w, response)
}

func (handler AuthHandler) TokenInfoHandler(w http.ResponseWriter, request *http.Request) {
	logger := handler.logger.Request(request)

	idToken, err := GetIdTokenFromCookie(request)
	if err != nil {
		logger.Error("Error getting id_token cookie")
		logger.Error(err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = handler.AuthService.ValidateToken(idToken)
	if err != nil {
		logger.Error("Invalid Token")
		logger.Error(err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	tokenInfo, err := handler.AuthService.GetTokenInfo(idToken)
	if err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, "Error getting token info")
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(tokenInfo)
	if err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

}

func SetAuthCookies(w *http.ResponseWriter, response dto.LoginResponseDto) {
	http.SetCookie(*w, &http.Cookie{
		Name:     "id_token",
		Value:    response.IdToken,
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	})

	http.SetCookie(*w, &http.Cookie{
		Name:     "access_token",
		Value:    response.AccessToken,
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	})

	http.SetCookie(*w, &http.Cookie{
		Name:     "refresh_token",
		Value:    response.RefreshToken,
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	})
}
