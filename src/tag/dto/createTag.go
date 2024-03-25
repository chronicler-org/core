package tagDTO

type CreateTagDTO struct {
	Name  string `validate:"required,min=3,max=20" json:"name"`
	Color string `validate:"required,len=7" json:"color"`
}
