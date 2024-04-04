package salesDTO

type SaleItemDTO struct {
	ProductID string `validate:"required,uuid" json:"product_id"`
	Quantity  int    `validate:"required,min=1" json:"quantity"`
}
