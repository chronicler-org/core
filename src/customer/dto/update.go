package customerDTO

import (
	"time"
)

type UpdateCustomerDTO struct {
	Name      string    `validate:"max=50" json:"name,omitempty"`
	Email     string    `validate:"max=50" json:"email,omitempty"`
	Phone     string    `validate:"max=11" json:"phone,omitempty"`
	Job       string    `json:"job,omitempty"`
	BirthDate time.Time `json:"birth_date,omitempty"`
}

func (dto *UpdateCustomerDTO) ValidateEmail() bool {
	return true
}
