package message

import (
	"github.com/PaluMacil/walkabout/nivio"
	"github.com/google/uuid"
)

type PlayerActionRequest struct {
	SessionID       uuid.UUID
	CommandID       uint32
	ClientCommandID uint8
}

// NetworkWrite should have inherited the comment from the interface but because there is no way to know which interface it came from, it wants another comment. So here that it. Yay.
func (playerActionRequest *PlayerActionRequest) NetworkWrite(stream *nivio.Writer) error {
	// Write session ID
	err := stream.WriteBytes(playerActionRequest.SessionID[:])
	if err != nil {
		return err
	}

	// Write command ID
	err = stream.WriteUInt32(playerActionRequest.CommandID)
	if err != nil {
		return err
	}

	// Write the client command ID
	err = stream.WriteUInt8(playerActionRequest.ClientCommandID)
	if err != nil {
		return err
	}

	return nil
}

// NetworkRead should have inherited the comment from the interface but because there is no way to know which interface it came from, it wants another comment. So here that it. Yay.
func (playerActionRequest *PlayerActionRequest) NetworkRead(stream *nivio.Reader) error {
	// Read session ID
	uuidBytes, err := stream.ReadBytes(16)
	if err != nil {
		return err
	}
	// Set the session ID
	copy(playerActionRequest.SessionID[:], uuidBytes)

	// Read the command ID
	playerActionRequest.CommandID, err = stream.ReadUInt32()
	if err != nil {
		return err
	}

	// Read the client command ID
	playerActionRequest.ClientCommandID, err = stream.ReadUInt8()
	if err != nil {
		return err
	}

	return nil
}

type PlayerActionResponse struct {
	ClientCommandID uint8
	Result          interface{}
}
