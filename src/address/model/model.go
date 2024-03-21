package addressModel

import (
	"time"

	"github.com/google/uuid"
)

type Tag struct {
	ID uuid.UUID `gorm:primarykey`
	StreetName string `json:"streetName`
	Neighborhood string `json:"neighborhood"`
	Number string `json:"number"`
	City string `json:"city"`
	Estate string `json:"estate"`
	CEP string `json:"cep"`
	Complement string `json:"complement"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}