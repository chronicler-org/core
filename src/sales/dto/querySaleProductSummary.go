package salesDTO

import (
	"github.com/google/uuid"

	appDto "github.com/chronicler-org/core/src/app/dto"
)

type QuerySalesProductSummaryDTO struct {
	appDto.PaginationDTO
	SaleID    uuid.UUID `validate:"omitempty,uuid" query:"sale_id"`
	ProductID uuid.UUID `validate:"omitempty,uuid" query:"product_id"`
}
