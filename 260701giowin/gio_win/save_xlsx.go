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

// SetRowValue - Функція для додавання до xlsx об'єкту цілого рядка
func SetRowValue(f *excelize.File, sheet string, row int, values []string) error {
	for i, v := range values {
		col, _ := excelize.ColumnNumberToName(i + 1)
		cell := fmt.Sprintf("%s%d", col, row)
		if err := f.SetCellValue(sheet, cell, v); err != nil {
			log.Fatal(err)
			return err
		}
	}
	return nil
}

// SaveReportPPD - Функція для запису звіту ППД
func SaveReportPPD(ppd_counter_ptr *map[string]Distribution,
	ppd_list_ptr *[][]ShortPersData, text_in_window string) (string , error) {

	fmt.Println("SaveReportPPD() called")
	if len(*ppd_counter_ptr) == 0 || len(*ppd_list_ptr) == 0 {
		errPtr := fmt.Errorf("Будь ласка, завантажте і підготуйте дані ШПК для звіту ППД.")
		log.Println(errPtr)
		return "", errPtr
	}

	// Створюємо об'єкт xlsx і додаємо до нього дані
	xlsx := excelize.NewFile()
	sheetName := xlsx.GetSheetName(xlsx.GetActiveSheetIndex())
	now := time.Now()
	dateTime := now.Format("02.01.2006")
	sheetTitle := fmt.Sprintf("Розподіл особового складу 3бо станом на %v", dateTime)
	if errCell := xlsx.SetCellValue(sheetName, "A1", sheetTitle); errCell != nil{
		log.Println(errCell)
		return "", errCell
	}
	if errCell := xlsx.SetCellValue(sheetName, "A2", ""); errCell != nil{
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
	for _, element := range ppd_report_list {
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
		dataRow := []string{"", "", "", "", ppd_report_list[r]}
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
		if err := SetRowValue(xlsx, sheetName, idxRow, dataRow); err != nil {
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

	// Вирівнюємо динамічно по ширині колонки "B"..."I"
	maxRow := len(reportData)
	for col := 1; col <= 9; col++ {
		colName, _ := excelize.ColumnNumberToName(col)
		maxLen := 0
		for row := 2; row <= maxRow; row++ {
			cell, err := xlsx.GetCellValue(sheetName, fmt.Sprintf("%s%d", colName, row))
			if err != nil {
				log.Println("Помилка отримання кількості літерів в клітинці для вирівнювання по ширині: ", err)
				return "", err
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
			return "", err
		}
	}

	// Зберігаємо дані в файл
	timeStamp := now.Format("060102")
	dirPPD := ""
	if text_in_window == "" {
		text_in_window = "d:/tmp/звіт_ППД.xlsx"
	}

	dirPPD = filepath.Dir(text_in_window)
	if _, err := os.Stat(dirPPD); os.IsNotExist(err) {
		msgDir := fmt.Sprintf("Директорії %s не існує, створіть її cfvs!\n", dirPPD)
		log.Println(msgDir)
		errDir := fmt.Errorf("%s", msgDir)
		return text_in_window, errDir
	} else {
		log.Println("Успішно перевірено існування директорії: ", dirPPD)
	}

	filename := filepath.Join(dirPPD, timeStamp+"_"+filepath.Base(text_in_window))
	if err := xlsx.SaveAs(filename); err != nil {
		log.Println("Помилка збережання даних в xlsx файл: ",err)
		return filename, err
	}
	return filename, nil
}

func UpdateDistributionBO() {
	fmt.Println("UpdateDistributionBO() called")
}

func SaveVacationReport1() {
	fmt.Println("SaveVacationReport1() called")
}
