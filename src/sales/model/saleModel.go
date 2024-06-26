package salesModel

import (
	"time"

	"github.com/google/uuid"

	customerCareModel "github.com/chronicler-org/core/src/customerCare/model"
	saleEnum "github.com/chronicler-org/core/src/sales/enum"
)

type Sale struct {
	CustomerCareID uuid.UUID                      `gorm:"column:customer_care_id;primarykey" json:"id"`
	CustomerCare   customerCareModel.CustomerCare `gorm:"foreignKey:CustomerCareID" json:"customer_care"`
	TotalValue     float32                        `gorm:"type:decimal(10,2);not null" json:"total_value"`
	Status         saleEnum.SaleStatus            `gorm:"type:varchar(45);not null" json:"status"`
	PaymentMethod  saleEnum.PaymentMethod         `gorm:"not null" json:"payment_method"`
	CreatedAt      time.Time                      `json:"created_at"`
	UpdatedAt      time.Time                      `json:"updated_at"`
}
