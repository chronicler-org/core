package appDto

type PaginationDTO struct {
	Limit int `validate:"omitempty,min=1,max=100" json:"limit" query:"limit" `
	Page  int `validate:"omitempty,min=1" json:"page" query:"page" `
}

func (p *PaginationDTO) GetLimit() int {
	if p.Limit <= 0 {
		return 10
	}
	return p.Limit
}

func (p *PaginationDTO) GetPage() int {
	if p.Page <= 0 {
		return 1
	}
	return p.Page
}
