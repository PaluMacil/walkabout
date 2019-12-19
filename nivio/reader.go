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
func (stream *Reader) ReadBytes(numBytes uint64) ([]byte, error) {
	bytes := make([]byte, numBytes)
	_, error := stream.uw.Read(bytes)
	if error != nil {
		return nil, error
	}
	return bytes, nil
}

// ReadString Reads a string from the stream
func (stream *Reader) ReadString() (string, error) {
	// Reads the length of the string
	numBytes, err := stream.ReadUInt64()
	if err != nil {
		return "", err
	}

	// Read that many bytes from the stream
	bytes, err := stream.ReadBytes(numBytes)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// ReadStruct reads the struct!
func (stream *Reader) ReadStruct(data interface{}) error {
	return binary.Read(stream.uw, binary.LittleEndian, data)
}

// ReadUInt64 Reads the int from the stream
func (stream *Reader) ReadUInt64() (uint64, error) {
	var num uint64
	err := binary.Read(stream.uw, binary.LittleEndian, &num)
	if err != nil {
		return 0, err
	}
	return num, nil
}
