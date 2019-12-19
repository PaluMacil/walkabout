package message

import "github.com/google/uuid"

type PlayerActionRequest struct {
	Header
	SessionID       uuid.UUID
	CommandID       uint
	ClientCommandID uint8
	DataLength      uint
	Data            interface{}
}

type PlayerActionResponse struct {
	Header
	ClientCommandID uint8
	Result          interface{}
}
