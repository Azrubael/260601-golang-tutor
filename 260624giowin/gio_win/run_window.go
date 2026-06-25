package gio_win

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

// C та D - скорочення для layout.Context та layout.Dimensions
type C = layout.Context
type D = layout.Dimensions

// RunWindow - головна функція, яка запускає графічне вікно програми
func RunWindow(window *app.Window) error {
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
		// SHPK_XLSX *excelize.File
		// SHPK_FILE_PATH string
		// BO_XLSX *excelize.File
		// BO_FILE_PATH string
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
			if file_btn.Clicked(gtx) {
				open_file = !open_file
				open_action = false
				open_help = false
				text_in_window = input_window.Text()
			} else if action_btn.Clicked(gtx) {
				open_action = !open_action
				open_file = false
				open_help = false
			} else if help_btn.Clicked(gtx) {
				open_help = !open_help
				open_action = false
				open_file = false
			}
			// Кнопки для вибору дій, які відображаються в головному вікні програми
			layout.Flex{Axis: layout.Vertical, Spacing: layout.SpaceStart,
				Alignment: layout.Start}.Layout(gtx,
				// Текстовий заголовок для поля вводу імені файлу для звіту
				layout.Rigid(func(gtx C) D {
						lbl := material.Body1(theme, "Ім'я файлу для звіту:")
						lbl.Alignment = text.Middle
						return lbl.Layout(gtx)
				}),
				layout.Rigid( // Поле вводу імені файлу для звіту
					func(gtx C) D {
						// Обгортка поля вводу в матеріальний дизайн
						ed := material.Editor(theme, input_window, text_in_window)

						// Визначення характеристик поля вводу
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

			// Відображення кнопок для вибору дій, якщо прапорець відповідної кнопки кореневого меню задіяний
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
								/*
								SHPK_XLSX, SHPK_FILE_PATH, err := OpenFileXlsx()
								*/
							}),
						layout.Rigid(func(gtx C) D {
							return renderMenuButton(gtx, theme, proto_distrib_btn, "Визначити прототип розподілу",
								&define_distrib, &define_shpk)
								/*OpenFileXlsx*/
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
							return renderMenuButton(gtx, theme, prep_shpk_btn, "Підготувати дані ШПК",
								&prepare_shpk, &prepare_ppd, &refresh_distrib, &write_vacation)
								/*ReadShpkData*/
							}),
						layout.Rigid(func(gtx C) D {
							return renderMenuButton(gtx, theme, prep_ppd_btn, "Записати звіт для стройової",
								&prepare_ppd, &prepare_shpk, &refresh_distrib, &write_vacation)
								/*SaveReportPPD*/
							}),
						layout.Rigid(func(gtx C) D {
							return renderMenuButton(gtx, theme, refresh_distrib_btn, "Оновити весь розподіл",
								&refresh_distrib, &prepare_ppd, &prepare_shpk, &write_vacation)
								/*UpdateDistributionBO*/
							}),
						layout.Rigid(func(gtx C) D {
							return renderMenuButton(gtx, theme, write_vacation_btn, "Записати звіт по І відпустці",
								&write_vacation, &prepare_ppd, &prepare_shpk, &refresh_distrib)
								/*SaveVacationReport1()*/
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
			// Вивід заголовку програми, вирівняного по центру вікна
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

// renderMenuButton - Функція для відображення кнопки меню з можливістю вибору
func renderMenuButton(
	gtx C, theme *material.Theme, btn *widget.Clickable, name string,
	/* handler func()*/ current_flag *bool , other_flags ...*bool) D {
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
				*current_flag = !*current_flag
				for _, o := range other_flags {
						*o = false
				}
				// handler() // виклик переданої функції для обробки натискання кнопки
			}
			return d
		},
	)
}
