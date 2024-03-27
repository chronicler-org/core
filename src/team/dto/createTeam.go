package tagDTO

type CreateTeamDTO struct {
	Name string `validate:"required,min=3,max=100" json:"name"`
}
