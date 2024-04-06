package customerDTO

import appDto "github.com/chronicler-org/core/src/app/dto"

type QueryCustomerAddressDTO struct {
	appDto.PaginationDTO
	City       string `validate:"omitempty" query:"city" pg:"operation=like" `
	StreetName string `validate:"omitempty" query:"street_name" pg:"operation=like" `
	CEP        string `validate:"omitempty" query:"cep" pg:"operation=equal" `
}
