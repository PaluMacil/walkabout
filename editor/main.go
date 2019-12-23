package main

import (
	"image"
	"log"
	"math"
	"strconv"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	tileSize          int = 16
	tileSelectorWidth int = 512
	levelWidth        int = 1024
	screenHeight      int = 768
)

const (
	screenWidth       int = levelWidth + tileSelectorWidth
	levelGridWidth    int = levelWidth / tileSize
	selectorGridWidth int = tileSelectorWidth / tileSize
	gridHeight        int = screenHeight / tileSize
)

// NivTileSet should be commented
type NivTileSet struct {
	tileset *ebiten.Image
	tiles   []*ebiten.Image
}

// LevelData should be commented
type LevelData struct {
	tilesetID int
	tileID    int
}

var tilesets []*NivTileSet
var level [levelGridWidth][gridHeight]LevelData
var cursorHover *ebiten.Image
var cursorSelect *ebiten.Image
var selectedTile int = 32

func init() {
	// Load the images
	tileset0Image, _, err := ebitenutil.NewImageFromFile("./assets/0_RPGMaker_Outside_A2.png", ebiten.FilterLinear)
	cursorHover, _, err = ebitenutil.NewImageFromFile("./assets/CursorHover.png", ebiten.FilterLinear)
	cursorSelect, _, err = ebitenutil.NewImageFromFile("./assets/CursorSelect.png", ebiten.FilterLinear)

	// Break into tiles
	tileset0 := tileify(tileset0Image)

	// Add to the tilesets
	tilesets = append(tilesets, tileset0)

	if err != nil {
		log.Fatal(err)
	}
}

func tileify(img *ebiten.Image) *NivTileSet {
	// Create the tileset
	var tileSet = &NivTileSet{
		tileset: img,
	}

	// Calculate the tile bounds
	tileBounds := img.Bounds().Size().Div(tileSize)

	// For each tile
	for y := 0; y < tileBounds.Y; y++ {
		for x := 0; x < tileBounds.X; x++ {
			// Calculate position in the tileset
			u, v := x*tileSize, y*tileSize
			// Create the subimage and append to the slice
			tileSet.tiles = append(
				tileSet.tiles,
				img.SubImage(image.Rectangle{
					Min: image.Point{
						X: u,
						Y: v,
					},
					Max: image.Point{
						X: u + tileSize,
						Y: v + tileSize,
					},
				}).(*ebiten.Image))
		}
	}

	return tileSet
}

func stepScale(value int, stepSize int) int {
	return int(math.Floor(float64(value / stepSize)))
}

// i = x + (y * width)
func xytoi(x int, y int, width int) int {
	return x + (y * width)
}

// (x,y) = (id % width, id / width)
func itoxy(id int, width int) (int, int) {
	return (id % width), (id / width)
}

func clampInt(v int, min int, max int) int {
	if v > max {
		return max
	}
	if v < min {
		return min
	}
	return v
}

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// Get mouse and bound position
	mx, my := ebiten.CursorPosition()
	cursor := image.Point{X: clampInt(mx, 0, screenWidth), Y: clampInt(my, 0, screenHeight)}

	// Create selector rect(todo move out)
	selectorBounds := image.Rect(0, 0, tileSelectorWidth, screenHeight)
	// Create level rect(todo move out)
	levelBounds := image.Rect(tileSelectorWidth, 0, screenWidth, screenHeight)

	// Clicked?
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		// In the bounds of the selector?
		if cursor.In(selectorBounds) {
			// Select something
			selectedTile = xytoi(stepScale(mx, tileSize), stepScale(my, tileSize), selectorGridWidth)
		} else if cursor.In(levelBounds) {
			// Change tile
			level[stepScale(mx-tileSelectorWidth, tileSize)][stepScale(my, tileSize)].tileID = selectedTile
		}
	} else if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		if cursor.In(levelBounds) {
			// Copy tile
			selectedTile = level[stepScale(mx-tileSelectorWidth, tileSize)][stepScale(my, tileSize)].tileID
		}
	}

	var drawOptions = &ebiten.DrawImageOptions{
		GeoM: ebiten.GeoM{},
	}

	// Draw tileset
	screen.DrawImage(tilesets[0].tileset, drawOptions)

	// Draw level
	for y := 0; y < gridHeight; y++ {
		// Back to 0,0
		drawOptions.GeoM.Reset()
		// Down and left
		drawOptions.GeoM.Translate(float64(tileSelectorWidth-tileSize), float64(y*tileSize))

		for x := 0; x < levelGridWidth; x++ {
			// Right
			drawOptions.GeoM.Translate(float64(tileSize), 0)

			// Get the tile IDs from the level
			tileIDs := level[x][y]

			// Draw the tile
			screen.DrawImage(tilesets[tileIDs.tilesetID].tiles[tileIDs.tileID], drawOptions)
		}
	}

	// Draw selection
	sx, sy := itoxy(selectedTile, selectorGridWidth)
	drawOptions.GeoM.Reset()
	drawOptions.GeoM.Translate(float64(sx*tileSize), float64(sy*tileSize))
	screen.DrawImage(cursorSelect, drawOptions)

	// Draw cursor
	drawOptions.GeoM.Reset()
	drawOptions.GeoM.Translate(float64(stepScale(mx, tileSize)*tileSize), float64(stepScale(my, tileSize)*tileSize))
	screen.DrawImage(cursorHover, drawOptions)

	ebiten.SetWindowTitle(strconv.Itoa(mx) + "," + strconv.Itoa(my))
	return nil
}

func main() {

	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "Title"); err != nil {
		log.Fatal(err)
	}
}
