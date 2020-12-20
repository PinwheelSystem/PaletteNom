# PaletteNom
> 🧵 Load a color palette from an image.

PaletteNom *noms* out 64 colors from an 8x8 image and spits them out to an array.  
The use of this in Pinwheel is to get its 64 color palette easily from a single image.

# Example
```golang
package main

import (
	"fmt"

	"github.com/PinwheelSystem/PaletteNom"
)

func main() {
	palettelib := palettenom.New()
	colors, err := palettelib.Load("palette.png")
	if err != nil {
		panic(err)
	}

	r, g, b, _ := colors[0].RGBA()
	fmt.Println("The 1st color in RGB is:", r, g, b)
}
```

# License
[BSD 3-Clause](LICENSE)