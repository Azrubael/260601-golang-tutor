package main

import (
	"log"
	"os"

	"gioui.org/app"
)

func main() {
	go func() {
		w := new(app.Window)

		err := window(w)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}
