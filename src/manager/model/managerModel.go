package managerModel

import (
	"time"

	teamModel "github.com/chronicler-org/core/src/team/model"
	"github.com/google/uuid"
)

type Manager struct {
	ID        uuid.UUID      `gorm:"primarykey"  json:"id"`
	CPF       string         `gorm:"unique" json:"cpf"`
	Name      string         `json:"name"`
	Email     string         `gorm:"unique" json:"email"`
	Password  string         `json:"-"`
	Team      teamModel.Team `gorm:"not null, unique" json:"team"`
	BirthDate time.Time      `json:"birth_date"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}
