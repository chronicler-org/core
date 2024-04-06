package productDTO

type CreateProductDTO struct {
	Model  string  `validate:"required,model" json:"model"`
	Size   string  `validate:"required,size" json:"size"`
	Value  float32 `validate:"required,number,min=0" json:"value"`
	Fabric string  `validate:"required,min=2" json:"fabric"`
	Stock  uint32   `validate:"required,number,min=0" json:"stock"`
}
