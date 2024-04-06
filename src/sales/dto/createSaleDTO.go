package salesDTO

import (
	salesSubsDTO "github.com/chronicler-org/core/src/sales/dto/subs"
	saleEnum "github.com/chronicler-org/core/src/sales/enum"
)

type CreateSaleDTO struct {
	CustomerCareID string                     `validate:"required,uuid" json:"customer_care_id"`
	PaymentMethod  saleEnum.PaymentMethod     `validate:"required,paymentMethod" json:"payment_method"`
	SalesItems     []salesSubsDTO.SaleItemDTO `validate:"required,gt=0,dive" json:"sales_items"`
}
