package managerDTO

import (
	"time"
)

type CreateManagerDTO struct {
	CPF       string    `validate:"required,cpf" json:"cpf"`
	Name      string    `validate:"required,min=10,max=50" json:"name"`
	Email     string    `validate:"required,email,max=50" json:"email"`
	Password  string    `validate:"required,min=8" json:"password"`
	TeamId    string    `validate:"required,uuid" json:"team_id"`
	BirthDate time.Time `validate:"required" json:"birth_date"`
}
