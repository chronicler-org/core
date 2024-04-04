package salesDTO

import "github.com/google/uuid"

type CreateSaleDTO struct {
	CostumerCareID uuid.UUID     `validate:"required,uuid" json:"customer_care_id"`
	TotalValue     int           `validate:"required" json:"total_value"`
	Status         string        `validate:"required,max=45" json:"status"`
	PaymentMethod  string        `validate:"required,max=45" json:"payment_method"`
	SalesItems     []SaleItemDTO `validate:"required,dive" json:"sales_items"`
}
