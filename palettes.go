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

var initPalette = []color.Color{
	color.RGBA{0, 0, 0, 1},
}

func paletteSetter(s string) {
	pChoice = s
}

var palettes = []string{"pico8"}
var paletteMap = map[string][]color.Color{
	"pico8": pico8,
}
