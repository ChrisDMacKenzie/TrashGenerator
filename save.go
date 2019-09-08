package main

import (
	"image/gif"
	"os"
	"os/user"
	"path"
	"runtime"
)

func save(t *trash, text string) {
	myself, _ := user.Current()
	homedir := myself.HomeDir
	var desktop string
	if runtime.GOOS == "windows" {
		desktop = homedir + "\\Desktop\\"
	} else {
		desktop = homedir + "/Desktop/"
	}
	filename := text + ".gif"
	pathToFile := path.Join(desktop, filename)
	f, _ := os.OpenFile(pathToFile, os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	gif.EncodeAll(f, t.gif)
}
