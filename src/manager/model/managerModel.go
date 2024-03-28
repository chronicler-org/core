package managerModel

import (
	"time"

	"github.com/google/uuid"
)

type Manager struct {
	ID        uuid.UUID `gorm:"primarykey"  json:"id"`
	CPF       string    `gorm:"unique" json:"cpf"`
	Name      string    `json:"name"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"-"`
	TeamID    uuid.UUID `gorm:"column:team_id;not null" json:"team_id"`
	BirthDate time.Time `json:"birth_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
