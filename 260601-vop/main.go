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
        return s  + second, ""
    } else {
        return rank, "Помилка в званні"
    }

}


// Отримати координати ячейки в файлі Excel
func coord(col, row int) string {
    colName, _ := excelize.ColumnNumberToName(col)
    return fmt.Sprintf("%s%d", colName, row)
}


func ensureWrapText(f *excelize.File, sheet, coord string) error {
    // Отримати поточний styleID (0 якщо немає)
    styleID, err := f.GetCellStyle(sheet, coord)
    if err != nil {
        return err
    }

    // Якщо styleID == 0 — немає явного стилю; створимо новий із wrap_text
    if styleID == 0 {
        newStyle, err := f.NewStyle(&excelize.Style{
            Alignment: &excelize.Alignment{WrapText: true},
        })
        if err != nil {
            return err
        }
        return f.SetCellStyle(sheet, coord, coord, newStyle)
    }

    // Якщо style є — прочитаємо його як XML, змодифікуємо alignment.wrapText і створимо новий стиль, зберігши інші властивості.
    style, err := f.GetStyle(styleID)
    if err != nil {
        // Якщо не вдається прочитати повний стиль, застосуємо простий wrap-only стиль
        newStyle, err2 := f.NewStyle(&excelize.Style{
            Alignment: &excelize.Alignment{WrapText: true},
        })
        if err2 != nil {
            return err2
        }
        return f.SetCellStyle(sheet, coord, coord, newStyle)
    }

    // Модифікуємо alignment (створимо копію)
    if style.Alignment == nil {
        style.Alignment = &excelize.Alignment{}
    }
    style.Alignment.WrapText = true

    // Створимо новий стиль на основі модифікованого опису
    newStyleID, err := f.NewStyle(style)
    if err != nil {
        return err
    }

    // Застосуємо новий стиль до клітинки
    return f.SetCellStyle(sheet, coord, coord, newStyleID)
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
    vopi_file := prefix_path + work_dir + os.Args[1]

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
    vopi_xlsx, err_vopi := excelize.OpenFile(vopi_file)
    if err_vopi != nil {
        fmt.Println(fmt.Errorf("Помилка відкриття %s: %w", vopi_file, err_vopi))
        os.Exit(4)
    }

    // Отримання списку листів в файлі з даними ВОПів
    vopi_sheets := vopi_xlsx.GetSheetList()
    fmt.Println(vopi_sheets)



    for i := 1; i < len(vopi_sheets); i++ {
        vopi_sheet := vopi_sheets[i]
        // Отримання таблиці даних "ВОП" у вигляді рядків
        vopi_rows, err_vopi := vopi_xlsx.GetRows(vopi_sheet)
        if err_vopi != nil {
            fmt.Println(fmt.Errorf("Помилка зчитування змісту %s: %w", vopi_file, err_vopi))
            os.Exit(5)
        } else {
            message := fmt.Sprintf("Зчитано %d рядків з листа %s", len(vopi_rows), vopi_sheet)
            fmt.Println(message)
        }
        // Оновлення змісту в структурі даних, пропускаючи заголовки ВОП
        for vopi_row := 4; vopi_row < len(vopi_rows); vopi_row++ {
            coord_rank := coord(3, vopi_row)
            coord_name := coord(4, vopi_row)
            coord_dep := coord(5, vopi_row)
            coord_tel := coord(6, vopi_row)
            name, err_name := vopi_xlsx.GetCellValue(vopi_sheet, coord_name)
            if err_name != nil {
                fmt.Println(fmt.Errorf("Помилка зчитування імені: %w", err_name))
                os.Exit(10)
            }
            cleaned_name := cleanName(name)

            if person, ok := shpk_data[cleaned_name]; ok {
                err_rank := vopi_xlsx.SetCellValue(vopi_sheet, coord_rank, person.Rank)
                if err_rank != nil {
                    fmt.Println(person.Rank)
                    fmt.Println(fmt.Errorf("Помилка запису звання: %w", err_rank))
                    os.Exit(12)
                }
                err_name := vopi_xlsx.SetCellValue(vopi_sheet, coord_name, cleaned_name)
                if err_name != nil {
                    fmt.Println(cleaned_name)
                    fmt.Println(fmt.Errorf("Помилка запису імені: %w", err_name))
                    os.Exit(13)
                }
                err_department := vopi_xlsx.SetCellValue(vopi_sheet, coord_dep, person.Department)
                if err_department != nil {
                    fmt.Println(fmt.Errorf("Помилка запису підрозділу: %w", err_department))
                    os.Exit(14)
                }
                err_telephone := vopi_xlsx.SetCellValue(vopi_sheet, coord_tel, person.Telephone)
                if err_telephone != nil {
                    fmt.Println(fmt.Errorf("Помилка запису телефону: %w", err_telephone))
                    os.Exit(15)
                }
                fmt.Println(person.Rank + " " + cleaned_name)

                // Переконатись, що wrap_text увімкнено, не втративши інших властивостей
                err_tel_wrap := ensureWrapText(vopi_xlsx, vopi_sheet, coord_tel)
                if err_tel_wrap != nil {
                    fmt.Println(fmt.Errorf("Помилка форматування ячейки для телефонного номеру: %w", err_tel_wrap))
                    os.Exit(1)
                }


            } else if cleaned_name != "" {
                message := fmt.Sprintf("Ім'я %s не знайдено в ШПС %s", cleaned_name, shpk_file)
                fmt.Println(message)
            }
        }
    }

    current_date := time.Now().Format("060102")
    new_vopi_file := prefix_path + work_dir + current_date + "-Склад_ВОПів_перевірено.xlsx"
    fmt.Printf("Файл %s успішно записаний\n", new_vopi_file)
    vopi_xlsx.SaveAs(new_vopi_file)

}
