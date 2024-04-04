package salesDTO

import (
	"github.com/google/uuid"
)

type SaleItemDTO struct {
	ProductID uuid.UUID `gorm:"column:product_id;primaryKey;not null" json:"product_id"`
	Quantity  int       `gorm:"not null" json:"quantity"`
}
