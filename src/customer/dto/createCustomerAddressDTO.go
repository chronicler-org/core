package customerDTO

type CreateCustomerAddressDTO struct {
	CEP          string `validate:"required,len=9" json:"cep"`
	City         string `validate:"required,max=30" json:"city"`
	Number       string `validate:"required,max=20" json:"number"`
	Estate       string `validate:"required,len=2" json:"estate"`
	StreetName   string `validate:"required,max=50" json:"street_name"`
	Complement   string `validate:"omitempty,max=200" json:"complement"`
	Neighborhood string `validate:"omitempty,max=30" json:"neighborhood"`
}
