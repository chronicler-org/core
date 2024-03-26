package tagModel

import (
	"time"

	"github.com/google/uuid"
)

type Tag struct {
	ID        uuid.UUID `gorm:"primarykey"  json:"id"`
	Name      string    `gorm:"uniqueIndex" json:"name"`
	Color     string    `json:"color"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
