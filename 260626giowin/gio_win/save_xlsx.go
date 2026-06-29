package gio_win

import (
	"fmt"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"
)

// Функція для додавання до xlsx об'єкту цілого рядка
func SetRowValue(f *excelize.File, sheet string, row int, values []any) (int, error) {
	for i, v := range values {
		col, _ := excelize.ColumnNumberToName(i + 1)
		cell := fmt.Sprintf("%s%d", col, row)
		if err := f.SetCellValue(sheet, cell, v); err != nil {
			return row, err
		}
	}
	row++
	return row, nil
}

func SaveReportPPD(ppd_counter_ptr *map[string]Distribution,
	ppd_list_ptr *[][]ShortPersData, text_in_window string) error {

	fmt.Println("SaveReportPPD() called")
	currentTime := time.Now().Format("02.01.2006")

	xlsx := excelize.NewFile()
	sheetName := xlsx.GetSheetName(xlsx.GetActiveSheetIndex())

	rowPtr := 1
	tableHat := []string{"Призначення", "Офіцери", "Сержанти", "Солдати", "Загалом"}

	// Функція - помічник для спрощення синтаксису при додаванні клітинок
	setCell := func(i, row int, v any) error {
		col, _ := excelize.ColumnNumberToName(i + 1)
		cell := fmt.Sprintf("%s%d", col, row)
    return xlsx.SetCellValue(sheetName, cell, v)
	}

	sheetTitle := fmt.Sprintf("Розподіл особового складу 3бо станом на %v", currentTime)
	if err := setCell(1, rowPtr, sheetTitle); err != nil {
		return err
	}
	rowPtr++

	for i, element := range tableHat {
		if err := setCell(i, rowPtr, element); err != nil {
			return err
		}
	}
	rowPtr++

	// Створюємо таблицю лічильників розподілу
	reportData := [][]string{}
	for _, element := range ppd_report_list {
		dist := (*ppd_counter_ptr)[element]
		content := []string{
			element,
			fmt.Sprintf("%d", dist.Offi),
			fmt.Sprintf("%d", dist.Serg),
			fmt.Sprintf("%d", dist.Sold),
			fmt.Sprintf("%d", dist.Total)}
		reportData = append(reportData, content)
	}

	// Додаємо до об'єкту xlsx таблицю лічильників розподілу
	for _, dataRow := range reportData {
		for j, dataCell := range dataRow {
			if err := setCell(j, rowPtr, dataCell); err != nil {
				return err
			}
			rowPtr++
		}
	}

	// Створюємо таблицю списку особового складу за призначенням
	reportData = [][]string{}
	for r, assignment := range *ppd_list_ptr {
		// Створюємо заголовок призначення особового складу
		dataRow := []string{"","","","", ppd_report_list[r]}
		for i, element := range dataRow {
			if err := setCell(i, rowPtr, element); err != nil {
				return err
			}
		}
		// Додаєм заголовок призначення до таблиці reportData
		reportData = append(reportData, dataRow)
		rowPtr++
		for i, person := range assignment {
			dataRow = []string{"","","","","", strconv.Itoa(i),
			person.Rank, person.Name, person.Department}
			for j, dataCell := range dataRow {
				if err := setCell(j, rowPtr, dataCell); err != nil {
					return err
				}
			}
			rowPtr++
		}
	}
	return nil
}

func UpdateDistributionBO() {
	fmt.Println("UpdateDistributionBO() called")
}

func SaveVacationReport1() {
	fmt.Println("SaveVacationReport1() called")
}
