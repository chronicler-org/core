package salesDTO

type UpdateSaleDTO struct {
	Transition string `validate:"required" json:"transition"`
}
