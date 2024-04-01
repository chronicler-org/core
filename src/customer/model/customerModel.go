package customerModel

import (
	"time"

	tagModel "github.com/chronicler-org/core/src/tag/model"
	"github.com/google/uuid"
)

type Customer struct {
	CPF       string          `gorm:"primaryKey;type:varchar(11)" json:"cpf"`
	Name      string          `json:"name"`
	Email     string          `gorm:"uniqueIndex" json:"email"`
	Phone     string          `json:"phone"`
	Job       string          `json:"job"`
	BirthDate time.Time       `json:"birth_date"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	Tags      []*tagModel.Tag `gorm:"many2many:customer_tags;onDelete:CASCADE" json:"tags"`
	AddressID uuid.UUID       `gorm:"column:address_id;not null" json:"-"`
	Address   CustomerAddress `gorm:"foreignKey:AddressID" json:"address"`
}
