package dto

type LoginResponseDto struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	IdToken      string `json:"id_token"`
}
type LoginRequestDto struct {
	Code string `json:"code"`
}

type TokenInfoDto struct {
	Email   string `json:"email"`
	Expires int64  `json:"expires"`
}
