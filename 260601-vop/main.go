package main

/* Cкріпт, який робить веріфікацію даних персоналу, і виконує розрахунок
 * кількості особового складу по категоріям: звання і підрозділи
 */

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
)

type Person struct {
	Department string
	Rank       string
	Telephone  string
}

// Очистка імени від зайвих символів
func cleanName(name string) string {
	if name == "" {
		return ""
	}
	return strings.TrimSpace(strings.ReplaceAll(name, "\n", " "))
}

// Скорочення напису звання
func shortenRank(rank string) (string, string) {
	if rank == "" {
		return "", "Звання відсутнє"
	}
	parts := strings.FieldsFunc(rank, func(r rune) bool {
		return r == ' ' || r == '\t'
	})

	if len(parts) >= 3 {
		return "", "Забагато слів у званні"
	}
	if len(parts) == 1 {
		return rank, ""
	}
	replacements := map[string]string{
		"Старший":  "Ст.",
		"Молодший": "Мол.",
		"Головний": "Гол.",
	}
	first, second := parts[0], parts[1]
	if s, ok := replacements[first]; ok {
		return s + second, ""
	} else {
		return rank, "Помилка в званні"
	}

}

// Отримати координати ячейки в файлі Excel
func coord(col, row int) string {
	colName, _ := excelize.ColumnNumberToName(col)
	return fmt.Sprintf("%s%d", colName, row)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Використання: go run %s <vop.xlsx>\n", filepath.Base(os.Args[0]))
		fmt.Printf("\tПримітка: Файл <vop.xlsx> має бути в папці d:/tmp/250706-vop/")
		os.Exit(1)
	}
	prefix_path := "d:/tmp/"
	work_dir := "250706-vop/"
	shpk := "ШПС-T0320.xlsx"

	shpk_file := prefix_path + shpk
	vop_file := prefix_path + work_dir + os.Args[1]

	// Відкриття файлу з ШПС в форматі Excel
	shpk_xlsx, err_shpk := excelize.OpenFile(shpk_file)
	if err_shpk != nil {
		fmt.Println(fmt.Errorf("Помилка відкриття %s: %w", shpk_file, err_shpk))
		os.Exit(2)
	}

	// Отрмання таблиці даних "ШПС" у вигляді рядків
	shpk_rows, err_shpk := shpk_xlsx.GetRows("ШПС")
	if err_shpk != nil {
		fmt.Println(fmt.Errorf("Помилка зчитування змісту %s: %w", shpk_file, err_shpk))
		os.Exit(3)
	}

	// Структура даних для персоналу
	shpk_data := make(map[string]Person)

	// Заповнення структури даних персоналу змістом, пропускаючи заголовки ШПС
	for i := 2; i < len(shpk_rows) && i < 630; i++ { // index 2 = row 3
		row := shpk_rows[i]
		if len(row) > 16 && row[8] != "" {
			cleaned_name := cleanName(row[8])
			if cleaned_name != "" {
				short_rank, err_rank := shortenRank(row[7])
				if err_rank == "" {
					shpk_data[cleaned_name] = Person{
						Department: row[10],
						Rank:       short_rank,
						Telephone:  row[16],
					}

				}
			}
		}
	}

	// Відкриття файлу з даними ВОПів в форматі Excel
	vop_xlsx, err_vopi := excelize.OpenFile(vop_file)
	if err_vopi != nil {
		fmt.Println(fmt.Errorf("Помилка відкриття %s: %w", vop_file, err_vopi))
		os.Exit(4)
	}

	// Отримання доступу до аркушів в файлі з даними ВОПів
	for _, vop_sheet := range vop_xlsx.GetSheetList() {
		// Отримання таблиці даних "ВОП" у вигляді рядків
		vop_rows, err_vopi := vop_xlsx.GetRows(vop_sheet)
		if err_vopi != nil {
			fmt.Println(fmt.Errorf("Помилка зчитування змісту %s:\n %w", vop_file, err_vopi))
			os.Exit(5)
		} else {
			message := fmt.Sprintf("Зчитано %d рядків з листа %s", len(vop_rows), vop_sheet)
			fmt.Println(message)
		}
		// Оновлення змісту в структурі даних, пропускаючи заголовки ВОП
		for vop_row := 4; vop_row < len(vop_rows); vop_row++ {
			coord_rank := coord(3, vop_row)
			coord_name := coord(4, vop_row)
			coord_dep := coord(5, vop_row)
			coord_tel := coord(6, vop_row)
			name, err_name := vop_xlsx.GetCellValue(vop_sheet, coord_name)
			if err_name != nil {
				fmt.Println(fmt.Errorf("Помилка зчитування імені: %w", err_name))
				os.Exit(10)
			}
			cleaned_name := cleanName(name)

			if person, ok := shpk_data[cleaned_name]; ok {
				err_rank := vop_xlsx.SetCellValue(vop_sheet, coord_rank, person.Rank)
				if err_rank != nil {
					fmt.Println(person.Rank)
					fmt.Println(fmt.Errorf("Помилка запису звання: %w", err_rank))
					os.Exit(12)
				}
				err_name := vop_xlsx.SetCellValue(vop_sheet, coord_name, cleaned_name)
				if err_name != nil {
					fmt.Println(cleaned_name)
					fmt.Println(fmt.Errorf("Помилка запису імені: %w", err_name))
					os.Exit(13)
				}
				err_department := vop_xlsx.SetCellValue(vop_sheet, coord_dep, person.Department)
				if err_department != nil {
					fmt.Println(fmt.Errorf("Помилка запису підрозділу: %w", err_department))
					os.Exit(14)
				}
				err_telephone := vop_xlsx.SetCellValue(vop_sheet, coord_tel, person.Telephone)
				if err_telephone != nil {
					fmt.Println(fmt.Errorf("Помилка запису телефону: %w", err_telephone))
					os.Exit(15)
				}
				// fmt.Println(person.Rank + " " + cleaned_name)

				// Переконатись, що wrap_text увімкнено, не втративши інших властивостей
				err_tel_wrap := EnsureWrapText(vop_xlsx, vop_sheet, coord_tel)
				if err_tel_wrap != nil {
					fmt.Println(fmt.Errorf("Помилка форматування ячейки для телефонного номеру: %w", err_tel_wrap))
					os.Exit(16)
				}
				err_wrap_tel := SetRowHeightXlsx(vop_xlsx, vop_sheet, vop_row, 15.0, person.Telephone)
				if err_wrap_tel != nil {
					message := fmt.Errorf("Помилка встановлення висоти рядка для %s : %w", cleaned_name, err_wrap_tel)
					fmt.Println(message)
				}

			} else if cleaned_name != "" {
				message := fmt.Sprintf("Ім'я %s не знайдено в ШПС %s", cleaned_name, shpk_file)
				fmt.Println(message)
			}

		}
	}

	current_date := time.Now().Format("060102")
	new_vop_file := prefix_path + work_dir + current_date + "-Склад_ВОПів_перевірено.xlsx"
	fmt.Printf("Дані файлу %s успішно оновлено для розрахунку розподілу персоналу.\n", new_vop_file)
	vop_xlsx.SaveAs(new_vop_file)

	divisions_list := Divisions()

	for _, vop_sheet := range vop_xlsx.GetSheetList() {
		vop_rows, err_vopi := vop_xlsx.GetRows(vop_sheet)
		if err_vopi != nil {
			fmt.Println(fmt.Errorf("Помилка зчитування змісту %s:\n %w", vop_rows, err_vopi))
			os.Exit(17)
		}
		divisions_counter := PersonnelCounter(vop_rows)
		var found_row int
		var found_col int
		for row, vop_row := range vop_rows {
			for col, vop_cell := range vop_row {
				if strings.Contains(vop_cell, "1 рота") {
					found_row = row
					found_col = col + 1
					break
				}
			}
			if found_row != 0 {
				break
			}
		}

		get_counts := func(m map[string]int, k string) int {
			if m == nil {
				return 0
			}
			return m[k]
		}

		output_strings := make(map[string]string, len(divisions_counter))
		for el, counts := range divisions_counter {
			output_strings[el] = fmt.Sprintf("%d-%d-%d",
				get_counts(counts, "offi"),
				get_counts(counts, "serg"),
				get_counts(counts, "sold"),
			)
		}

		for i, el := range divisions_list {
			vop_xlsx.SetCellValue(vop_sheet, coord(found_col+1, found_row+i), output_strings[el])
			fmt.Println(found_col+1, found_row+i, output_strings[el])
		}

		vop_xlsx.SaveAs(new_vop_file)

	}
}
