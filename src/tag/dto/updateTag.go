package tagDTO

type UpdateTagDTO struct {
	Name  string `validate:"omitempty,min=3,max=20" json:"name,omitempty"`
	Color string `validate:"omitempty,len=7" json:"color,omitempty"`
}
