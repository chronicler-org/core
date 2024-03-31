package productDTO

type UpdateProductDTO struct {
	Model  string  `validate:"omitempty,min=4,max=200,len=11" json:"model"`
	Size   string  `validate:"omitempty" json:"size"`
	Value  float32 `validate:"omitempty,number,min=0" json:"value"`
	Fabric string  `validate:"omitempty,min=2" json:"fabric"`
	Stock  int64   `validate:"omitempty,number,min=0" json:"stock"`
}
