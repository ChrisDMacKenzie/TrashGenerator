package main

import (
	"fmt"
	"image"
	"image/gif"
	"math"
	"os"
	"strings"
)

func main() {
	var w, h int = 240, 240
	var ops []string
	operators := getRandomOperators(NumOperations)
	for _, o := range operators {
		o.setSecondaryOps()
		ops = append(ops, o.print())
	}

	expr := strings.Join(ops, "*")
	fmt.Println(expr)
	palette := pico8

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
