package customerModel

import (
	"time"

	"github.com/google/uuid"
)

type CustomerAddress struct {
	ID           uuid.UUID `gorm:"primarykey" json:"id"`
	CEP          string    `gorm:"not null" json:"cep"`
	City         string    `gorm:"not null" json:"city"`
	Number       string    `gorm:"not null" json:"number"`
	Estate       string    `gorm:"not null" json:"estate"`
	StreetName   string    `json:"street_name"`
	Complement   string    `json:"complement"`
	Neighborhood string    `json:"neighborhood"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
