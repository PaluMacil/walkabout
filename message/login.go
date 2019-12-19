package message

import (
	"github.com/google/uuid"
)

type LoginRequest struct {
	Header
	CharacterName string
}

type LoginResponse struct {
	Header
	EntityID  uuid.UUID
	SessionID uuid.UUID
}
