package gio_win

import (
	"fmt"
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
func RunWindow(
	window *app.Window, logger *log.Logger) error {
	var (
		ops            op.Ops
		input_window   = new(widget.Editor)
		BS             BtnState
		title material.LabelStyle
		default_msg = "Генератор звітів"
		text_in_window string
		err_msg 			 string
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
			BS, text_in_window, err_msg = handleButtonClicks(gtx, BS, &SHPK_XLSX, &BO_XLSX,
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

// handleButtonClicks - Функція для обробки натискань кнопок меню
func handleButtonClicks(
	gtx C,
	BS BtnState,
	shpk_xlsx_ptr, bo_xlsx_ptr *xlsxData,
	input_window *widget.Editor,
	logger *log.Logger,
) (BtnState, string, string) {

	var (
		text_in_window = "d:\\tmp\\звіт_ППД.xlsx" // Ім'я файлу для запису звіту ППД
		default_file_path = text_in_window
		msg = ""
	)

	switch {

	case BS.file_btn.Clicked(gtx):
		logger.Println("Натиснуто кнопку: 'Файл'")
		BS.open_file = !BS.open_file
		BS.open_action = false
		BS.open_help = false

	case BS.action_btn.Clicked(gtx):
		logger.Println("Натиснуто кнопку: 'Звіти'")
		BS.open_action = !BS.open_action
		BS.open_file = false
		BS.open_help = false

	case BS.help_btn.Clicked(gtx):
		logger.Println("Натиснуто кнопку: 'Допомога'")
		BS.open_help = !BS.open_help
		BS.open_action = false
		BS.open_file = false

	case BS.shpk_btn.Clicked(gtx):
		logger.Println("Натиснуто кнопку: 'Визначити ШПК'")
		BS.define_shpk = !BS.define_shpk
		BS.define_distrib = false
		title_shpk := "Виберіть Excel файл ШПК"
		err_shpk := error(nil)
		*shpk_xlsx_ptr, err_shpk = OpenFileXlsx(title_shpk, "")
		if err_shpk != nil || SHPK_XLSX.Data == nil {
			msg = fmt.Sprintf("Помилка відкриття %s", SHPK_XLSX.FilePath)
			logger.Printf("Помилка відкриття %s: %v\n", SHPK_XLSX.FilePath, err_shpk)
		} else {
			msg = fmt.Sprintf("Файл %s успішно відкрито.", SHPK_XLSX.FilePath)
			logger.Println(msg)
		}

	case BS.proto_distrib_btn.Clicked(gtx):
		logger.Println("Натиснуто кнопку: 'Визначити прототип розподілу'")
		BS.define_distrib = !BS.define_distrib
		BS.define_shpk = false
		title_bo := "Виберіть Excel файл загального розподілу людей"
		err_bo := error(nil)
		*bo_xlsx_ptr, err_bo = OpenFileXlsx(title_bo, "")
		if err_bo != nil || BO_XLSX.Data == nil {
			msg = fmt.Sprintf("Помилка відкриття %s з даними розподілу людей",
				BO_XLSX.FilePath)
			logger.Printf("Помилка відкриття %s з даними розподілу людей: %v\n",
				BO_XLSX.FilePath, err_bo)
		} else {
			msg = fmt.Sprintf("Файл %s успішно відкрито.", BO_XLSX.FilePath)
			logger.Println(msg)
		}

	case BS.prep_shpk_btn.Clicked(gtx):
		logger.Println("Натиснуто кнопку: 'Підготувати дані ШПК'")
		BS.prepare_shpk = !BS.prepare_shpk
		BS.prepare_ppd = false
		BS.refresh_distrib = false
		BS.save_vacation = false
		err_shpk := error(nil)
		SHPK_DATA, err_shpk = ReadShpkData(shpk_xlsx_ptr)
		if err_shpk != nil || SHPK_DATA == nil {
			msg = fmt.Sprintf("Помилка перетворення даних ШПК в словник: %v\n", err_shpk)
			logger.Println(msg)
		} else {
			msg = "Дані ШПК успішно перетворено з формату xlsx в словник."
			logger.Println(msg)
		}

	case BS.prep_ppd_btn.Clicked(gtx):
		logger.Println("Натиснуто кнопку: 'Записати звіт для стройової'")
		BS.prepare_ppd = !BS.prepare_ppd
		BS.prepare_shpk = false
		BS.refresh_distrib = false
		BS.save_vacation = false
		if text_in_window != "" {
			text_in_window = input_window.Text()
		} else {
			text_in_window = default_file_path
			input_window.SetText(text_in_window)
		}
		logger.Println(text_in_window)

		err_ppd := []string{}
		PPD_COUNTER, PPD_LIST, err_ppd = PrepareReportPPD(SHPK_DATA)
		if len(err_ppd) != 0 || SHPK_DATA == nil {
			msg = fmt.Sprintf("Помилка підготовки звіту для ППД: %v", err_ppd)
			logger.Println(msg)
		} else {
			msg = "Дані ШПК успішно підготовлено для звіту ППД."
			logger.Println(msg)
		}
		saved_file, err_ppd_save := SaveReportPPD(&PPD_COUNTER, &PPD_LIST, text_in_window)
		if err_ppd_save != nil {
			msg = fmt.Sprintf("Помилка збереження звіту ППД до файлу %s", saved_file)
			logger.Printf("Помилка збереження звіту ППД до файлу %s: %v\n", saved_file, err_ppd_save)
		} else {
			msg = fmt.Sprintf("Звіт для ППД успішно збережений в файл %s.", saved_file)
			logger.Println(msg)
		}

	case BS.refresh_distrib_btn.Clicked(gtx):
		logger.Println("Натиснуто кнопку: 'Оновити весь розподіл'")
		BS.refresh_distrib = !BS.refresh_distrib
		BS.prepare_shpk = false
		BS.prepare_ppd = false
		BS.save_vacation = false
		UpdateDistributionBO()

	case BS.save_vacation_btn.Clicked(gtx):
		logger.Println("Натиснуто кнопку: 'Записати звіт по І відпустці'")
		BS.save_vacation = !BS.save_vacation
		BS.prepare_shpk = false
		BS.prepare_ppd = false
		BS.refresh_distrib = false
		SaveVacationReport1()
	}

	return BS, text_in_window, msg
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
