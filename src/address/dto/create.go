package addressDTO

type CreateAddressDTO struct {
	StreetName string `validate:"required, max=50" json:"streetName"`
	Neighborhood  string  `validate:"required, max=30" json:"neighborhood"`
	Number string `validate:"required, max=20" json:"number"`
	City string `validate:"required, max=30" json:"city"`
	Estate string `validate:"required, max=30" json:"estate"`
	CEP string `validate:"required, max=8" json:"cep"`
	Complement string `json:"complement"`
}

