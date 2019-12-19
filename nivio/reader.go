package nivio

import (
	"bufio"
	"encoding/binary"
	"io"
)

// Reader , The :D
type Reader struct {
	uw *bufio.Reader
}

// NewReader Creates a new NivIO reader backed by the sourceReader
func NewReader(sourceReader io.Reader) *Reader {
	return &Reader{
		uw: bufio.NewReader(sourceReader),
	}
}

// ReadBytes Reads at maximum the given number of bytes from the stream and returns them
func (stream *Reader) ReadBytes(numBytes uint64) []byte {
	bytes := make([]byte, numBytes)
	stream.uw.Read(bytes)
	return bytes
}

// ReadString Reads a string from the stream
func (stream *Reader) ReadString() string {
	// Reads the length of the string, followed by the slice of bytes, and converts to string
	return string(stream.ReadBytes(stream.ReadUInt64()))
}

// ReadStruct reads the struct!
func (stream *Reader) ReadStruct(data interface{}) {
	binary.Read(stream.uw, binary.LittleEndian, data)
}

// ReadUInt64 Reads the int from the stream
func (stream *Reader) ReadUInt64() uint64 {
	var num uint64
	binary.Read(stream.uw, binary.LittleEndian, &num)
	return num
}
