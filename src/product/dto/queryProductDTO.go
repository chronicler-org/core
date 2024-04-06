package productDTO

import appDto "github.com/chronicler-org/core/src/app/dto"

type QueryProductDTO struct {
	appDto.PaginationDTO
	Model  string `validate:"omitempty,model" query:"model" pg:"operation=like"`
	Fabric string `validate:"omitempty" query:"fabric" pg:"operation=like"`
	Size   string `validate:"omitempty" query:"size" pg:"operation=equal"`
}
