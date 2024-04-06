package salesDTO

type UpdateSaleDTO struct {
	Transition string `validate:"required,transition" json:"transition"`
}
