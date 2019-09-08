package main

import (
	"image"
	"image/gif"
)

func NothingGIF() *gif.GIF {
	img := image.NewPaletted(image.Rect(0, 0, 240, 240), initPalette)
	delays := []int{0}

	for x := 0; x < 240; x++ {
		for y := 0; y < 240; y++ {
			img.Set(x, y, img.Palette[0])
		}
	}

	return &gif.GIF{
		Image: []*image.Paletted{img},
		Delay: delays,
	}
}
