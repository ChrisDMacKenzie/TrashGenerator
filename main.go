package main

import (
	"image/gif"
	"os"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

var giffile = "rgb.gif"
var pChoice = defaultPalette
var rChoice = defaultRatio

func main() {
	init := NothingGIF()
	var t = &trash{
		gif:      init,
		image:    init.Image[0],
		imageIdx: 0,
		length:   1,
		hidden:   false,
		size:     fyne.Size{500, 500},
	}
	app := app.New()

	w := app.NewWindow("Trash Generator")
	w.SetOnClosed(cleanup)
	w.SetContent(widget.NewVBox(
		PaletteSelector(),
		RatioSelector(),
		widget.NewButton("Create", func() {
			generate(pChoice, rChoice)
			f, _ := os.Open(giffile)
			g, _ := gif.DecodeAll(f)
			t.gif = g
			t.image = g.Image[0]
			t.imageIdx = 0
			t.length = len(g.Image)
			t.hidden = false
			newSize := fyne.Size{
				aspectRatios[rChoice][0] * 2,
				aspectRatios[rChoice][1] * 2,
			}
			if rChoice == "big" {
				newSize.Width = newSize.Width / 4
				newSize.Height = newSize.Height / 4
			}
			t.Resize(newSize)
			newPosition := fyne.Position{
				(1000 - t.size.Width) / 2,
				99 + ((750 - t.size.Height) / 2),
			}
			t.Move(newPosition)
			f.Close()
		}),
		t,
		SaveButton(t, w),
	))
	t.animate()
	w.ShowAndRun()
}

func PaletteSelector() *widget.Select {
	s := widget.NewSelect(palettes, paletteSetter)
	s.SetSelected(defaultPalette)
	return s
}

func RatioSelector() *widget.Select {
	s := widget.NewSelect(ratioNames, ratioSetter)
	s.SetSelected(defaultRatio)
	return s
}

func cleanup() {
	os.Remove(giffile)
}

func SaveButton(t *trash, w fyne.Window) *widget.Button {
	b := widget.NewButton("Save", func() {
		content := widget.NewEntry()
		content.SetPlaceHolder("name this trash")
		confirmCallback := func(response bool) {
			if response {
				save(t, content.Text)
			}
		}
		dialog.ShowCustomConfirm("Save GIF", "save", "cancel", content, confirmCallback, w)
	})
	return b
}
