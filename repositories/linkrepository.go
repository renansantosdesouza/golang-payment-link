package repositories

import "github.com/payment-link/models"

func CreateLink(paymentLink models.LinkModel) {
	db := DBConnetion()
	defer db.Close()

	create, err := db.Prepare("insert into links (Id, Name, Description, Price, ShortUrl, BodyJson) values ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		panic(err.Error())
	}

	_, errCreate := create.Exec(paymentLink.Id, paymentLink.Name, paymentLink.Description, paymentLink.Price, paymentLink.ShortUrl, paymentLink.BodyJson)
	if errCreate != nil {
		panic(errCreate.Error())
	}
}

func DeleteLink(id string) {
	db := DBConnetion()

	delete, err := db.Prepare("delete from links where id=$1")
	if err != nil {
		panic(err.Error())
	}

	delete.Exec(id)
	defer db.Close()

}

func ReadLink(linkId string) models.LinkModel {
	db := DBConnetion()
	defer db.Close()

	read, err := db.Query("select Id, Name, Description, Price, ShortUrl, BodyJson from links where id=$1", linkId)
	if err != nil {
		panic(err.Error())
	}

	link := models.LinkModel{}

	for read.Next() {
		var id, name, description, shortUrl, bodyJson string
		var price int32

		err = read.Scan(&id, &name, &description, &price, &shortUrl, &bodyJson)
		if err != nil {
			panic(err.Error())
		}
		link.Id = id
		link.Name = name
		link.Description = description
		link.Price = price
		link.ShortUrl = shortUrl
		link.BodyJson = bodyJson
	}

	return link
}

func GetAll() []models.LinkModel {
	db := DBConnetion()
	defer db.Close()

	getList, err := db.Query("select Id, Name, Description, Price, ShortUrl, BodyJson from links")
	if err != nil {
		panic(err.Error())
	}

	link := models.LinkModel{}
	listLinks := []models.LinkModel{}

	for getList.Next() {
		var id, name, description, shortUrl, bodyJson string
		var price int32

		err = getList.Scan(&id, &name, &description, &price, &shortUrl, &bodyJson)

		if err != nil {
			panic(err.Error())
		}

		link.Id = id
		link.Name = name
		link.Description = description
		link.Price = price
		link.ShortUrl = shortUrl
		link.BodyJson = bodyJson

		listLinks = append(listLinks, link)
	}

	return listLinks
}

func UpdateLink(paymentLink models.LinkModel) {
	db := DBConnetion()
	defer db.Close()

	update, err := db.Prepare("update links set Name=$1, Description=$2, Price=$3, ShortUrl=$4, BodyJson=$5 where Id=$6")
	if err != nil {
		panic(err.Error())
	}

	update.Exec(paymentLink.Name, paymentLink.Description, paymentLink.Price, paymentLink.ShortUrl, paymentLink.BodyJson, paymentLink.Id)
}
