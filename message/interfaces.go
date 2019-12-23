package message

import (
	"github.com/PaluMacil/walkabout/nivio"
	"net"
)

// NetworkSerializable objects can be wrote to and read from the network
type NetworkSerializable interface {
	// NetworkWrite Writes the object for transport over the network
	NetworkWrite(stream *nivio.Writer) error

	// NetworkRead Reads the object from the network
	NetworkRead(stream *nivio.Reader) error
}

type Handler interface {
	HandleConnection(conn net.Conn)
	Send(envelope Envelope)
}

type Server interface {
	HandleSessionConnections()
	WaitForInterrupt()
	EnqueueMessage(envelope Envelope)
}
