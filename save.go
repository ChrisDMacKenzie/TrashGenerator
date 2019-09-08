package main

import (
	"image/gif"
	"os"
	"path"
)

func save(t *trash, text string) {
	filename := text + ".gif"
	pathToFile := path.Join(currentDir, filename)
	f, _ := os.OpenFile(pathToFile, os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	gif.EncodeAll(f, t.gif)
}
