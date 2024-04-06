package salesDTO

import salesSubsDTO "github.com/chronicler-org/core/src/sales/dto/subs"

type CreateSaleDTO struct {
	CustomerCareID string                     `validate:"required,uuid" json:"customer_care_id"`
	PaymentMethod  string                     `validate:"required,max=45" json:"payment_method"`
	SalesItems     []salesSubsDTO.SaleItemDTO `validate:"required,gt=0,dive" json:"sales_items"`
}
