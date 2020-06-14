package clients

import (
	"fmt"

	resty "github.com/go-resty/resty/v2"
	"github.com/payment-link/config"
	"github.com/payment-link/contracts"
)

func Create(body string) contracts.LinkContract {
	url := config.GetBaseUrl() + "/api/public/v1/products/"

	var response contracts.LinkContract

	client := resty.New()
	_, err := client.R().
		SetAuthToken(GetToken()).
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetResult(&response).
		Post(url)

	if err != nil {
		fmt.Println("Failed to Create! Error:", err)
	}

	return response
}

func Update(linkId string, body string) contracts.LinkContract {
	url := config.GetBaseUrl() + "/api/public/v1/products/" + linkId

	var response contracts.LinkContract

	client := resty.New()
	_, err := client.R().
		SetAuthToken(GetToken()).
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetResult(&response).
		Put(url)

	if err != nil {
		fmt.Println("Failed to Update! Error:", err)
	}

	return response
}

func Delete(linkId string) {
	url := config.GetBaseUrl() + "/api/public/v1/products/" + linkId

	client := resty.New()
	_, err := client.R().
		SetAuthToken(GetToken()).
		SetHeader("Content-Type", "application/json").
		Delete(url)

	if err != nil {
		fmt.Println("Failed to Delete! Error:", err)
	}
}
