package auth

import (
	"backend/internal/auth/dto"
	"context"
	"google.golang.org/api/idtoken"
	"os"
)

type AuthService struct{}

func (service AuthService) GetAuthUrl(authCode string) string {
	googleClient := NewGoogleClientFromEnv()

	return googleClient.GetAuthUrl()
}

func (service AuthService) Authenticate(code string) (dto.LoginResponseDto, error) {
	googleClient := NewGoogleClientFromEnv()

	return googleClient.Authenticate(code)
}
func (service AuthService) ValidateToken(idToken string) error {
	_, err := idtoken.Validate(context.Background(), idToken, os.Getenv("GOOGLE_CLIENT_ID"))
	return err
}
