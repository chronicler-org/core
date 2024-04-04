package salesDTO

import "github.com/google/uuid"

type UpdateSaleDTO struct {
	CustomerCareID uuid.UUID `validate:"omitempty,uuid" json:"customer_care,omitempty"`
	TotalValue     int       `validate:"omitempty" json:"total_value,omitempty"`
	Status         string    `validate:"omitempty,lte=45" json:"status,omitempty"`
	PaymentMethod  string    `validate:"omitempty,lte=45" json:"payment_method,omitempty"`
}
