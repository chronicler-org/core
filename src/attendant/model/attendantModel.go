package attendantModel

import (
	"time"

	"github.com/google/uuid"

	teamModel "github.com/chronicler-org/core/src/team/model"
)

type Attendant struct {
	ID        uuid.UUID      `gorm:"primarykey"  json:"id"`
	CPF       string         `gorm:"unique" json:"cpf"`
	Name      string         `json:"name"`
	Email     string         `gorm:"unique" json:"email"`
	Password  string         `json:"-"`
	TeamID    uuid.UUID      `gorm:"column:team_id;not null" json:"-"`
	Team      teamModel.Team `gorm:"foreignKey:TeamID" json:"team"`
	BirthDate string         `json:"birth_date"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

func (a Attendant) GetID() uuid.UUID {
	return a.ID
}
