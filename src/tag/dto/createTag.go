package tagDTO

import "strings"

type CreateTagDTO struct {
	Name string `validate:"required,max=20" json:"name"`
	Color string `validate:"required,max=20" json:"color"`
}

func (dto *CreateTagDTO) ValidateHexColor() bool {
	color := strings.ToLower(dto.Color)
	return hexRegex.MatchString(color)
}
