package message

import (
	"bytes"
	"encoding/gob"
	"github.com/google/uuid"
	"log"
	"net"
)

type PlayerActionRequest struct {
	SessionID       uuid.UUID
	CommandID       uint32
	ClientCommandID uint8
}

func (z PlayerActionRequest) Envelope(conn net.Conn) Envelope {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(z); err != nil {
		log.Println("Could not encode PlayerActionRequest")
	}
	messageBytes := buf.Bytes()
	return Envelope{
		Conn: conn,
		Raw: Raw{
			Header: Header{
				ProtocolVersion: CurrentProtocolVersion,
				MessageType:     TypePlayerActionRequest,
				Length:          uint16(len(messageBytes)),
			},
			MessageBytes: messageBytes,
		},
	}
}

type PlayerActionResponse struct {
	ClientCommandID uint8
	Result          interface{}
}

func (z PlayerActionResponse) Envelope(conn net.Conn) Envelope {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(z); err != nil {
		log.Println("Could not encode PlayerActionResponse")
	}
	messageBytes := buf.Bytes()
	return Envelope{
		Conn: conn,
		Raw: Raw{
			Header: Header{
				ProtocolVersion: CurrentProtocolVersion,
				MessageType:     TypePlayerActionResponse,
				Length:          uint16(len(messageBytes)),
			},
			MessageBytes: messageBytes,
		},
	}
}
