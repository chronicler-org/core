package productModel

import (
	"time"

	"github.com/google/uuid"

	productEnum "github.com/chronicler-org/core/src/product/enum"
)

type Product struct {
	ID        uuid.UUID                 `gorm:"primarykey"  json:"id"`
	Model     productEnum.ClothingModel `gorm:"not null" json:"model"`
	Size      productEnum.Size          `gorm:"not null" json:"size"`
	Value     float32                   `gorm:"not null" json:"value"`
	Fabric    string                    `gorm:"not null" json:"fabric"`
	Stock     uint32                     `gorm:"not null" json:"stock"`
	CreatedAt time.Time                 `json:"created_at"`
	UpdatedAt time.Time                 `json:"updated_at"`
}
