package attendantDTO

import (
	"time"
)

type CreateAttendantDTO struct {
	CPF       string    `validate:"required,number,len=11" json:"cpf"`
	Name      string    `validate:"required,min=10,max=50" json:"name"`
	Email     string    `validate:"required,email,max=50" json:"email"`
	Password  string    `validate:"required,min=8" json:"password"`
	BirthDate time.Time `validate:"required" json:"birth_date"`
}
