package salesDTO

import "github.com/google/uuid"

type CreateSaleDTO struct {
	CostumerCareID uuid.UUID `validate:"required,uuid" json:"customer_care"`
	TotalValue     int       `validate:"required," json:"total_value"`
	Status         string    `validate:"required,lte=45" json:"status"`
	PaymentMethod  string    `validate:"required,lte=45" json:"payment_method"`
}
