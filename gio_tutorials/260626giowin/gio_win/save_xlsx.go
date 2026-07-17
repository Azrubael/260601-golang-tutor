package gio_win

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

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
	ppd_list_ptr *[][]ShortPersData, text_in_window string) error {

	fmt.Println("SaveReportPPD() called")

	// Таблиця звіту, з которої буде створений xlsx файл
	reportData := [][]string{}
	tableHat := []string{
		"Призначення", "Офіцери", "Сержанти", "Солдати", "Загалом",
		"","","","",
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
			"","","","",
		}
		reportData = append(reportData, dataRow)
	}

	// Створюємо список особового складу за призначенням
	for r, assignment := range *ppd_list_ptr {
		dataRow := []string{"","","","", ppd_report_list[r]}
		reportData = append(reportData, dataRow)
		for i, person := range assignment {
			dataRow = []string{"","","","","", strconv.Itoa(i),
				person.Rank, person.Name, person.Department,
			}
			reportData = append(reportData, dataRow)
		}
	}

	// Створюємо об'єкт xlsx і додаємо до нього дані
	xlsx := excelize.NewFile()
	sheetName := xlsx.GetSheetName(xlsx.GetActiveSheetIndex())
	now := time.Now()
	dateTime := now.Format("02.01.2006")
	sheetTitle := fmt.Sprintf("Розподіл особового складу 3бо станом на %v", dateTime)
	if errCell := xlsx.SetCellValue(sheetName, "A1", sheetTitle); errCell != nil{
		log.Fatal(errCell)
		return errCell
	}

	for row, dataRow := range reportData {
		if errRow := SetRowValue(xlsx, sheetName, row+1, dataRow); errRow != nil {
			log.Fatal(errRow)
			return errRow
		}
	}

	// Зберігаємо дані в файл
	timeStamp := now.Format("060102")
	dirPPD := filepath.Dir(text_in_window)
	if _, err := os.Stat(dirPPD); os.IsNotExist(err) {
		log.Printf("Такої директорії не існує: %s\n", dirPPD)
	} else {
		log.Printf("Успішно перевірено існування директорії: %s\n", dirPPD)
	}
	filename := filepath.Join(dirPPD, timeStamp + "_" + filepath.Base(text_in_window))
	if err := xlsx.SaveAs(filename); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func UpdateDistributionBO() {
	fmt.Println("UpdateDistributionBO() called")
}

func SaveVacationReport1() {
	fmt.Println("SaveVacationReport1() called")
}
