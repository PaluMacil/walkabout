package message

import "github.com/PaluMacil/walkabout/nivio"

// NetworkSerializable objects can be wrote to and read from the network
type NetworkSerializable interface {
	// NetworkWrite Writes the object for transport over the network
	NetworkWrite(stream *nivio.Writer) error

	// NetworkRead Reads the object from the network
	NetworkRead(stream *nivio.Reader) error
}
