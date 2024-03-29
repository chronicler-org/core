package managerModel

import (
	"time"

	"github.com/google/uuid"

	teamModel "github.com/chronicler-org/core/src/team/model"
)

type Manager struct {
	ID        uuid.UUID      `gorm:"primarykey"  json:"id"`
	CPF       string         `gorm:"unique" json:"cpf"`
	Name      string         `gorm:"column:name" json:"name"`
	Email     string         `gorm:"unique" json:"email"`
	Password  string         `json:"-"`
	TeamID    uuid.UUID      `gorm:"column:team_id;not null, uniqueIndex" json:"-"`
	Team      teamModel.Team `gorm:"foreignKey:TeamID" json:"team"`
	BirthDate time.Time      `json:"birth_date"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}
