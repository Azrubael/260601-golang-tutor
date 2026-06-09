package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type C = layout.Context
type D = layout.Dimensions

var progress float32
var progressIncrementer chan float32

func main() {
	// Setup a separate channel to provide ticks to increment progress
	progressIncrementer = make(chan float32)
	go func() {
		for {
			time.Sleep(time.Second / 25)
			progressIncrementer <- 0.004
		}
	}()
	go func() {
		// create new window
		w := new(app.Window)
		err := draw(w)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func draw(w *app.Window) error {
	w_width := 320
	w_height := 480
	w.Option(app.Title("GIO UI app"))
	w.Option(app.Size(unit.Dp(w_width), unit.Dp(w_height)))

	var ops op.Ops
	var startButton widget.Clickable    // startButton is a clickable widget
	var boiling bool                    // is the egg boiling?
	var boilDurationInput widget.Editor // A textfield to input boil duration
	var boilDuration float32

	th := material.NewTheme() // th defines the material design style

	go func() { // listen for events in the incrementer channel
		for p := range progressIncrementer {
			if boiling && progress < 1 {
				progress += p
				w.Invalidate() // Force a redraw by invalidating the frame
			}
		}
	}()

	for { // listen for events in the window

		switch typ := w.Event().(type) {
		case app.FrameEvent: // when the application should re-render
			gtx := app.NewContext(&ops, typ)
			if startButton.Clicked(gtx) {
				boiling = !boiling
				if progress >= 1 { // Resetting the boil
					progress = 0
				}
				// Read from the input box
				inputString := boilDurationInput.Text()
				inputString = strings.TrimSpace(inputString)
				inputFloat, _ := strconv.ParseFloat(inputString, 32)
				boilDuration = float32(inputFloat)
				boilDuration = boilDuration / (1 - progress)
			}

			layout.Flex{
				Axis:    layout.Vertical,   // From top to bottom
				Spacing: layout.SpaceStart, // Leftover space at the top
			}.Layout(gtx,
				layout.Rigid( // Draw an ellipce
					func(gtx C) D {
						circle := clip.Ellipse{
							Min: image.Pt(gtx.Constraints.Max.X/2-120, 0),
							Max: image.Pt(gtx.Constraints.Max.X/2+120, 240),
						}.Op(gtx.Ops)
						color := color.NRGBA{R: 255, G: uint8(239 * (1 - progress)), B: uint8(174 * (1 - progress)), A: 255}
						paint.FillShape(gtx.Ops, color, circle)
						d := image.Point{Y: 400}
						return layout.Dimensions{Size: d}
					},
				),
				// The inputbox
				layout.Rigid(
					func(gtx C) D {
						// Wrap the editor in material design
						ed := material.Editor(th, &boilDurationInput, "sec")

						// Define characteristics of the input box
						boilDurationInput.SingleLine = true
						boilDurationInput.Alignment = text.Middle

						// Count down the text when boiling
						if boiling && progress < 1 {
							boilRemain := (1 - progress) * boilDuration
							inputStr := fmt.Sprintf("%.1f", math.Round(float64(boilRemain)*10)/10)
							boilDurationInput.SetText(inputStr)
						}

						// Define insets ...
						margins := layout.Inset{
							Top:    unit.Dp(0),
							Right:  unit.Dp(w_width / 3),
							Bottom: unit.Dp(25),
							Left:   unit.Dp(w_width / 3),
						}

						// ... and borders ...
						border := widget.Border{
							Color:        color.NRGBA{R: 204, G: 204, B: 204, A: 255},
							CornerRadius: unit.Dp(3),
							Width:        unit.Dp(2),
						}

						// ... before laying it out, one inside the other
						return margins.Layout(gtx,
							func(gtx C) D {
								return border.Layout(gtx, ed.Layout)
							},
						)
					},
				),
				layout.Rigid( // Draw a progress bar
					func(gtx C) D {
						bar := material.ProgressBar(th, progress)
						return bar.Layout(gtx)
					},
				),
				layout.Rigid(
					func(gtx C) D { // Define button with margins around it
						margins := layout.Inset{
							Top:    unit.Dp(25),
							Bottom: unit.Dp(25),
							Right:  unit.Dp(35),
							Left:   unit.Dp(35),
						}
						return margins.Layout(gtx,
							func(gtx C) D {
								var text string
								if !boiling {
									text = "Start"
								} else {
									text = "Stop"
								}
								return material.Button(th, &startButton, text).Layout(gtx)
							},
						)

					},
				),
			)

			typ.Frame(gtx.Ops)

		case app.DestroyEvent:
			return typ.Err
		}
	}
}
