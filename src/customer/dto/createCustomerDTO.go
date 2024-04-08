package customerDTO

type CreateCustomerDTO struct {
	CPF       string   `validate:"required,cpf" json:"cpf"`
	Name      string   `validate:"required,min=10,max=50" json:"name"`
	Email     string   `validate:"required,email,max=50"  json:"email"`
	Phone     string   `validate:"required,number,len=11" json:"phone"`
	Job       string   `validate:"required,min=5,max=30" json:"job"`
	BirthDate string   `validate:"required,birthdate,notfuturedate" json:"birth_date"`
	TagIDs    []string `validate:"omitempty,dive,uuid" json:"tag_ids"`
	AddressId string   `validate:"required,uuid" json:"address_id"`
}
