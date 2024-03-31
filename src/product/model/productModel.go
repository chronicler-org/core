package productModel

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID        uuid.UUID `gorm:"primarykey"  json:"id"`
	Model     string    `gorm:"not null" json:"model"`
	Size      string    `gorm:"not null" json:"size"`
	Value     float32   `gorm:"not null" json:"value"`
	Fabric    string    `gorm:"not null" json:"fabric"`
	Stock     int64     `gorm:"not null" json:"stock"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
