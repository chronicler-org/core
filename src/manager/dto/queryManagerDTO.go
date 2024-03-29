package managerDTO

import appDto "github.com/chronicler-org/core/src/app/dto"

type QueryManagerDTO struct {
	appDto.PaginationDTO
	TeamID string `validate:"omitempty,uuid" query:"team_id"`
	Name   string `validate:"omitempty,uuid" query:"name"`
}
