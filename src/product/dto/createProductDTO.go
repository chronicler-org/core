package productDTO

type CreateProductDTO struct {
	Model  string  `validate:"required,min=4,max=200,len=11" json:"model"`
	Size   string  `validate:"required" json:"size"`
	Value  float32 `validate:"required,number,min=0" json:"value"`
	Fabric string  `validate:"required,min=2" json:"fabric"`
	Stock  int64   `validate:"required,number,min=0" json:"stock"`
}
