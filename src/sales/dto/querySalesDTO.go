package salesDTO

import (
	appDto "github.com/chronicler-org/core/src/app/dto"
)

type QuerySalesDTO struct {
	appDto.PaginationDTO
	CustomerCareID string `validate:"omitempty,uuid" query:"customer_care_id" pg:"operation=equal"`
	Status  string `validate:"omitempty,status" query:"status" pg:"operation=equal"`
	PaymentMethod string `validate:"omitempty,paymentMethod" query:"payment_method" pg:"operation=equal"`
	Order  string `validate:"omitempty" json:"order" query:"order"`
}
