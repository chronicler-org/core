package customerServiceModel

import (
	"time"

	"github.com/google/uuid"

	customerModel "github.com/chronicler-org/core/src/customer/model"
	teamModel "github.com/chronicler-org/core/src/team/model"
)

type CustomerService struct {
	ID          uuid.UUID              `gorm:"primarykey"  json:"id"`
	Date        time.Time              `json:"date"`
	CustomerCPF uuid.UUID              `gorm:"column:customer_cpf;not null" json:"_"`
	Customer    customerModel.Customer `gorm:"foreignKey:CustomerCPF" json:"customer"`
	TeamID      uuid.UUID              `gorm:"column:team_id;not null" json:"-"`
	Team        teamModel.Team         `gorm:"foreignKey:TeamID" json:"team"`
	CreatedAt   time.Time              `json:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at"`
}
