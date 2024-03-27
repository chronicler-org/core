package tagModel

import (
	"time"

	"github.com/google/uuid"
)

type Team struct {
	ID        uuid.UUID `gorm:"primarykey"  json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}