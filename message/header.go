package message

import (
	"encoding/binary"
	"net"
)

var CurrentProtocolVersion = uint8(0)

// Header starts every message with a simple preamble
type Header struct {
	ProtocolVersion uint8
	MessageType     Type
	Length          uint16
}

func (h Header) Bytes() []byte {
	buf := make([]byte, HeaderSize)
	buf[0] = h.ProtocolVersion
	binary.LittleEndian.PutUint16(buf[1:3], uint16(h.MessageType))
	binary.LittleEndian.PutUint16(buf[3:5], h.Length)

	return buf
}

const HeaderSize = 5

func GetHeader(conn net.Conn) Header {
	var header Header

	buf := make([]byte, HeaderSize)
	conn.Read(buf)
	header.ProtocolVersion = buf[0]
	header.MessageType = Type(binary.LittleEndian.Uint16(buf[1:3]))
	header.Length = binary.LittleEndian.Uint16(buf[3:])

	return header
}

type Type uint16

const (
	// Client to Server
	TypeLoginRequest        Type = 0
	TypePlayerActionRequest Type = 1

	// Server to client
	TypeLoginResponse        Type = 1000
	TypePlayerActionResponse      = 1001
	TypeZoneReport                = 1002
)
