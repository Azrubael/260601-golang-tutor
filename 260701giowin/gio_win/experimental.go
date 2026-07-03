package gio_win

import (
	"fmt"
	"log"
	"strings"

	"github.com/xuri/excelize/v2"
)

// SetRowValue - Функція для додавання до xlsx об'єкту рядка з масиву цілих чисел
func SetRowValue(f *excelize.File, sheet string,
	row int, colOffcet int, values []string) error {
	for i, v := range values {
		col, _ := excelize.ColumnNumberToName(i + colOffcet)
		cell := fmt.Sprintf("%s%d", col, row)
		if err := f.SetCellValue(sheet, cell, v); err != nil {
			log.Fatal(err)
			return err
		}
	}
	return nil
}


// SetRowHeightXlsx - Встановлює висоту рядка в файлі excelize.File, зберігши інші властивості
func SetRowHeightXlsx(f *excelize.File, sheet string, row int, height float64, txt string) error {
	wrap_lines := len(strings.Fields(txt))
	if wrap_lines == 1 {
		return nil
	}
	required_height := (float64(wrap_lines)) * height
	err_height := f.SetRowHeight(sheet, row, required_height)

	return err_height
}
