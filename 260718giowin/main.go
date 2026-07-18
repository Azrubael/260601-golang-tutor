package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Azrubael/260601-golang-tutor/260718giowin/gio_win"

	"gioui.org/app"
)

var logger *log.Logger
var logFile *os.File

func main() {
	var err error
	// Відкрити чи створити логфайл
	logFile, err = os.OpenFile("go_output.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer logFile.Close()

	// Створюється логер, який записує дані як у файл, так і в термінал.
	logger = log.New(io.MultiWriter(logFile, os.Stdout), "APP_LOG: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("Starting app. Logs are also visible in the terminal.")

	go func() {
		window := new(app.Window)

		if err := gio_win.RenderWindow(window, logger); err != nil {
			logger.Printf("window error: %v", err)
		}
		logFile.Close()
		os.Exit(0)
	}()
	app.Main()
}
