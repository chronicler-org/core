package managerDTO

import (
	"net/mail"
	"time"

	"github.com/klassmann/cpfcnpj"
)

type UpdateManagerDTO struct {
	CPF       string    `json:"cpf,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"Password,omitempty"`
	BirthDate time.Time `json:"birth_date,omitempty"`
}

func (dto *UpdateManagerDTO) ValidateCPF() bool {
	return cpfcnpj.ValidateCPF(dto.CPF)
}

func (dto *UpdateManagerDTO) ValidateEmail() bool {
	_, err := mail.ParseAddress(dto.Email)
	return err == nil
}
