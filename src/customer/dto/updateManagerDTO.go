package customerDTO

import (
	"time"
)

type UpdateCustomerDTO struct {
	CPF       string    `validate:"required,max=11" json:"cpf,omitempty"`
	Name      string    `validate:"required,max=50" json:"name,omitempty"`
	Email     string    `validate:"required,max=50" json:"email,omitempty"`
	Phone     string    `validate:"required,max=11" json:"phone,omitempty"`
	Job       string    `validate:"required" json:"job,omitempty"`
	Password  string    `validate:"required,min=8" json:"password,omitempty"`
	BirthDate time.Time `validate:"required" json:"birth_date,omitempty"`
}

func (dto *UpdateCustomerDTO) ValidateEmail() bool {
	return true
}

func (dto *UpdateCustomerDTO) ValidateCPF() bool {
	return true
}
