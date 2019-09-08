package main

import (
	"image"
	"image/color"
	"image/gif"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

type trash struct {
	length   int
	imageIdx int
	image    image.Image
	gif      *gif.GIF
	size     fyne.Size
	position fyne.Position
	hidden   bool
}

func (t *trash) Size() fyne.Size {
	return t.size
}

func (t *trash) Resize(size fyne.Size) {
	t.size = size
	widget.Renderer(t).Layout(size)
}

func (t *trash) Position() fyne.Position {
	return t.position
}

func (t *trash) Move(pos fyne.Position) {
	t.position = pos
	widget.Renderer(t).Layout(t.size)
}

func (t *trash) MinSize() fyne.Size {
	return widget.Renderer(t).MinSize()
}

func (t *trash) Visible() bool {
	return !t.hidden
}

func (t *trash) Show() {
	t.hidden = false
}

func (t *trash) Hide() {
	t.hidden = true
}

type trashRenderer struct {
	render  *canvas.Raster
	objects []fyne.CanvasObject

	trash *trash
}

func (t *trashRenderer) MinSize() fyne.Size {
	return fyne.NewSize(750, 600)
}

func (t *trashRenderer) Layout(size fyne.Size) {
	t.render.Resize(size)
}

func (t *trashRenderer) ApplyTheme() {
}

func (t *trashRenderer) BackgroundColor() color.Color {
	return theme.BackgroundColor()
}

func (t *trashRenderer) Refresh() {
	canvas.Refresh(t.render)
}

func (t *trashRenderer) Objects() []fyne.CanvasObject {
	return t.objects
}

func (t *trashRenderer) Destroy() {
}

func (t *trashRenderer) draw(w, h int) image.Image {
	return t.trash.image
}

func (t *trash) CreateRenderer() fyne.WidgetRenderer {
	renderer := &trashRenderer{trash: t}
	render := canvas.NewRaster(renderer.draw)
	renderer.render = render
	renderer.objects = []fyne.CanvasObject{render}

	return renderer
}

func (t *trash) animate() {
	go func() {
		tick := time.NewTicker(time.Second / 10)
		for {
			select {
			case <-tick.C:
				t.image = t.gif.Image[t.imageIdx]
				t.imageIdx = (t.imageIdx + 1) % t.length
				widget.Refresh(t)
			}
		}
	}()
}
