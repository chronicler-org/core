package addressDTO

type CreateAddressDTO struct {
	CEP          string `validate:"required,max=8" json:"cep"`
	City         string `validate:"required,max=30" json:"city"`
	Number       string `validate:"required,max=20" json:"number"`
	Estate       string `validate:"required,max=30" json:"estate"`
	StreetName   string `validate:"required,max=50" json:"streetName"`
	Complement   string `json:"complement"`
	Neighborhood string `validate:"required,max=30" json:"neighborhood"`
}
