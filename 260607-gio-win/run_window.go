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

			layout.Flex{Axis: layout.Vertical, Spacing: layout.SpaceStart, Alignment: layout.Start}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					d := material.Button(theme, fileBtn, "File").Layout(gtx)
					if actionBtn.Clicked(gtx) {
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
						return material.Button(theme, new(widget.Clickable), "New").Layout(gtx)
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
