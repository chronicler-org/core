package managerDTO

import (
	"time"
)

type UpdateManagerDTO struct {
	CPF       string    `json:"cpf,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"Password,omitempty"`
	BirthDate time.Time `json:"birth_date,omitempty"`
}

func (dto *UpdateManagerDTO) ValidateEmail() bool {
	return true
}

func (dto *UpdateManagerDTO) ValidateCPF() bool {
	return true
}
