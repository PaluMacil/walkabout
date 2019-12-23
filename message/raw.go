package message

import "net"

type Envelope struct {
	Conn net.Conn
	Raw  Raw
}

type Raw struct {
	Header       Header
	MessageBytes []byte
	StateBytes   []byte
}

func (r Raw) Bytes() []byte {
	bytes := append(r.Header.Bytes(), r.MessageBytes...)
	bytes = append(bytes, r.StateBytes...)
	return bytes
}
