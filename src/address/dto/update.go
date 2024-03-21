package addressDTO

type UpdateAddressDTO struct {
	StreetName string `validate:"max=50" json:"streetName, omitempty"`
	Neighborhood  string  `validate:"max=30" json:"neighborhood, omitempty"`
	Number string `validate:"max=20" json:"number"`
	City string `validate:"max=30" json:"city, omitempty"`
	Estate string `validate:"max=30" json:"estate, omitempty"`
	CEP string `validate:"max=8" json:"cep"`
	Complement string `json:"complement, omitempty"`
}

