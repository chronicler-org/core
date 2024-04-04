package salesDTO

type CreateSaleDTO struct {
	CustomerCareID string        `validate:"required,uuid" json:"customer_care_id"`
	Status         string        `validate:"required,max=45,status" json:"status"`
	PaymentMethod  string        `validate:"required,max=45" json:"payment_method"`
	SalesItems     []SaleItemDTO `validate:"required,gt=0,dive" json:"sales_items"`
}
