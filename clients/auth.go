package clients

import (
	"fmt"

	resty "github.com/go-resty/resty/v2"
	"github.com/payment-link/config"
)

type AccessToken struct {
	Token   string `json:"access_token"`
	Type    string `json:"token_type"`
	Expires string `json:"expires_in"`
}

func GetToken() string {
	url := config.GetBaseUrl() + "/api/public/v2/token"

	var response AccessToken

	client := resty.New()
	_, err := client.R().
		SetBasicAuth(config.GetClient(), config.GetSecret()).
		SetBody("grant_type=client_credentials").
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetResult(&response).
		Post(url)

	if err != nil {
		fmt.Println("Failed to get Authorization! Error:", err)
	}

	return response.Token
}
