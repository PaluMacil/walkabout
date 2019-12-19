package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	ebitenutil.DebugPrint(screen, "Hello, World!")
	return nil
}

func main() {
	if err := ebiten.Run(update, 1024, 768, 1, "Hello, World!"); err != nil {
		log.Fatal(err)
	}
}
