package customerModel

import (
	"time"

	"github.com/google/uuid"
)

type CustomerAddress struct {
	ID 						uuid.UUID	`gorm:"primarykey"`
	CEP 					string		`json:"cep"`
	City 					string 		`json:"city"`
	Number 				string		`json:"number"`
	Estate 				string		`json:"estate"`
	StreetName 		string		`json:"street_name"`
	Complement 		string		`json:"complement"`
	Neighborhood	string 		`json:"neighborhood"`
	CreatedAt			time.Time	`json:"created_at"`
	UpdatedAt 		time.Time `json:"updated_at"`
}