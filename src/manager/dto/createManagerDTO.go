package managerDTO

import (
	"time"
)

type CreateManagerDTO struct {
	CPF       string    `validate:"required" json:"cpf"`
	Name      string    `validate:"required,max=50" json:"name"`
	Email     string    `validate:"required,email,max=50" json:"email"`
	Password  string    `validate:"required,min=8" json:"password"`
	BirthDate time.Time `validate:"required" json:"birth_date"`
}
