package main

import (
	"log"
	"os"

	"github.com/Azrubael/260601-golang-tutor/260626giowin/gio_win"

	"gioui.org/app"
)


func main() {
	go func() {
		window := new(app.Window)

		err := gio_win.RunWindow(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}
