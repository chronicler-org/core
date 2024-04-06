package salesDTO

type QuerySaleItemsDTO struct {
	SaleID string `validate:"required,uuid" json:"sale_id" query:"sale_id" pg:"operation=equal"`
}
