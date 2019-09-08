package main

import (
	"image/gif"
	"os"
)

func save(t *trash) {
	filename := "randomTrash.gif"
	f, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	gif.EncodeAll(f, t.gif)
}
