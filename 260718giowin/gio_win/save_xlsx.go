package gio_win

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
	"unicode/utf8"

	"github.com/xuri/excelize/v2"
)

// AlignColumnWidth - Функція для вирівнювання динамічно по ширині колонок "B"..."I", а колонку "А" - статично. Вирівнювання для першого рядка не відбувається
func AlignXlsxColumnWidth(xlsx *excelize.File, sheetName string,
	maxRow int, maxCol int) (string, error) {

		for col := 1; col <= maxCol; col++ {
		colName, _ := excelize.ColumnNumberToName(col)
		maxLen := 0
			for row := 2; row <= maxRow; row++ {
				cell, err := xlsx.GetCellValue(sheetName, fmt.Sprintf("%s%d", colName, row))
				if err != nil {
					log.Println("Помилка отримання кількості літерів в клітинці для вирівнювання по ширині: ", err)
					return sheetName, err
				}
				if l := utf8.RuneCountInString(cell); l > maxLen {
					maxLen = l
				}
		}

		var width float64 = 0
		if col == 1 {
			width = 15.0
		} else {
			width = float64(maxLen) + 2.0
		}

		if width < 10 {
			width = 10
		}

		if err := xlsx.SetColWidth(sheetName, colName, colName, width); err != nil {
			log.Println("Помилка вирівнювання колонок по ширині: ", err)
			return sheetName, err
		}
	}
	return  "", nil
}

// SetRowValueGeneric - Функція для додавання до xlsx об'єкту рядка даних []cellValue
func SetRowValueGeneric[T cellValue](f *excelize.File, sheet string,
	row int, colOffset int, values []T) error {

	for i, v := range values {

		colName, err := excelize.ColumnNumberToName(i + colOffset)
		if err != nil {
			return fmt.Errorf("invalid column number: %w", err)
		}
		cell := fmt.Sprintf("%s%d", colName, row)
		var toWrite any

		// Заміна нульових значень на пусті рядки
		switch val := any(v).(type) {
		case int, int8, int16, int32, int64,
			uint, uint8, uint16, uint32, uint64:
			if fmt.Sprintf("%v", v) == "0" {
				toWrite = ""
			} else {
				toWrite = val
			}
		case float32, float64:
			if z := fmt.Sprintf("%v", v); z == "0.0" || z == "0,0" || z == "0" {
				toWrite = ""
			} else {
				toWrite = val
			}
		case string:
			toWrite = val
		default:
			toWrite = fmt.Sprintf("%v", val)
		}

		if err = f.SetCellValue(sheet, cell, toWrite); err != nil {
			return err
		}
	}
	return nil
}

// saveXlsxFile - Функція для  запису типового файлу *.xlsx на жорсткий диск
func saveXlsxFile(xlsxPtr *XlsxData, defaultPath string) (
	factFilepath string, err error) {

	if xlsxPtr == nil {
		err_arg := fmt.Errorf("Дані прототипу розподілу особового не передано до 'saveXlsxFile()'!\n")
		log.Println(err_arg)
		return "", err
	}
	xlsx := (*xlsxPtr).Data
	selectedPath := (*xlsxPtr).FilePath
	if selectedPath == "" {
		selectedPath = defaultPath
	}

	timeStamp := time.Now().Format("060102")
	directory := filepath.Dir(defaultPath)
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		msgDir := fmt.Sprintf("Директорії %s не існує, створіть її самі!\n", directory)
		log.Println(msgDir)
		errDir := fmt.Errorf("%s", msgDir)
		return defaultPath, errDir
	} else {
		log.Println("Успішно перевірено існування директорії: ", directory)
	}
	filename := filepath.Join(directory, timeStamp+"-"+filepath.Base(defaultPath))
	if err := xlsx.SaveAs(filename); err != nil {
		log.Println("Помилка збережання даних в xlsx файл: ", err)
		return filename, err
	}

	return filename, nil
}

// SaveReportPPD - Функція для запису звіту ППД
func SaveReportPPD(ppd_counter_ptr *map[string]Distribution,
	ppd_list_ptr *[][]ShortPersData, pathReportPPD string) (string, error) {

	// fmt.Println("SaveReportPPD() called")
	if len(*ppd_counter_ptr) == 0 || len(*ppd_list_ptr) == 0 {
		errPtr := fmt.Errorf("Завантажте і підготуйте дані ШПК для звіту ППД!")
		log.Println(errPtr)
		return "", errPtr
	}

	// Створюємо об'єкт xlsx і додаємо до нього дані
	xlsx := excelize.NewFile()
	sheetName := xlsx.GetSheetName(xlsx.GetActiveSheetIndex())
	now := time.Now()
	dateTime := now.Format("02.01.2006")
	sheetTitle := fmt.Sprintf("Розподіл особового складу 3бо станом на %v", dateTime)
	if errCell := xlsx.SetCellValue(sheetName, "A1", sheetTitle); errCell != nil {
		log.Println(errCell)
		return "", errCell
	}
	if errCell := xlsx.SetCellValue(sheetName, "A2", ""); errCell != nil {
		log.Println(errCell)
		return "", errCell
	}

	// Таблиця звіту, з которої буде створений xlsx файл
	reportData := [][]string{}
	emptyRow := []string{"", "", "", "", "", "", "", "", ""}
	reportData = append(reportData, emptyRow)
	tableHat := []string{
		"Призначення", "Офіцери", "Сержанти", "Солдати", "Загалом",
		"", "", "", "",
	}
	reportData = append(reportData, tableHat)

	// Додаємо в таблицю звіту лічильників розподілу
	for _, element := range PPD_report_list {
		dist := (*ppd_counter_ptr)[element]
		dataRow := []string{
			element,
			fmt.Sprintf("%d", dist.Offi),
			fmt.Sprintf("%d", dist.Serg),
			fmt.Sprintf("%d", dist.Sold),
			fmt.Sprintf("%d", dist.Total),
			"", "", "", "",
		}
		reportData = append(reportData, dataRow)
	}
	reportData = append(reportData, emptyRow)

	// Створюємо список особового складу за призначенням
	for r, assignment := range *ppd_list_ptr {
		reportData = append(reportData, emptyRow)
		dataRow := []string{"", "", "", "", PPD_report_list[r]}
		reportData = append(reportData, dataRow)
		for i, person := range assignment {
			dataRow = []string{"", "", "", "", "", strconv.Itoa(i),
				person.Rank, person.Name, person.Department,
			}
			reportData = append(reportData, dataRow)
		}
	}

	// Допоміжна безіменна функція для перевірки пустих рядків
	isBlankRow := func(row []string) bool {
		for i := range 9 {
			if i < len(row) && row[i] != "" {
				return false
			}
		}
		return true
	}

	// Оголошення об'єкту "жирні літери"
	boldStyle, err := xlsx.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
			Size: 12,
		},
	})
	if err != nil {
		log.Println("Помилка оголошення об'єкту 'жирні літери':", err)
		return "", err
	}

	// Цикл запису значень з масиву reportData в об'єкт xlsx
	for rowIdx, dataRow := range reportData {
		idxRow := rowIdx + 2
		if err := SetRowValueGeneric(xlsx, sheetName, idxRow, 1, dataRow); err != nil {
			log.Printf("Помилка запису значень рядка %d в об'єкт xlsx:\n %v", idxRow, err)
			return "", err
		}

		// Робиться перший рядок жирними літерами
		if idxRow == 2 {
			if err := xlsx.SetCellStyle(sheetName, "A1", "I1", boldStyle); err != nil {
				log.Printf("Помилка встановлення артібуту 'жирні літери' для першого рядка: \n%v", err)
				return "", err
			}
			continue
		}

		// Робиться рядок жирними літерами після пустого A:I рядка
		if rowIdx > 0 && isBlankRow(reportData[rowIdx-1]) && !isBlankRow(dataRow) {
			endCol, _ := excelize.ColumnNumberToName(9)
			startCell := fmt.Sprintf("A%d", idxRow)
			endCell := fmt.Sprintf("%s%d", endCol, idxRow)
			if err := xlsx.SetCellStyle(sheetName, startCell, endCell, boldStyle); err != nil {
				log.Printf("Помилка встановлення артібуту 'жирні літери' для рядка %d:\n %v", rowIdx, err)
				return "", err
			}
		}
	}

	// Вирівнюємо динамічно по ширині колонки xlsx
	maxRow := len(reportData) + 1
	msg, err_align := AlignXlsxColumnWidth(xlsx, sheetName, maxRow, 9)
	if err_align != nil {
		log.Printf("При вирівнюванні xlsx аркуша %s виникла помилка:\n%v", msg, err_align)
	}

	// Зберігаємо дані в файл
	var xlsxFile = XlsxData{
		Data:     xlsx,
		FilePath: pathReportPPD,
	}
	factFilepath, err_save := saveXlsxFile(&xlsxFile, "d:/tmp/звіт_ППД.xlsx")
	if err_save != nil {
		log.Println(err_save)
		return factFilepath, err_save
	}
	return factFilepath, nil
}

// UpdateDistributionBO - Оновлення загального розподілу особового складу та запис оновлених даних в новий файл
func UpdateDistributionBO(
	boReportCounter map[string]map[string]Distribution, bo_xlsx_ptr *XlsxData,
	pathReportBO string) (string, error) {

	switch {
	case len(boReportCounter) == 0:
		err_arg := fmt.Errorf("Завантажте і підготуйте дані ШПК для для оновлення загального розподілу.\n")
		log.Println(err_arg)
		return "", err_arg
	case bo_xlsx_ptr == nil:
		err_arg := fmt.Errorf("Дані прототипу розподілу особового не передано до 'UpdateDistributionBO()'!\n")
		log.Println(err_arg)
		return "", err_arg
	}

	xlsx := (*bo_xlsx_ptr).Data
	sheetName := "3БО"
	distribMatrix := [][]int{}
	distribLine := []int{}
	for _, comp := range COMP_list {
		for _, brl := range BO_report_list {
			el := boReportCounter[comp][brl]
			distribLine = append(distribLine,
				el.Offi, el.Serg, el.Sold,
			)
		}
		fmt.Println(comp, distribLine)
		distribMatrix = append(distribMatrix, distribLine)
		distribLine = []int{}
	}

	// Оновлення об'єкту, що містить дані в форматі xlsx, перед записом в файл
	startCol := 6 // Починаємо заповнювати зі стовпчика 'F'
	for rowIdx, dataRow := range distribMatrix {
		idxRow := rowIdx + 3
		if err := SetRowValueGeneric(xlsx, sheetName, idxRow, startCol, dataRow); err != nil {
			log.Printf("Помилка запису значень рядка %d в об'єкт xlsx:\n %v", idxRow, err)
			return "", err
		}
	}

	// Зберігаємо дані в файл
	factFilepath, err_save := saveXlsxFile(bo_xlsx_ptr, "d:/tmp/3бо.xlsx")
	if err_save != nil {
		log.Println(err_save)
		return factFilepath, err_save
	}
	return factFilepath, nil
}

func SaveVacationReport1(VacReport1_ptr *[][]string,
	pathReportVac1 string) (filePath string, err error) {

	// fmt.Println("SaveVacationReport1() called")
	if len(*VacReport1_ptr) == 0 {
		errPtr := fmt.Errorf("Завантажте і підготуйте дані ШПК для звіту по відпусткам 1 черги!")
		log.Println(errPtr)
		return "", errPtr
	}

	// Створюємо об'єкт xlsx і додаємо до нього дані
	xlsx := excelize.NewFile()
	sheetName := xlsx.GetSheetName(xlsx.GetActiveSheetIndex())

	now := time.Now()
	dateTime := now.Format("02.01.2006")
	firstLineText := fmt.Sprintf("Кількість особового складу 3бо, що відгуляла першу частину щорічної відпустки станом на %v", dateTime)
	if errCell := xlsx.SetCellValue(sheetName, "A1", firstLineText); errCell != nil {
		log.Println(errCell)
		return "", errCell
	}

	for rowIdx, dataRow := range *VacReport1_ptr{
		idxRow := rowIdx + 2
		if err := SetRowValueGeneric(xlsx, sheetName, idxRow, 1, dataRow); err != nil {
			log.Printf("Помилка запису значень рядка %d в об'єкт xlsx:\n %v", idxRow, err)
			return "", err
		}
	}

	// Оголошення об'єкту "жирні літери"
	boldStyle, err := xlsx.NewStyle(&excelize.Style{
		Font: &excelize.Font{Bold: true, Size: 12},
	})
	if err != nil {
		log.Println("Помилка оголошення об'єкту 'жирні літери':", err)
		return "", err
	}

	// Робиться перший рядок жирними літерами
	if err := xlsx.SetCellStyle(sheetName, "A1", "I1", boldStyle); err != nil {
		log.Printf("Помилка встановлення артібуту 'жирні літери' для першого рядка: \n%v", err)
		return "", err
	}

	// Вирівнюємо динамічно по ширині колонки xlsx
	maxRow := len(*VacReport1_ptr) + 1
	msg, err_align := AlignXlsxColumnWidth(xlsx, sheetName, maxRow, 9)
	if err_align != nil {
		log.Printf("При вирівнюванні xlsx аркуша %s виникла помилка:\n%v", msg, err_align)
	}

	// Зберігаємо дані в файл
	var xlsxFile = XlsxData{
		Data:     xlsx,
		FilePath: pathReportVac1,
	}
	filePath, err_save := saveXlsxFile(&xlsxFile, "d:/tmp/відпустки1черги.xlsx")
	if err_save != nil {
		log.Println(err_save)
		return filePath, err_save
	}

	return filePath, nil
}