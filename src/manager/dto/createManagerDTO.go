package managerDTO

import (
	"time"

	"github.com/klassmann/cpfcnpj"
)

type CreateManagerDTO struct {
	CPF       string    `validate:"required,max=11" json:"cpf"`
	Name      string    `validate:"required,max=50" json:"name"`
	Email     string    `validate:"required,max=50" json:"email"`
	Password  string    `validate:"required,min=8" json:"password"`
	BirthDate time.Time `validate:"required" json:"birth_date"`
}

func (dto *CreateManagerDTO) Validate() bool {
	return cpfcnpj.ValidateCPF(dto.CPF)
}
