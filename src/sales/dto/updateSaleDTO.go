package salesDTO

type UpdateSaleDTO struct {
	Status string `validate:"omitempty,lte=45" json:"status,omitempty"`
}
