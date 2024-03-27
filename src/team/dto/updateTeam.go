package tagDTO

type UpdateTeamDTO struct {
	Name  string `validate:"omitempty,min=3,max=100" json:"name,omitempty"`
}
