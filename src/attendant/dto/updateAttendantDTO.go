package attendantDTO

import (
	"time"
)

type UpdateAttendantDTO struct {
	CPF       string    `validate:"omitempty,number,len=11" json:"cpf,omitempty"`
	Name      string    `validate:"omitempty,min=10,max=50" json:"name,omitempty"`
	Email     string    `validate:"omitempty,email,max=50" json:"email,omitempty"`
	Password  string    `validate:"omitempty,min=8" json:"Password,omitempty"`
	BirthDate time.Time `json:"birth_date,omitempty"`
}
