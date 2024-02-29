package customerDTO

import (
	"time"
)

type CreateCustomerDTO struct {
	CPF       string    `validate:"required,max=11" json:"cpf"`
	Name      string    `validate:"required,max=50" json:"name"`
	Email     string    `validate:"required,max=50" json:"email"`
	Phone     string    `validate:"required,max=11" json:"phone"`
	Job       string    `validate:"required" json:"job"`
	BirthDate time.Time `validate:"required" json:"birth_date"`
}

func (dto *CreateCustomerDTO) ValidateEmail() bool {
	return true
}

func (dto *CreateCustomerDTO) ValidateCPF() bool {
	return true
}
