package main

import (
	"bytes"
	"io"
	"log"

	"github.com/PaluMacil/walkabout/nivio"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type messageHeader struct {
	Version uint32
}

type position struct {
	X uint32
	Y uint32
}

type entity struct {
	Name     string
	Position position
}

func (ent *entity) WriteTo(stream *nivio.Writer) {
	stream.WriteString(ent.Name)
	stream.WriteStruct(ent.Position)
}

func (ent *entity) ReadFrom(stream *nivio.Reader) {
	ent.Name = stream.ReadString()
	stream.ReadStruct(&ent.Position)
}

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	ebitenutil.DebugPrint(screen, "Hello, World!")
	return nil
}

func main() {

	buff := &bytes.Buffer{}

	sendMessage(buff)
	readMessage(buff)

	// if err := ebiten.Run(update, 1024, 768, 1, "Hello, World!"); err != nil {
	// 	log.Fatal(err)
	// }
}

func sendMessage(sourceWriter io.Writer) {
	x := messageHeader{
		Version: 12,
	}

	ent := entity{
		Name: "Bob",
		Position: position{
			X: 12,
			Y: 13,
		},
	}

	// Create buffer
	stream := nivio.NewWriter(sourceWriter)

	// Write header
	stream.WriteStruct(x)

	// Write the entity
	ent.WriteTo(stream)

	// Flush / Send
	stream.Flush()
}

func readMessage(sourceReader io.Reader) {

	// Create the reader
	stream := nivio.NewReader(sourceReader)

	// Read header
	var x messageHeader
	stream.ReadStruct(&x)

	// Read the entity
	ent := entity{}
	ent.ReadFrom(stream)

	log.Print("")

}
