package auth

import (
	"backend/internal/auth/dto"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"google.golang.org/api/idtoken"
)

const GOOGLE_AUTH_URL = "https://accounts.google.com/o/oauth2"

type GoogleClient struct {
	ClientId     string
	ClientSecret string
	RedirectUrl  string
	Scope        string
}

func (client GoogleClient) SetClientId(clientId string) GoogleClient {
	client.ClientId = clientId
	return client
}

func (client GoogleClient) SetClientSecret(clientSecret string) GoogleClient {
	client.ClientSecret = clientSecret
	return client
}

func (client GoogleClient) SetRedirectUrl(redirectUrl string) GoogleClient {
	client.RedirectUrl = redirectUrl
	return client
}

func (client GoogleClient) SetScope(scope string) GoogleClient {
	client.Scope = scope
	return client
}

func NewGoogleClientFromEnv() GoogleClient {
	return GoogleClient{
		ClientId:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectUrl:  os.Getenv("GOOGLE_REDIRECT_URI"),
		Scope:        "email",
	}
}

func (client GoogleClient) GetAccessTokenUrl(code string) string {
	return fmt.Sprintf("%s/token?client_id=%s&client_secret=%s&redirect_uri=%s&grant_type=authorization_code&code=%s",
		GOOGLE_AUTH_URL,
		client.ClientId,
		client.ClientSecret,
		client.RedirectUrl,
		code)
}

func (client GoogleClient) GetTokenInfo(idToken string) (dto.TokenInfoDto, error) {
	token, err := idtoken.Validate(context.Background(), idToken, os.Getenv("GOOGLE_CLIENT_ID"))
	if err != nil {
		return dto.TokenInfoDto{}, err
	}

	return dto.TokenInfoDto{
		Email:   token.Claims["email"].(string),
		Expires: token.Expires,
	}, nil
}

func (client GoogleClient) Authenticate(code string) (dto.LoginResponseDto, error) {
	res, err := http.NewRequest("POST", client.GetAccessTokenUrl(code), nil)
	if err != nil {
		return dto.LoginResponseDto{}, err
	}

	googleClient := &http.Client{
		Timeout: time.Second * 10,
	}

	googleResponse, err := googleClient.Do(res)
	if err != nil {
		return dto.LoginResponseDto{}, err
	}

	var response dto.LoginResponseDto
	defer googleResponse.Body.Close()
	body, err := io.ReadAll(googleResponse.Body)

	if googleResponse.StatusCode != 200 {
		return dto.LoginResponseDto{}, fmt.Errorf("Error authenticating with Google %s", string(body))
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return dto.LoginResponseDto{}, err
	}

	return response, nil
}
