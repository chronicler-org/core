package customerCareModel

import (
	"time"

	"github.com/google/uuid"

	customerModel "github.com/chronicler-org/core/src/customer/model"
)

type CustomerCareEvaluation struct {
	ID             uuid.UUID              `gorm:"primarykey"  json:"id"`
	Score          uint8                  `gorm:"not null" json:"score"`
	Description    string                 `gorm:"not null" json:"description"`
	CustomerCareID uuid.UUID              `gorm:"type:uuid;column:customer_care_id;not null" json:"-"`
	CustomerCare   CustomerCare           `gorm:"foreignKey:ID;references:CustomerCareID" json:"customer_care"`
	CustomerCPF    string                 `gorm:"column:customer_cpf;primaryKey;type:varchar(11);not null" json:"-"`
	Customer       customerModel.Customer `gorm:"foreignKey:CustomerCPF" json:"customer"`
	CreatedAt      time.Time              `json:"created_at"`
	UpdatedAt      time.Time              `json:"updated_at"`
}
