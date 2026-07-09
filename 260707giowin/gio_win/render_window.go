package gio_win

import (
	"log"

	"image/color"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// RunWindow - головна функція, яка запускає графічне вікно програми
func RenderWindow(
	window *app.Window, logger *log.Logger) error {
	var (
		ops            op.Ops
		input_window   = new(widget.Editor)
		BS             BtnState
		title          material.LabelStyle
		default_msg    = "Генератор звітів"
		text_in_window string
		err_msg        string
		w_width        = 480
		w_height       = 640
	)

	theme := material.NewTheme()
	title = material.Body1(theme, default_msg)
	BS.file_btn = new(widget.Clickable)
	BS.action_btn = new(widget.Clickable)
	BS.help_btn = new(widget.Clickable)

	BS.shpk_btn = new(widget.Clickable)
	BS.proto_distrib_btn = new(widget.Clickable)

	BS.prep_shpk_btn = new(widget.Clickable)
	BS.prep_ppd_btn = new(widget.Clickable)
	BS.refresh_distrib_btn = new(widget.Clickable)
	BS.save_vacation_btn = new(widget.Clickable)

	window.Option(app.Title("Звіти XLSX"))
	window.Option(app.Size(unit.Dp(w_width), unit.Dp(w_height)))

	for {
		switch typ := window.Event().(type) {
		case app.DestroyEvent:
			return typ.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, typ)
			BS, text_in_window, err_msg = handleButtonClicks(gtx, BS, &SHPK_XLSX, &BO_XLSX, &SHPK_DATA,
				input_window, logger)

			// Кнопки для вибору дій, які відображаються в головному вікні програми
			layout.Flex{Axis: layout.Vertical, Spacing: layout.SpaceStart,
				Alignment: layout.Start}.Layout(gtx,
				// Текстовий заголовок для поля вводу імені файлу для звіту
				layout.Rigid(func(gtx C) D {
					lbl := material.Body1(theme, "Ім'я файлу для звіту ППД:")
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
					return renderMenuButton(gtx, theme, BS.file_btn, "Файл",
						BS.open_file, BS.open_action, BS.open_help)
				}),
				layout.Rigid(func(gtx C) D {
					return renderMenuButton(gtx, theme, BS.action_btn, "Звіти",
						BS.open_action, BS.open_file, BS.open_help)
				}),
				layout.Rigid(func(gtx C) D {
					return renderMenuButton(gtx, theme, BS.help_btn, "Допомога",
						BS.open_help, BS.open_file, BS.open_action)
				}),
			)

			// Відображення кнопок для вибору дій, якщо прапорець відповідної кнопки кореневого меню задіяний
			if BS.open_file {
				layout.Inset{
					Top:   unit.Dp(120),
					Left:  unit.Dp(25),
					Right: unit.Dp(25),
				}.Layout(gtx, func(gtx C) D {
					return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
						layout.Rigid(func(gtx C) D {
							return renderMenuButton(gtx, theme, BS.shpk_btn, "Визначити ШПК",
								BS.define_shpk, BS.define_distrib)
						}),
						layout.Rigid(func(gtx C) D {
							return renderMenuButton(gtx, theme, BS.proto_distrib_btn, "Визначити прототип розподілу",
								BS.define_distrib, BS.define_shpk)
						}),
					)
				})

			}
			if BS.open_action {
				layout.Inset{
					Top:   unit.Dp(100),
					Left:  unit.Dp(25),
					Right: unit.Dp(25),
				}.Layout(gtx, func(gtx C) D {
					return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
						layout.Rigid(func(gtx C) D {
							return renderMenuButton(gtx, theme, BS.prep_shpk_btn,
								"Підготувати дані ШПК",
								BS.prepare_shpk, BS.prepare_ppd, BS.refresh_distrib, BS.save_vacation)
						}),
						layout.Rigid(func(gtx C) D {
							return renderMenuButton(gtx, theme, BS.prep_ppd_btn,
								"Записати звіт для стройової",
								BS.prepare_ppd, BS.prepare_shpk, BS.refresh_distrib, BS.save_vacation)
						}),
						layout.Rigid(func(gtx C) D {
							return renderMenuButton(gtx, theme, BS.refresh_distrib_btn,
								"Оновити весь розподіл",
								BS.refresh_distrib, BS.prepare_ppd, BS.prepare_shpk, BS.save_vacation)
						}),
						layout.Rigid(func(gtx C) D {
							return renderMenuButton(gtx, theme, BS.save_vacation_btn,
								"Записати звіт по І відпустці",
								BS.save_vacation, BS.prepare_ppd, BS.prepare_shpk, BS.refresh_distrib)
						}),
					)
				})
			}
			if BS.open_help {
				layout.Inset{
					Top:   unit.Dp(25),
					Left:  unit.Dp(25),
					Right: unit.Dp(25),
				}.Layout(gtx, func(gtx C) D {
					err_msg = ""
					title = material.Body1(theme, err_msg)
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
			}

			if err_msg != "" {
				title = material.Body1(theme, err_msg)
			}
			// Вивід заголовку програми, вирівняного по центру вікна
			title.Alignment = text.Middle

			layout.Inset{Top: unit.Dp(15)}.Layout(gtx, func(gtx C) D {
				return title.Layout(gtx)
			})

			typ.Frame(gtx.Ops)
		}
	}
}

// renderMenuButton - Функція для відображення кнопки меню з можливістю вибору
func renderMenuButton(gtx C, theme *material.Theme, btn *widget.Clickable,
	name string, current_flag bool, other_flags ...bool) D {
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
				current_flag = !current_flag
				for f := range other_flags {
					other_flags[f] = false
				}
				// handler() // виклик переданої функції для обробки натискання кнопки
			}
			return d
		},
	)
}
