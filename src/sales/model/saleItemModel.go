package salesModel

import (
	"time"

	"github.com/google/uuid"

	productModel "github.com/chronicler-org/core/src/product/model"
)

type SaleItem struct {
	SaleID    uuid.UUID            `gorm:"column:sale_id;primaryKey;not null" json:"-"`
	Sale      Sale                 `gorm:"foreignKey:SaleID;constraint:OnDelete:CASCADE" json:"sale"`
	ProductID uuid.UUID            `gorm:"column:product_id;primaryKey;not null" json:"-"`
	Product   productModel.Product `gorm:"foreignKey:ProductID" json:"product"`
	Quantity  int                  `gorm:"not null" json:"quantity"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
}
