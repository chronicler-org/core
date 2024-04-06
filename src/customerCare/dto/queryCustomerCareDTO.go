package customerCareDTO

import appDto "github.com/chronicler-org/core/src/app/dto"

type QueryCustomerCareDTO struct {
	appDto.PaginationDTO
	TeamID      string `validate:"omitempty,uuid" query:"team_id" pg:"operation=equal"`
	CustomerCPF string `validate:"omitempty,uuid" query:"customer_cpf" pg:"operation=like"`
}
