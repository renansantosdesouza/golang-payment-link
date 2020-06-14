package models

import (
	"encoding/json"
	"fmt"

	"github.com/payment-link/contracts"
)

type LinkModel struct {
	Id          string
	Name        string
	Description string
	Price       int32
	ShortUrl    string
	BodyJson    string
}

func (c *LinkModel) MapperFromContract(link contracts.LinkContract) {
	c.Id = link.Id
	c.Name = link.Name
	c.Description = link.Description
	c.Price = link.Price
	c.ShortUrl = link.ShortUrl

	linkJson, err := json.Marshal(link)
	if err != nil {
		fmt.Println(err)
	}
	c.BodyJson = string(linkJson)
}
