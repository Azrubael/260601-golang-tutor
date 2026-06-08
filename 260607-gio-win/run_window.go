package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func run_window(window *app.Window) error {
	var ops op.Ops
	var (
		fileBtn   = new(widget.Clickable)
		actionBtn = new(widget.Clickable)
		helpBtn   = new(widget.Clickable)

		fileOpen, actionOpen, helpOpen bool
	)

	theme := material.NewTheme()

	window.Option(app.Title("GIO UI app"))
	window.Option(app.Size(unit.Dp(320), unit.Dp(480)))
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			// Menu bar
			// Spacing: Leftover space will be at the start
			layout.Flex{Axis: layout.Vertical, Spacing: layout.SpaceStart, Alignment: layout.Start}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					d := material.Button(theme, fileBtn, "File").Layout(gtx)
					if fileBtn.Clicked(gtx) {
						fileOpen = !fileOpen
						// close others if desired:
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
				layout.Rigid(
					// The height of the spacer is 25 Device independent pixels
					layout.Spacer{Height: unit.Dp(25)}.Layout,
				),
			)

			// Simple dropdowns under the menu bar
			// You can improve positioning with more advanced layout (e.g., layout.Stack).
			if fileOpen {
				layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return material.Button(theme, new(widget.Clickable), "New").Layout(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return material.Button(theme, new(widget.Clickable), "Open").Layout(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return material.Button(theme, new(widget.Clickable), "Exit").Layout(gtx)
					}),
				)
			}
			if actionOpen {
				layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return material.Button(theme, new(widget.Clickable), "Run").Layout(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return material.Button(theme, new(widget.Clickable), "Stop").Layout(gtx)
					}),
				)
			}
			if helpOpen {
				layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return material.Button(theme, new(widget.Clickable), "About").Layout(gtx)
					}),
				)
			}

			// Center-alligned text in the main window
			title := material.H6(theme, "Hello, Gio!")
			title.Alignment = text.Middle

			// Add some top padding so title isn't drawn on top of menu bar
			layout.Inset{Top: unit.Dp(12)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return title.Layout(gtx)
			})

			e.Frame(gtx.Ops)
		}
	}
}
