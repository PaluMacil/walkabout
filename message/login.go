package message

import (
	"bytes"
	"encoding/gob"
	"github.com/google/uuid"
	"log"
	"net"
)

type LoginRequest struct {
	CharacterName string
}

type LoginResponse struct {
	EntityID  uuid.UUID
	SessionID uuid.UUID
}

func (l LoginResponse) Envelope(conn net.Conn) Envelope {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(l); err != nil {
		log.Println("Could not encode LoginResponse")
	}
	messageBytes := buf.Bytes()
	return Envelope{
		Conn: conn,
		Raw: Raw{
			Header: Header{
				ProtocolVersion: CurrentProtocolVersion,
				MessageType:     TypeLoginResponse,
				Length:          uint16(len(messageBytes)),
			},
			MessageBytes: messageBytes,
			StateBytes:   nil,
		},
	}
}
