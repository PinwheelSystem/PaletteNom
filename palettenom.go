// PaletteNom loads a 64 color palette from a 8x8 image.
package palettenom

import (
	"os"
	"image"
	"image/color"

	_ "image/png"
	_ "golang.org/x/image/bmp"
)
// Contains our palette data
type PaletteNom struct {
	Palette []color.Color
}

// Creates a new PaletteNom instance
func New() PaletteNom {
	return PaletteNom{make([]color.Color, 64)}
}

func (p *PaletteNom) Load(filename string) []color.Color {
	file, _ := os.Open(filename)
	image, _, _ := image.Decode(file)

	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			color := image.At(x, y)
			p.Palette[x + (8 * y)] = color
		}
	}

	return p.Palette
}