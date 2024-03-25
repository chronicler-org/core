package managerModel

import (
	"time"

	"github.com/google/uuid"

	tagModel "github.com/chronicler-org/core/src/tag/model"
)

type Manager struct {
	ID        uuid.UUID       `gorm:"primarykey"`
	CPF       string          `gorm:"uniqueIndex" json:"cpf"`
	Name      string          `json:"name"`
	Email     string          `gorm:"uniqueIndex" json:"email"`
	Password  string          `json:"-"`
	BirthDate time.Time       `json:"birth_date"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	Tags      []*tagModel.Tag `gorm:"many2many:manager_tags" json:"tags"`
}
