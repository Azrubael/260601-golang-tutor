package main

import (
	"fmt"

	"image/color"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type C = layout.Context
type D = layout.Dimensions

func run_window(window *app.Window) error {
	var (
		ops                               op.Ops
		file_btn                          = new(widget.Clickable)
		action_btn                        = new(widget.Clickable)
		help_btn                          = new(widget.Clickable)
		input_window                      = new(widget.Editor)
		file_open, action_open, help_open bool
		// new_open, open_open,exit_open bool

		w_width               = 480
		w_height              = 640
		text_in_window string = "d:\\tmp\\filename.xlsx"
	)

	theme := material.NewTheme()

	window.Option(app.Title("XLSX processing app"))
	window.Option(app.Size(unit.Dp(w_width), unit.Dp(w_height)))

	for {
		switch typ := window.Event().(type) {
		case app.DestroyEvent:
			return typ.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, typ)
			if file_btn.Clicked(gtx) {
				file_open = !file_open
				action_open = false
				help_open = false
				text_in_window = input_window.Text()
			}
			if action_btn.Clicked(gtx) {
				action_open = !action_open
				file_open = false
				help_open = false
			}
			if help_btn.Clicked(gtx) {
				help_open = !help_open
				action_open = false
				file_open = false
			}
			// Menu bar
			// Spacing: Leftover space will be at the start
			layout.Flex{Axis: layout.Vertical, Spacing: layout.SpaceStart,
				Alignment: layout.Start}.Layout(gtx,

				layout.Rigid( // The inputbox
					func(gtx C) D {
						// Wrap the editor in material design
						ed := material.Editor(theme, input_window, text_in_window)

						// Define characteristics of the input box
						input_window.SingleLine = true
						input_window.Alignment = text.Middle

						if text_in_window != "" {
							input_str := fmt.Sprint(text_in_window)
							input_window.SetText(input_str)
						}

						margins := layout.Inset{
							Top:    unit.Dp(0),
							Right:  unit.Dp(w_width / 6),
							Bottom: unit.Dp(25),
							Left:   unit.Dp(w_width / 6),
						}

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
				layout.Rigid(func(gtx C) D {
					return renderMenuButton(gtx, theme, file_btn, "File",
						&file_open, &action_open, &help_open)
				}),
				layout.Rigid(func(gtx C) D {
					return renderMenuButton(gtx, theme, action_btn, "Action",
						&action_open, &file_open, &help_open)
				}),
				layout.Rigid(func(gtx C) D {
					return renderMenuButton(gtx, theme, help_btn, "Help",
						&help_open, &file_open, &action_open)
				}),
			)

			// Simple dropdowns under the menu bar
			// You can improve positioning with more advanced layout (e.g., layout.Stack).
			if file_open {
				layout.Inset{
					Top:   unit.Dp(100),
					Left:  unit.Dp(25),
					Right: unit.Dp(25),
				}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return material.Button(theme, new(widget.Clickable), "New").Layout(gtx)
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return layout.Spacer{Height: unit.Dp(5)}.Layout(gtx)
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return material.Button(theme, new(widget.Clickable), "Open").Layout(gtx)
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return layout.Spacer{Height: unit.Dp(5)}.Layout(gtx)
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return material.Button(theme, new(widget.Clickable), "Exit").Layout(gtx)
						}),
					)
				})
			}
			if action_open {
				layout.Inset{
					Top:   unit.Dp(100),
					Left:  unit.Dp(25),
					Right: unit.Dp(25),
				}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return material.Button(theme, new(widget.Clickable), "Run").Layout(gtx)
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return layout.Spacer{Height: unit.Dp(5)}.Layout(gtx)
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return material.Button(theme, new(widget.Clickable), "Stop").Layout(gtx)
						}),
					)
				})
			}
			if help_open {
				layout.Inset{
					Top:   unit.Dp(100),
					Left:  unit.Dp(25),
					Right: unit.Dp(25),
				}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return material.Button(theme, new(widget.Clickable), "About").Layout(gtx)
						}),
					)
				})
			}

			// Center-alligned text in the main window
			title := material.H6(theme, "Hello, Gio!")
			title.Alignment = text.Middle

			// Add some top padding so title isn't drawn on top of menu bar
			layout.Inset{Top: unit.Dp(15)}.Layout(gtx, func(gtx C) D {
				return title.Layout(gtx)
			})

			typ.Frame(gtx.Ops)
		}
	}
}

func renderMenuButton(gtx C, theme *material.Theme, btn *widget.Clickable,
	name string, current *bool, other1 *bool, other2 *bool) D {
	margins := layout.Inset{
		Top:    unit.Dp(5),
		Bottom: unit.Dp(0),
		Right:  unit.Dp(5),
		Left:   unit.Dp(5),
	}
	d := material.Button(theme, btn, name).Layout(gtx)
	return margins.Layout(gtx,
		func(gtx C) D {
			if btn.Clicked(gtx) {
				*current = !*current
				*other1 = false
				*other2 = false
			}
			return d
		},
	)
}
