package auth

import (
	"backend/internal/auth/dto"
	"fmt"
	"time"
)

type AuthService struct{}

// func (service AuthService) GetAuthUrl(authCode string) string {
// 	googleClient := NewGoogleClientFromEnv()
//
// 	return googleClient.GetAuthUrl()
// }

func (service AuthService) Authenticate(code string) (dto.LoginResponseDto, error) {
	googleClient := NewGoogleClientFromEnv()

	return googleClient.Authenticate(code)
}

func (service AuthService) GetTokenInfo(idToken string) (dto.TokenInfoDto, error) {
	googleClient := NewGoogleClientFromEnv()

	return googleClient.GetTokenInfo(idToken)
}

func (service AuthService) ValidateToken(idToken string) error {
	token, err := service.GetTokenInfo(idToken)

	if err != nil {
		return err
	}

	if token.Expires < time.Now().Unix() {
		return fmt.Errorf("Token expired")
	}

	return nil
}
