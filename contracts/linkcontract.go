package contracts

type LinkContract struct {
	Id                      string        `json:"id"`
	Name                    string        `json:"name"`
	Description             string        `json:"description"`
	Price                   int32         `json:"price"`
	ShortUrl                string        `json:"shortUrl"`
	Type                    string        `json:"type"`
	ShowDescription         bool          `json:"showDescription"`
	Weight                  int32         `json:"weight"`
	SoftDescriptor          string        `json:"softDescriptor"`
	ExpirationDate          string        `json:"expirationDate"`
	MaxNumberOfInstallments int32         `json:"maxNumberOfInstallments"`
	Shipping                ShippingModel `json:"thipping"`
}

type ShippingModel struct {
	Type          string `json:"type"`
	OriginZipcode string `json:"originZipcode"`
}
