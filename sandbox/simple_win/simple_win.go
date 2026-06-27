package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

var logger *log.Logger
var logFile *os.File

func main() {
	var err error
	// Open or create a log file
	logFile, err = os.OpenFile("output.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer logFile.Close()

	// Create a logger that writes to both the file and the terminal
	logger = log.New(io.MultiWriter(logFile, os.Stdout), "APP_LOG: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("Starting app. Logs are also visible in the terminal.")

	go func() {
		w := new(app.Window)

		if err := window(w); err != nil {
			logger.Printf("window error: %v", err)
		}
		logFile.Close()
		os.Exit(0)
	}()

	app.Main()
}

func window(w *app.Window) error {
	var ops op.Ops
	var (
		fileBtn   = new(widget.Clickable)
		actionBtn = new(widget.Clickable)
		helpBtn   = new(widget.Clickable)

		fileOpen, actionOpen, helpOpen bool
	)

	theme := material.NewTheme()

	for {
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			if fileBtn.Clicked(gtx) {
				fileOpen = !fileOpen
				actionOpen = false
				helpOpen = false
				logger.Println("Clicked: File")
			} else if actionBtn.Clicked(gtx) {
				logger.Println("Clicked: Action")
				actionOpen = !actionOpen
				fileOpen = false
				helpOpen = false
			} else if helpBtn.Clicked(gtx) {
				logger.Println("Clicked: Help")
				helpOpen = !helpOpen
				actionOpen = false
				fileOpen = false
			}
			layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					d := material.Button(theme, fileBtn, "File").Layout(gtx)
					if fileBtn.Clicked(gtx) {
						fileOpen = !fileOpen
						actionOpen = false
						helpOpen = false
					}
					return d
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					d := material.Button(theme, actionBtn, "Action").Layout(gtx)
					if actionBtn.Clicked(gtx) {
						actionOpen = !actionOpen
						fileOpen = false
						helpOpen = false
					}
					return d
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					d := material.Button(theme, helpBtn, "Help").Layout(gtx)
					if helpBtn.Clicked(gtx) {
						helpOpen = !helpOpen
						fileOpen = false
						actionOpen = false
					}
					return d
				}),
			)

			if fileOpen {
				layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						newBtn := new(widget.Clickable)
						d := material.Button(theme, newBtn, "New").Layout(gtx)
						if newBtn.Clicked(gtx) {
							logger.Println("Clicked: New")
						}
						return d
					}),
				)
			}

			title := material.H5(theme, "Hello, Gio")
			title.Alignment = text.Start
			layout.Inset{Top: unit.Dp(12)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return title.Layout(gtx)
			})

			e.Frame(gtx.Ops)
		}
	}
}
