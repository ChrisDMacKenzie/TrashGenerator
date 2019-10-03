package main

import (
	"image/color"
)

var pico8 = []color.Color{
	color.RGBA{0, 0, 0, 1},
	color.RGBA{29, 43, 83, 1},
	color.RGBA{129, 37, 83, 1},
	color.RGBA{0, 135, 81, 1},
	color.RGBA{171, 82, 54, 1},
	color.RGBA{95, 87, 79, 1},
	color.RGBA{194, 195, 199, 1},
	color.RGBA{255, 241, 232, 1},
	color.RGBA{255, 0, 77, 1},
	color.RGBA{255, 163, 0, 1},
	color.RGBA{255, 255, 39, 1},
	color.RGBA{0, 231, 86, 1},
	color.RGBA{41, 173, 255, 1},
	color.RGBA{131, 118, 156, 1},
	color.RGBA{255, 119, 168, 1},
	color.RGBA{255, 204, 170, 1},
}

var girlgang1 = []color.Color{
	color.RGBA{0, 0, 0, 1},
	color.RGBA{208, 119, 211, 1},
	color.RGBA{122, 207, 214, 1},
	color.RGBA{23, 187, 187, 1},
	color.RGBA{255, 108, 140, 1},
	color.RGBA{42, 0, 255, 1},
	color.RGBA{65, 73, 164, 1},
	color.RGBA{175, 100, 157, 1},
	color.RGBA{245, 0, 186, 1},
	color.RGBA{255, 52, 75, 1},
	color.RGBA{255, 120, 160, 1},
	color.RGBA{82, 82, 212, 1},
	color.RGBA{255, 152, 156, 1},
	color.RGBA{255, 234, 146, 1},
	color.RGBA{254, 204, 235, 1},
	color.RGBA{116, 167, 207, 1},
	color.RGBA{121, 228, 174, 1},
}

var initPalette = []color.Color{
	color.RGBA{0, 0, 0, 1},
}

func paletteSetter(s string) {
	pChoice = s
}

var palettes = []string{"pico8", "girlgang1"}
var paletteMap = map[string][]color.Color{
	"pico8": pico8,
	"girlgang1": girlgang1,
}
