package salesModel

import (
	"time"

	"github.com/google/uuid"

	customerCareModel "github.com/chronicler-org/core/src/customerCare/model"
)

type Sale struct {
	CustomerCareID uuid.UUID                      `gorm:"column:customer_care_id;primarykey" json:"id"`
	CustomerCare   customerCareModel.CustomerCare `gorm:"foreignKey:CustomerCareID" json:"customer_care"`
	TotalValue     float32                        `gorm:"not null" json:"total_value"`
	Status         string                         `gorm:"type:varchar(45);not null" json:"status"`
	PaymentMethod  string                         `gorm:"type:varchar(45);not null" json:"payment_method"`
	CreatedAt      time.Time                      `json:"created_at"`
	UpdatedAt      time.Time                      `json:"updated_at"`
}
