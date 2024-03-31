package attendantModel

import (
	"time"

	"github.com/google/uuid"
)

type AttendantEvaluation struct {
	ID          uuid.UUID `gorm:"primarykey"  json:"id"`
	Score       uint8     `gorm:"not null" json:"score"`
	Description string    `gorm:"not null" json:"description"`
	AvaluatedID uuid.UUID `gorm:"type:uuid;column:avaluated_id;not null" json:"-"`
	Avaluated   Attendant `gorm:"foreignKey:AvaluatedID" json:"avaluated"`
	AvaluatorID uuid.UUID `gorm:"type:uuid;column:avaluator_id;not null" json:"-"`
	Avaluator   Attendant `gorm:"foreignKey:AvaluatorID" json:"avaluator"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
