package message

import (
	"bytes"
	"encoding/gob"
	"log"
	"net"
)

type ZoneReport struct {
	Header
	Layers []Layer
}

func (z ZoneReport) Envelope(conn net.Conn) Envelope {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(z); err != nil {
		log.Println("Could not encode ZoneReport")
	}
	messageBytes := buf.Bytes()
	return Envelope{
		Conn: conn,
		Raw: Raw{
			Header: Header{
				ProtocolVersion: CurrentProtocolVersion,
				MessageType:     TypeZoneReport,
				Length:          uint16(len(messageBytes)),
			},
			MessageBytes: messageBytes,
		},
	}
}

type Layer struct {
	ID    uint8
	Asset Asset
}

type Asset struct {
	ID          uint
	X           uint
	Y           uint
	StateLength uint
	State       interface{}
}
