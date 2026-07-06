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
	input_window *widget.Editor,
	logger *log.Logger,
) (BtnState, string, string) {

	var (
		default_path_ppd = "d:/tmp/звіт_ППД.xlsx" // Ім'я файлу для запису звіту ППД
		default_path_bo  = "d:/tmp/3бо.xlsx"      // Ім'я файлу для запису звіту ППД
		text_in_window   = default_path_ppd
		msg              = ""
		shpkDataPtr = &SHPK_DATA
	)

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
		err_shpk := error(nil)
		shpkXlsxPtr, err_shpk = OpenFileXlsx(title_shpk, "")
		if err_shpk != nil || SHPK_XLSX.Data == nil {
			msg = fmt.Sprintf("Помилка відкриття %s", SHPK_XLSX.FilePath)
			logger.Printf("%s: %v\n", msg, err_shpk)
		} else {
			msg = fmt.Sprintf("Файл %s успішно відкрито.", SHPK_XLSX.FilePath)
			logger.Println(msg)
		}

	case BS.proto_distrib_btn.Clicked(gtx):
		logger.Println("Натиснуто кнопку: 'Визначити прототип розподілу'.")
		BS.define_distrib = !BS.define_distrib
		BS.define_shpk = false
		title_bo := "Виберіть Excel файл загального розподілу людей"
		err_bo := error(nil)
		boXlsxPtr, err_bo = OpenFileXlsx(title_bo, "")
		if err_bo != nil || BO_XLSX.Data == nil {
			msg = fmt.Sprintf("Помилка відкриття %s з даними розподілу людей",
				BO_XLSX.FilePath)
			logger.Printf("%s: %v\n", msg, err_bo)
		} else {
			msg = fmt.Sprintf("Файл %s успішно відкрито.", BO_XLSX.FilePath)
			logger.Println(msg)
		}

	case BS.prep_shpk_btn.Clicked(gtx):
		logger.Println("Натиснуто кнопку: 'Підготувати дані ШПК'.")
		BS.prepare_shpk = !BS.prepare_shpk
		BS.prepare_ppd = false
		BS.refresh_distrib = false
		BS.save_vacation = false
		err_shpk := error(nil)
		shpkDataPtr, err_shpk = ReadShpkData(shpkXlsxPtr)
		SHPK_DATA = *shpkDataPtr
		if err_shpk != nil || SHPK_DATA == nil {
			msg = fmt.Sprintf("Помилка перетворення даних ШПК в словник: %v\n", err_shpk)
			logger.Println(msg)
			} else {
				msg = "Дані ШПК успішно перетворено з формату xlsx в словник."
				logger.Println(msg)
			}

		case BS.prep_ppd_btn.Clicked(gtx):
			logger.Println("Натиснуто кнопку: 'Записати звіт для стройової'.")
			BS.prepare_ppd = !BS.prepare_ppd
			BS.prepare_shpk = false
		BS.refresh_distrib = false
		BS.save_vacation = false
		if text_in_window != "" {
			text_in_window = input_window.Text()
			} else {
				text_in_window = default_path_ppd
				input_window.SetText(text_in_window)
			}
			logger.Println(text_in_window)
			err_ppd := []string{}
			PPD_COUNTER, PPD_LIST, err_ppd = PrepareReportPPD(shpkDataPtr)
		if len(err_ppd) != 0 {
			msg = fmt.Sprintf("Помилка підготовки звіту для ППД: %v", err_ppd)
			logger.Println(msg)
		} else {
			msg = "Дані ШПК успішно підготовлено для звіту ППД."
			logger.Println(msg)
		}
		saved_file, err_ppd_save := SaveReportPPD(&PPD_COUNTER, &PPD_LIST,
			text_in_window)
		if err_ppd_save != nil {
			msg = fmt.Sprintf("Помилка збереження звіту ППД до файлу %s", saved_file)
			logger.Printf("%s: %v\n", msg, err_ppd_save)
		} else {
			msg = fmt.Sprintf("Звіт для ППД успішно збережений до файлу\n%s.", saved_file)
			logger.Println(msg)
		}

	case BS.refresh_distrib_btn.Clicked(gtx):
		logger.Println("Натиснуто кнопку: 'Оновити весь розподіл'.")
		BS.refresh_distrib = !BS.refresh_distrib
		BS.prepare_shpk = false
		BS.prepare_ppd = false
		BS.save_vacation = false
		if text_in_window != "" {
			text_in_window = input_window.Text()
		} else {
			text_in_window = default_path_bo
			input_window.SetText(text_in_window)
		}
		BO_COUNTER, err_bo_count := PrepareReportBO(shpkDataPtr)
		switch {
		case len(err_bo_count) != 0:
			msg = fmt.Sprintf("Помилка підготовки звіту для оновлення розподілу:\n%v",
				err_bo_count)
		case len(BO_COUNTER) == 0:
			msg = "Помилка рахування розподілу особового складу, BO_COUNTER=0."
		case BO_XLSX.Data == nil:
			msg = fmt.Sprintf("Помилка зчитування даних прототипу розподілу:\n%v",
				err_bo_count)
		default:
			msg = "Дані успішно підготовлено для оновлення розподілу."
			logger.Println(msg)
		text_in_window, err_bo_upd := UpdateDistributionBO(BO_COUNTER,
			boXlsxPtr, default_path_bo)
		BO_XLSX.FilePath = text_in_window
		if err_bo_upd != nil {
			msg = fmt.Sprintf("Помилка збереження загального розподілу особового складу до файлу %s.",
				text_in_window)
			logger.Printf("%s: %v\n",	msg, err_bo_upd)
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
		SaveVacationReport1(nil)
	}

	return BS, text_in_window, msg
}
