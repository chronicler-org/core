package customerCareDTO

import appDto "github.com/chronicler-org/core/src/app/dto"

type QueryCustomerCareDTO struct {
	appDto.PaginationDTO
	TeamID      string `validate:"omitempty,uuid" query:"team_id"`
	CustomerCPF string `validate:"omitempty,uuid" query:"customer_cpf"`
}
