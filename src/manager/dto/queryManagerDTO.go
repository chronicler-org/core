package managerDTO

import appDto "github.com/chronicler-org/core/src/app/dto"

type QueryManagerDTO struct {
	appDto.PaginationDTO
	TeamID string `validate:"omitempty,uuid" query:"team_id" pg:"operation=equal"`
	Name   string `validate:"omitempty" query:"name" pg:"operation=like"`
	CPF    string `validate:"omitempty,cpf" query:"cpf" pg:"operation=like"`
}
