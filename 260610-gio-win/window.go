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

		shpk_btn                          = new(widget.Clickable)
		proto_distrib_btn                 = new(widget.Clickable)

		prep_shpk_btn                     = new(widget.Clickable)
		prep_ppd_btn                      = new(widget.Clickable)
		refresh_distrib_btn               = new(widget.Clickable)
		write_vacation_btn                = new(widget.Clickable)

		input_window                      = new(widget.Editor)
		open_file, open_action, open_help bool
		define_shpk, define_distrib bool
		prepare_shpk, prepare_ppd, refresh_distrib, write_vacation bool

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

			// Menu bar
			layout.Flex{Axis: layout.Vertical, Spacing: layout.SpaceStart,
				Alignment: layout.Start}.Layout(gtx,
				// Text above the button
				layout.Rigid(func(gtx C) D {
						lbl := material.Body1(theme, "Ім'я файлу для звіту:")
						lbl.Alignment = text.Middle
						return lbl.Layout(gtx)
				}),
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
					return renderMenuButton(gtx, theme, file_btn, "Файл",
						&open_file, &open_action, &open_help)
				}),
				layout.Rigid(func(gtx C) D {
					return renderMenuButton(gtx, theme, action_btn, "Звіти",
						&open_action, &open_file, &open_help)
				}),
				layout.Rigid(func(gtx C) D {
					return renderMenuButton(gtx, theme, help_btn, "Допомога",
						&open_help, &open_file, &open_action)
				}),
			)

			// Simple dropdowns under the menu bar
			if open_file {
				layout.Inset{
					Top:   unit.Dp(100),
					Left:  unit.Dp(25),
					Right: unit.Dp(25),
				}.Layout(gtx, func(gtx C) D {
					return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
						layout.Rigid(func(gtx C) D {
							return renderMenuButton(gtx, theme, shpk_btn, "Визначити ШПК",
								&define_shpk, &define_distrib)
							}),
						layout.Rigid(func(gtx C) D {
							return renderMenuButton(gtx, theme, proto_distrib_btn, "Визначити прототип розподілу",
								&define_distrib, &define_shpk)
							}),
					)
				})

			}
			if open_action {
				layout.Inset{
					Top:   unit.Dp(100),
					Left:  unit.Dp(25),
					Right: unit.Dp(25),
				}.Layout(gtx, func(gtx C) D {
					return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
						layout.Rigid(func(gtx C) D {
							return renderMenuButton(gtx, theme, prep_shpk_btn, "Підготувати дані",
								&prepare_shpk, &prepare_ppd, &refresh_distrib, &write_vacation)
							}),
						layout.Rigid(func(gtx C) D {
							return renderMenuButton(gtx, theme, prep_ppd_btn, "Записати звіт для стройової",
								&prepare_ppd, &prepare_shpk, &refresh_distrib, &write_vacation)
							}),
						layout.Rigid(func(gtx C) D {
							return renderMenuButton(gtx, theme, refresh_distrib_btn, "Оновити весь розподіл",
								&refresh_distrib, &prepare_ppd, &prepare_shpk, &write_vacation)
							}),
						layout.Rigid(func(gtx C) D {
							return renderMenuButton(gtx, theme, write_vacation_btn, "Записати звіт по І відпустці",
								&write_vacation, &prepare_ppd, &prepare_shpk, &refresh_distrib)
							}),
					)
				})
			}
			if open_help {
				layout.Inset{
					Top:   unit.Dp(25),
					Left:  unit.Dp(25),
					Right: unit.Dp(25),
				}.Layout(gtx, func(gtx C) D {
					return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
						layout.Rigid(func(gtx C) D {
							var zmist = "  Ця маленька програма призначена прискорити\n" +
													"підготовку трьох типів звітів.\n" +
													"  Після запуску цієї програми, з'явиться вікно, " +
													"де потрібно обрати вхідний файл `xlsx` розподілу " +
													"особового складу.\n" +
													"  Потрібно переконатись, що обрано саме правильний " +
													"файл ШПС, бо існує лише вбудована перевірка " +
													"назви аркуша `ШПС`.\n" +
													"За замовчуванням відбуваються спроба відкриття\n" +
													"ШПС за шляхом:    `d:/tmp/ШПС-T0320.xlsx`\n" +
													"  При обранні прототипу для звіту розподілу " +
													"особового складу, потрібно переконатись, що цей " +
													"файл може бути прототипом, бо існує лише " +
													"перевірка назви аркуша `3БО`.\n" +
													"  Для генерації звіту по відпусткам потрібно " +
													"обрати вхідний файл `xlsx` розподілу особового " +
													"складу, а потім обрати команду `Підготувати дані`."
							return material.Body1(theme, zmist).Layout(gtx)
						}),
					)
				})
			} else {
			// Center-alligned text in the main window
			title := material.H6(theme, "Генератор звітів")
			title.Alignment = text.Middle

			layout.Inset{Top: unit.Dp(15)}.Layout(gtx, func(gtx C) D {
				return title.Layout(gtx)
				})
			}

			typ.Frame(gtx.Ops)
		}
	}
}

func renderMenuButton(gtx C, theme *material.Theme, btn *widget.Clickable,
	name string, current *bool /*, handler func()*/, others ...*bool) D {
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
				for _, o := range others {
						*o = false
				}
				// handler() // call the function passed in
			}
			return d
		},
	)
}
