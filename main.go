package main

import (
	"image/gif"
	"os"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

var file = "rgb.gif"

func main() {
	f, _ := os.Open(file)
	g, _ := gif.DecodeAll(f)
	f.Close()
	var t = &trash{
		gif:      g,
		image:    g.Image[0],
		imageIdx: 1,
		length:   NumSteps,
		hidden:   false,
		size:     fyne.Size{500, 500},
	}
	app := app.New()

	w := app.NewWindow("Trash Generator")
	w.SetContent(widget.NewVBox(
		widget.NewButton("Create", func() {
			generate()
			f, _ := os.Open(file)
			g, _ = gif.DecodeAll(f)
			t.gif = g
			t.image = g.Image[0]
			t.imageIdx = 1
			t.length = NumSteps
			t.hidden = false
			f.Close()
		}),
		t,
		widget.NewButton("Save", func() {
			save(t)
		}),
	))
	t.animate()
	w.ShowAndRun()
}
