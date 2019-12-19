package nivio

import (
	"bufio"
	"encoding/binary"
	"io"
)

// Writer The Writer :D
type Writer struct {
	uw *bufio.Writer
}

// NewWriter Creates a new NivIO writer backed by the sourceWriter
func NewWriter(sourceWriter io.Writer) *Writer {
	return &Writer{
		uw: bufio.NewWriter(sourceWriter),
	}
}

// Flush Flushes the stream
func (stream *Writer) Flush() error {
	return stream.uw.Flush()
}

func (stream *Writer) Write(p []byte) (n int, err error) {
	return stream.uw.Write(p)
}

// WriteBytes Writes the given slice of bytes into the stream
func (stream *Writer) WriteBytes(bytes []byte) error {
	_, error := stream.uw.Write(bytes)
	return error
}

// WriteString Writes the string into the stream
func (stream *Writer) WriteString(str string) error {
	// Write the len
	error := stream.WriteUInt64(uint64(len(str)))
	if error != nil {
		return error
	}

	// Get the bytes
	return stream.WriteBytes([]byte(str))
}

// WriteStruct writes the struct!
func (stream *Writer) WriteStruct(data interface{}) error {
	return binary.Write(stream.uw, binary.LittleEndian, data)
}

// WriteUInt64 Writes an unsigned sixtyfour bit integer into the stream
func (stream *Writer) WriteUInt64(num uint64) error {
	return binary.Write(stream.uw, binary.LittleEndian, &num)
}
