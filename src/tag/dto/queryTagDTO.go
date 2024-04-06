package tagDTO

import appDto "github.com/chronicler-org/core/src/app/dto"

type QueryTagDTO struct {
	appDto.PaginationDTO
	Name   string `validate:"omitempty" query:"name" pg:"operation=like"`
}
