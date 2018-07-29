package auth

import (
	"github.com/dghubble/sling"
	"github.com/spf13/viper"
	"net/http"
)

type AuthenticationService struct {
	sling *sling.Sling
}

func NewAuthenticationService(client *http.Client) *AuthenticationService {
	return &AuthenticationService{
		sling: sling.New().Client(client).Base(viper.GetString("services.auth.addr")),
	}
}

func (auth *AuthenticationService) SetBase(url string) *AuthenticationService {
	auth.SetBase(url)
	return auth
}

func (auth *AuthenticationService) Account(token string) (*AccountApiResponse, *http.Response, error) {
	var resp = &AccountApiResponse{}
	accSling := auth.sling.New().Get("/api/account").QueryStruct(&AuthRequestQuery{token})
	httpResp, err := accSling.ReceiveSuccess(resp)
	return resp, httpResp, err
}

type AccountApiResponse struct {
	Id    string `json:"user_uuid,omitempty"`
	Email string `json:"email,omitempty"`
}

type AuthRequestQuery struct {
	AccessToken string `url:"access_token"`
}
