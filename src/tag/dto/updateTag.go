package tagDTO

import "strings"

type UpdateTagDTO struct {
	Name  string `validate:"max=20" json:"name,omitempty"`
	Color string `validate:"max=20" json:"color,omitempty"`
}

func (dto *UpdateTagDTO) ValidateHexColor() bool {
	color := strings.ToLower(dto.Color)
	return hexRegex.MatchString(color)
}
