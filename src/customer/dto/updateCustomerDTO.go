package customerDTO

import (
	"time"
)

type UpdateCustomerDTO struct {
	CPF       string    `validate:"omitempty,number,len=11" json:"cpf,omitempty"`
	Name      string    `validate:"omitempty,min=10,max=50" json:"name,omitempty"`
	Email     string    `validate:"omitempty,email,max=50"  json:"email,omitempty"`
	Phone     string    `validate:"omitempty,number,len=12" json:"phone,omitempty"`
	Job       string    `validate:"omitempty,min=10,max=30" json:"job,omitempty"`
	BirthDate time.Time `validate:"omitempty" json:"birth_date,omitempty"`
}
