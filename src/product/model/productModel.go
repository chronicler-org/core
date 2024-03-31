package productModel

import (
	"database/sql/driver"
	"time"

	"github.com/google/uuid"

	appUtil "github.com/chronicler-org/core/src/app/utils"
)

type Size string

const (
	SizePP Size = "PP"
	SizeP  Size = "P"
	SizeM  Size = "M"
	SizeG  Size = "G"
	SizeGG Size = "GG"
)

type ClothingModel string

const (
	TShirt     ClothingModel = "Camiseta"
	Shirt      ClothingModel = "Camisa"
	Pants      ClothingModel = "Calça"
	Skirt      ClothingModel = "Saia"
	Dress      ClothingModel = "Vestido"
	Jacket     ClothingModel = "Casaco"
	Sportswear ClothingModel = "Roupas Esportivas"
	Underwear  ClothingModel = "Roupa Íntima"
	Swimwear   ClothingModel = "Roupa de Banho"
	FormalWear ClothingModel = "Roupa Formal"
)

type Product struct {
	ID        uuid.UUID     `gorm:"primarykey"  json:"id"`
	Model     ClothingModel `gorm:"not null" json:"model"`
	Size      Size          `gorm:"not null" json:"size"`
	Value     float32       `gorm:"not null" json:"value"`
	Fabric    string        `gorm:"not null" json:"fabric"`
	Stock     int64         `gorm:"not null" json:"stock"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

func (s *Size) Scan(value interface{}) error {
	size, ok := value.([]byte)
	if !ok {
		validSizes := []interface{}{SizePP, SizeP, SizeM, SizeG, SizeGG}
		return appUtil.GenerateEnumInvalidFieldError("Size", validSizes)
	}

	*s = Size(string(size))
	return nil
}

func (s Size) Value() (driver.Value, error) {
	return string(s), nil
}

func (m *ClothingModel) Scan(value interface{}) error {
	model, ok := value.([]byte)
	if !ok {
		validModels := []interface{}{TShirt, Shirt, Pants, Skirt, Dress, Jacket, Sportswear, Underwear, Swimwear, FormalWear}
		return appUtil.GenerateEnumInvalidFieldError("Model", validModels)
	}

	*m = ClothingModel(string(model))
	return nil
}

func (s ClothingModel) Value() (driver.Value, error) {
	return string(s), nil
}
