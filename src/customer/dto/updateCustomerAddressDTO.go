package customerDTO

type UpdateCustomerAddressDTO struct {
	CEP          string `validate:"omitempty,len=9" json:"cep"`
	City         string `validate:"omitempty,max=30" json:"city"`
	Estate       string `validate:"omitempty,len=2" json:"estate"`
	Number       string `validate:"omitempty,max=20" json:"number"`
	Complement   string `validate:"omitempty,max=200" json:"complement"`
	StreetName   string `validate:"omitempty,max=50" json:"street_name"`
	Neighborhood string `validate:"omitempty,max=30" json:"neighborhood"`
}
