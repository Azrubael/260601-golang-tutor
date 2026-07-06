package gio_win

import (
	"fmt"
	"log"
	"strings"

	"github.com/xuri/excelize/v2"
)

/* В даному файлі зібрано кілька функцій, з якими припускаю експериментувати
 * після завершення основного об'єму роботи на проекті
 * імплементація MakeListOfCompaniesPtr і IncRankCountPtr дасть певну оптимізацію
 */

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

// MakeListOfCompaniesPtr - Створення списку вказівника для мапи підрозділів з нульовими даними розподілу int(0)
func MakeListOfCompaniesPtr(list []string) *map[string]Distribution {
	companyDist := make(map[string]Distribution, len(list))
	for _, name := range list {
		companyDist[name] = Distribution{}
	}
	return &companyDist
}

// incRankCount_ptr -інкрементує відповідні лічильники в структурі Disribution,
// використовуючи як інтерфейс *Distribution
func (dist *Distribution) IncRankCountifc(rank string) {
	getRankCategory := ""
	switch true {
	case strings.HasSuffix(rank, "олдат"):
		getRankCategory = "Sold"
	case strings.HasSuffix(rank, "ержант"):
		getRankCategory = "Serg"
	default:
		getRankCategory = "Offi"
	}
	switch getRankCategory {
	case "Sold":
		dist.Sold++
	case "Serg":
		dist.Serg++
	case "Offi":
		dist.Offi++
	}
	dist.Total++
}

// CategorizePPDifc - додає персону до відповідного списку і збільшує
// лічильники, використовуючи як інтерфейс *Distribution
func (dist *Distribution) CategorizePPDifc(
	person ShortPersData,
	list *[]ShortPersData,
	) {
	*list = append(*list, person)
	dist.IncRankCountifc(person.Rank)
}
