package message

import "github.com/PaluMacil/walkabout/nivio"

// Header should have a comment
type Header struct {
	ProtocolVersion uint8
}

// NetworkWrite should have inherited the comment from the interface but because there is no way to know which interface it came from, it wants another comment. So here that it. Yay.
func (header *Header) NetworkWrite(stream *nivio.Writer) error {
	return stream.WriteUInt8(header.ProtocolVersion)
}

// NetworkRead should have inherited the comment from the interface but because there is no way to know which interface it came from, it wants another comment. So here that it. Yay.
func (header *Header) NetworkRead(stream *nivio.Reader) error {
	var err error
	header.ProtocolVersion, err = stream.ReadUInt8()
	return err
}
