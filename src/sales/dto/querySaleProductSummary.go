package salesDTO

import (
	appDto "github.com/chronicler-org/core/src/app/dto"
)

type QuerySalesProductSummaryDTO struct {
	appDto.PaginationDTO
	SaleID    string `validate:"omitempty,uuid" query:"sale_id"`
	ProductID string `validate:"omitempty,uuid" query:"product_id"`
}
