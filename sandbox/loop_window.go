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

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			// Menu bar
			layout.Flex{Axis: layout.Horizontal, Alignment: layout.Start}.Layout(gtx,
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
			)

			// Simple dropdowns under the menu bar
			// You can improve positioning with more advanced layout (e.g., layout.Stack).
			if fileOpen {
				layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return material.Button(theme, new(widget.Clickable), "New").Layout(gtx)
					}),
				)
			}

			// Left-aligned title
			title := material.H5(theme, "Hello, Gio")
			title.Alignment = text.Start
			// Add some top padding so title isn't drawn on top of menu bar
			layout.Inset{Top: unit.Dp(12)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return title.Layout(gtx)
			})

			e.Frame(gtx.Ops)
		}
	}
}