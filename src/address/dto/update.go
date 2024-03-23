package addressDTO

type UpdateAddressDTO struct {
	CEP          string `validate:"max=8" json:"cep"`
	City         string `validate:"max=30" json:"city,omitempty"`
	Estate       string `validate:"max=30" json:"estate,omitempty"`
	Number       string `validate:"max=20" json:"number"`
	Complement   string `json:"complement,omitempty"`
	StreetName   string `validate:"max=50" json:"streetName,omitempty"`
	Neighborhood string `validate:"max=30" json:"neighborhood,omitempty"`
}
