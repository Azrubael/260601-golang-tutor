package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type C = layout.Context
type D = layout.Dimensions

func main() {
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
	w.Option(app.Title("GIO UI app"))
	w.Option(app.Size(unit.Dp(320), unit.Dp(480)))
	var ops op.Ops

	var startButton widget.Clickable // startButton is a clickable widget

	th := material.NewTheme() // th defines the material design style

	for { // listen for events in the window
		evt := w.Event()

		switch typ := evt.(type) {
		case app.FrameEvent: // when the application should re-render
			gtx := app.NewContext(&ops, typ)
			layout.Flex{
				Axis:    layout.Vertical,   // From top to bottom
				Spacing: layout.SpaceStart, // Leftover space at the top
			}.Layout(gtx,
				layout.Rigid(
					func(gtx C) D {
						// Define margins around the button using layout.Inset
						margins := layout.Inset{
							Top:    unit.Dp(25),
							Bottom: unit.Dp(25),
							Right:  unit.Dp(35),
							Left:   unit.Dp(35),
						}
						return margins.Layout(gtx,
							func(gtx C) D {
								return material.Button(th, &startButton, "Start").Layout(gtx)
							},
						)

					},
				),
				layout.Rigid(
					layout.Spacer{Height: unit.Dp(25)}.Layout,
				),
			)

			typ.Frame(gtx.Ops)

		case app.DestroyEvent:
			return typ.Err
		}
	}
}
