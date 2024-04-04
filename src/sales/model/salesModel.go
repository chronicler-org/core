package salesModel

import (
	"time"

	customerCareModel "github.com/chronicler-org/core/src/customerCare/model"
	"github.com/google/uuid"
)

type Sales struct {
	CustomerCare   customerCareModel.CustomerCare `gorm:"foreignKey:CustomerCareID" json:"customer_care"`
	CustomerCareID uuid.UUID                      `gorm:"column:customer_care_id;primarykey" json:"id"`
	TotalValue     int                            `json:"total_value"`
	Status         string                         `gorm:"type:varchar(45)" json:"status"`
	PaymentMethod  string                         `gorm:"type:varchar(45)" json:"payment_method"`
	CreatedAt      time.Time                      `json:"created_at"`
	UpdatedAt      time.Time                      `json:"updated_at"`
}
