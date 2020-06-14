package services

import (
	"encoding/json"
	"fmt"

	"github.com/payment-link/clients"
	"github.com/payment-link/models"
	repo "github.com/payment-link/repositories"
)

func Create(body string) bool {
	linkContractRespose := clients.Create(body)

	var linkModel models.LinkModel
	linkModel.MapperFromContract(linkContractRespose)

	repo.CreateLink(linkModel)
	return true
}

func Read(linkId string) models.LinkModel {
	link := repo.ReadLink(linkId)
	return link
}

func GetAll() []models.LinkModel {
	links := repo.GetAll()
	return links
}

func Update(linkId string, body string) bool {
	clients.Update(linkId, body)

	var linkModel models.LinkModel
	err := json.Unmarshal([]byte(body), &linkModel)
	if err != nil {
		fmt.Println(err)
	}

	linkModel.BodyJson = body
	repo.UpdateLink(linkModel)
	return true
}

func Delete(linkId string) bool {
	clients.Delete(linkId)
	repo.DeleteLink(linkId)
	return true
}
