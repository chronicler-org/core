package authInterface

import "github.com/google/uuid"

type IUser interface {
	GetID() uuid.UUID
}
