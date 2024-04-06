package salesDTO

import (
	appDto "github.com/chronicler-org/core/src/app/dto"
)

type QuerySaleItemDTO struct {
	appDto.PaginationDTO
	SaleID       string `validate:"omitempty,uuid" json:"sale_id" query:"sale_id" `
	ProductID    string `validate:"omitempty,uuid" json:"product_id" query:"product_id" `
	CreatedMonth int    `validate:"omitempty,min=1,max=12" query:"created_month" pg:"operation=equal,date_extract=MONTH,name=sale_items.created_at" `
	CreatedYear  int    `validate:"omitempty,number" query:"created_year" pg:"operation=equal,date_extract=YEAR,name=sale_items.created_at" `
	Order        string `validate:"omitempty" json:"order" query:"order"`
}
