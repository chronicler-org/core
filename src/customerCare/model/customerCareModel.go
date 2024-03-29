package customerCareModel

import (
	"time"

	"github.com/google/uuid"

	customerModel "github.com/chronicler-org/core/src/customer/model"
	teamModel "github.com/chronicler-org/core/src/team/model"
)

type CustomerCare struct {
	ID          uuid.UUID              `gorm:"primarykey"  json:"id"`
	Date        time.Time              `json:"date"`
	CustomerCPF string                 `gorm:"column:customer_cpf;primaryKey;type:varchar(11);not null" json:"_"`
	Customer    customerModel.Customer `gorm:"foreignKey:CustomerCPF" json:"customer"`
	TeamID      uuid.UUID              `gorm:"column:team_id;not null" json:"-"`
	Team        teamModel.Team         `gorm:"foreignKey:TeamID" json:"team"`
	CreatedAt   time.Time              `json:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at"`
}
