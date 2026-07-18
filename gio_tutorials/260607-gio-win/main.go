package main

import (
	"log"
	"os"

	"gioui.org/app"
)

func main() {
	go func() {
		w := new(app.Window)

		if err := window(w); err != nil {
			log.SetOutput(os.Stderr)
			log.Printf("window error: %v\n", err)
		}
		os.Exit(0)
	}()
	app.Main()
}
