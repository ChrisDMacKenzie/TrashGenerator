package main

import (
	"image"
	"image/color"
	"image/gif"
	"math"
	"os"
	"strconv"
	"sync"
)

func generate(p, r string) {
	mapMutex := new(sync.Mutex)
	wg := new(sync.WaitGroup)
	frames := make(map[string]*image.Paletted, NumSteps)
	palette := paletteMap[p]
	ratio := aspectRatios[r]
	w, h := ratio[0], ratio[1]
	var ops []string
	operators := getRandomOperators(NumOperations)
	for _, o := range operators {
		o.setSecondaryOps()
		ops = append(ops, o.print())
	}

	var images []*image.Paletted
	var delays []int
	steps := NumSteps
	wg.Add(steps)
	for t := 0; t < steps; t++ {
		delays = append(delays, 0)
		go setFrame(w, h, t, operators, frames, palette, wg, mapMutex)
	}
	wg.Wait()
	for t := 0; t < steps; t++ {
		key := strconv.Itoa(t)
		images = append(images, frames[key])
	}

	f, _ := os.OpenFile("rgb.gif", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	gif.EncodeAll(f, &gif.GIF{
		Image: images,
		Delay: delays,
	})
}

func setFrame(w, h, t int, operators []operator, frames map[string]*image.Paletted, palette []color.Color, wg *sync.WaitGroup, mu *sync.Mutex) {
	img := image.NewPaletted(image.Rect(0, 0, w, h), palette)
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
	key := strconv.Itoa(t)
	mu.Lock()
	frames[key] = img
	mu.Unlock()
	wg.Done()
}
