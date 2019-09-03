package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"math"
	"os"
	"strings"
)

func main() {
	var w, h int = 240, 240
	var ops []string
	// numOps := r.Intn(5) + 1
	operators := getRandomOperators(3)
	// operators.print()
	for _, o := range operators {
		o.setSecondaryOps()
		ops = append(ops, o.print())
	}

	expr := strings.Join(ops, "*")
	fmt.Println(expr)
	var palette = []color.Color{
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

	var images []*image.Paletted
	var delays []int
	steps := 25
	for t := 0; t < steps; t++ {
		img := image.NewPaletted(image.Rect(0, 0, w, h), palette)
		images = append(images, img)
		delays = append(delays, 0)

		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				colorIdx := float64(1)
				for _, o := range operators {
					xf := float64(x)
					yf := float64(y)
					tf := float64(t)
					colorIdx = colorIdx * o.compute(xf, yf, tf)
				}
				finalIdx := int(math.Abs(float64(int(colorIdx) % len(img.Palette))))
				// fmt.Println(colorIdx, finalIdx)
				img.Set(x, y, img.Palette[finalIdx])
			}
		}
	}

	f, _ := os.OpenFile("rgb.gif", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	gif.EncodeAll(f, &gif.GIF{
		Image: images,
		Delay: delays,
	})
}
