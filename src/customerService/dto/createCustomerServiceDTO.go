package customerServiceDTO

import (
	"time"
)

type CreateCustomerServiceDTO struct {
	Date        time.Time `validate:"required" json:"date"`
	CustomerCPF string    `validate:"required,number,len=11" json:"customer_cpf"`
}
