package gio_win

import (
	"fmt"
	"log"

	"gioui.org/widget"
)

// handleButtonClicks - Функція для обробки натискань кнопок меню
func handleButtonClicks(
	gtx C,
	BS BtnState,
	shpkXlsxPtr, boXlsxPtr *xlsxData,
	shpkDataPtr *map[string]Person,
	input_window *widget.Editor,
	logger *log.Logger,
) (BtnState, string, string) {

	var (
		default_path_ppd = "d:/tmp/звіт_ППД.xlsx" // Звіт для ППД
		default_path_bo  = "d:/tmp/3бо.xlsx"      // Звіт загального розподілу
		default_path_vac1  = "d:/tmp/відпустки1черги.xlsx"  // Звіт по відпусткам
		text_in_window   = default_path_ppd
		msg              = ""
		shpkXlsx, boXlsx xlsxData
		shpkData map[string]Person
	)

	if shpkXlsxPtr != nil {
		shpkXlsx = *shpkXlsxPtr
	}

	if boXlsxPtr != nil {
		boXlsx = *boXlsxPtr
	}

	if shpkDataPtr != nil {
		shpkData = *shpkDataPtr
	}

	switch {

	case BS.file_btn.Clicked(gtx):
		logger.Println("Натиснуто кнопку: 'Файл'.")
		BS.open_file = !BS.open_file
		BS.open_action = false
		BS.open_help = false

	case BS.action_btn.Clicked(gtx):
		logger.Println("Натиснуто кнопку: 'Звіти'.")
		BS.open_action = !BS.open_action
		BS.open_file = false
		BS.open_help = false

	case BS.help_btn.Clicked(gtx):
		logger.Println("Натиснуто кнопку: 'Допомога'.")
		BS.open_help = !BS.open_help
		BS.open_action = false
		BS.open_file = false

	case BS.shpk_btn.Clicked(gtx):
		logger.Println("Натиснуто кнопку: 'Визначити ШПК'.")
		BS.define_shpk = !BS.define_shpk
		BS.define_distrib = false
		title_shpk := "Виберіть Excel файл ШПК"
		shpkXlsxPtr, err_shpk := OpenFileXlsx(title_shpk, "")
		shpkXlsx = *shpkXlsxPtr
		if err_shpk != nil || shpkXlsx.Data == nil {
			msg = fmt.Sprintf("Помилка відкриття %s", shpkXlsx.FilePath)
			logger.Printf("%s: %v\n", msg, err_shpk)
		} else {
			msg = fmt.Sprintf("Файл %s успішно відкрито.", shpkXlsx.FilePath)
			logger.Println(msg)
		}

	case BS.proto_distrib_btn.Clicked(gtx):
		logger.Println("Натиснуто кнопку: 'Визначити прототип розподілу'.")
		BS.define_distrib = !BS.define_distrib
		BS.define_shpk = false
		title_bo := "Виберіть Excel файл загального розподілу людей"
		if text_in_window != "" {
			text_in_window = input_window.Text()
			default_path_ppd = text_in_window
		} else {
			text_in_window = default_path_ppd
			input_window.SetText(text_in_window)
		}
		boXlsxPtr, err_bo := OpenFileXlsx(title_bo, "")
		boXlsx := *boXlsxPtr
		if err_bo != nil {
			msg := fmt.Sprintf("Помилка відкриття %s з даними розподілу людей:\n%v",
					boXlsx.FilePath, err_bo)
			logger.Println(msg)
		} else {
			msg = fmt.Sprintf("Файл %s успішно відкрито.", boXlsx.FilePath)
			logger.Println(msg)
		}

	case BS.prep_shpk_btn.Clicked(gtx):
		logger.Println("Натиснуто кнопку: 'Підготувати дані ШПК'.")
		BS.prepare_shpk = !BS.prepare_shpk
		BS.prepare_ppd = false
		BS.refresh_distrib = false
		BS.save_vacation = false
		err_shpk := error(nil)
		if shpkXlsx.Data == nil || shpkXlsx.FilePath == "" {
			msg = "Спробуйте знову зчитати ШПК!"
			logger.Println(msg)
		} else {
			shpkDataPtr, err_shpk = ReadShpkData(&shpkXlsx)
			if err_shpk != nil || len(*shpkDataPtr) == 0 {
				msg = fmt.Sprintf("Помилка перетворення даних ШПК в словник: %v\n", err_shpk)
				logger.Println(msg)
			} else {
				msg = "Дані ШПК успішно перетворено з формату xlsx в словник."
				logger.Println(msg)
			}
		}

		case BS.prep_ppd_btn.Clicked(gtx):
			logger.Println("Натиснуто кнопку: 'Записати звіт для стройової'.")
			BS.prepare_ppd = !BS.prepare_ppd
			BS.prepare_shpk = false
			BS.refresh_distrib = false
			BS.save_vacation = false
			err_ppd := []string{}
			if shpkXlsx.Data != nil || shpkDataPtr != nil || len(shpkData) != 0 {
				PPD_COUNTER, PPD_LIST, err_ppd = PrepareReportPPD(shpkDataPtr)
			} else {
				logger.Println("Завантажте і підготуйте дані ШПК для звіту ППД!")
			}
			switch {
			case len(err_ppd) != 0:
				msg = fmt.Sprintf("Помилка підготовки звіту для ППД: \n%v", err_ppd)
				logger.Println(msg)
			case PPD_COUNTER == nil || PPD_LIST == nil :
				logger.Println("Помилка обробки даних ШПК для звіту ППД!")
			default:
				text_in_window = input_window.Text()
				if text_in_window != "" {
					default_path_ppd = text_in_window
				} else {
					text_in_window = default_path_ppd
					input_window.SetText(text_in_window)
				}
				saved_file, err_ppd_save := SaveReportPPD(&PPD_COUNTER, &PPD_LIST,
					text_in_window)
					// msg = fmt.Sprintf("Дані ШПК успішно підготовлено для запису звіту ППД до файлу %s.", text_in_window)
					// logger.Println(msg)
				if err_ppd_save != nil {
					msg = fmt.Sprintf("Помилка збереження звіту ППД до файлу %s", saved_file)
					logger.Printf("%s: \n%v\n", msg, err_ppd_save)
				} else {
					msg = fmt.Sprintf("Звіт для ППД успішно збережений до файлу\n%s.", saved_file)
					logger.Println(msg)
				}
			}

	case BS.refresh_distrib_btn.Clicked(gtx):
		logger.Println("Натиснуто кнопку: 'Оновити весь розподіл'.")
		BS.refresh_distrib = !BS.refresh_distrib
		BS.prepare_shpk = false
		BS.prepare_ppd = false
		BS.save_vacation = false
		if text_in_window != "" {
			text_in_window = input_window.Text()
			default_path_bo = text_in_window
		} else {
			text_in_window = default_path_bo
			input_window.SetText(text_in_window)
		}
		bo_counter, err_bo_count := PrepareReportBO(shpkDataPtr)
		switch {
		case len(err_bo_count) != 0:
			msg = fmt.Sprintf("Помилка підготовки звіту для оновлення розподілу:\n%v",
				err_bo_count)
		case len(bo_counter) == 0:
			msg = "Помилка рахування розподілу особового складу, bo_counter=0."
		case boXlsx.Data == nil:
			msg = fmt.Sprintf("Помилка зчитування даних прототипу розподілу:\n%v",
				err_bo_count)
		default:
			msg = "Дані успішно підготовлено для оновлення розподілу."
			logger.Println(msg)
		text_in_window, err_bo_upd := UpdateDistributionBO(bo_counter,
			&boXlsx, default_path_bo)
		boXlsx.FilePath = text_in_window
		if err_bo_upd != nil {
			msg = fmt.Sprintf("Помилка збереження загального розподілу особового складу до файлу %s.",
				text_in_window)
			logger.Printf("%s: \n%v\n",	msg, err_bo_upd)
		} else {
			msg = fmt.Sprintf("Загальний розподіл особового складу успішно оновлений і збережений до файлу %s.",
				text_in_window)
			logger.Println(msg)
		}}

	case BS.save_vacation_btn.Clicked(gtx):
		logger.Println("Натиснуто кнопку: 'Записати звіт по І відпустці'.")
		BS.save_vacation = !BS.save_vacation
		BS.prepare_shpk = false
		BS.prepare_ppd = false
		BS.refresh_distrib = false
		if shpkXlsx.Data != nil {

			Vac1ReportPtr, err_vac1 := PrepareVacationReport1(shpkDataPtr)
			if err_vac1 != nil || *Vac1ReportPtr == nil {
				msg = "Помилка обробки даних для звіту по 1 черзі відпусток!"
				logger.Printf("%s: \n%v\n",	msg, err_vac1)
			}

			if text_in_window != "" {
				text_in_window = input_window.Text()
				default_path_vac1 = text_in_window
			} else {
				text_in_window = default_path_vac1
				input_window.SetText(text_in_window)
			}

			vac1Filepath, err_save := SaveVacationReport1(Vac1ReportPtr,
				default_path_vac1)
			if err_save != nil {
				msg = fmt.Sprintf("Помилка запису звіту по 1 черзі відпусток до файлу %s!",
				vac1Filepath)
				logger.Printf("%s: \n%v\n",	msg, err_vac1)
			} else {
				msg = fmt.Sprintf("Звіт стосовно 1 черги відпусток збережений до файлу\n%s.",
				text_in_window)
			logger.Println(msg)
			}

		} else {
			logger.Println("Виберіть Excel файл ШПК!")
		}
	}

	return BS, text_in_window, msg
}
