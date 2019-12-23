package message

import (
	"fmt"
	"github.com/google/uuid"
)

type SessionNotFoundError struct {
	ID uuid.UUID
}

func (e SessionNotFoundError) Error() string {
	return fmt.Sprintf("session for %s not found", e.ID)
}
