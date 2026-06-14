package gio_win

import (
	"log"
	"os"

	"gioui.org/app"
)

func main() {
	go func() {
		window := new(app.Window)

		err := run_window(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}
