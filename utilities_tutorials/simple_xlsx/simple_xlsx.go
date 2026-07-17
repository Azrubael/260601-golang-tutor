package main
/* Дуже простий скріпт для створення нового файлу `output.xlsx` і заповнення його
 * даними, збереженими в 2D масиві `data`.
 */
import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

func main() {
	// 2D масив даних: кожен внутрішній slice — це один рядок
	data := [][]interface{}{
		{"ID", "Name", "Score"},
		{1, "Alice", 95.5},
		{2, "Bob", 88.0},
		{3, "Carol", 91.2},
	}

	// 1) Створюємо нову книгу
	f := excelize.NewFile()

	// 2) Обираємо активний аркуш
	sheetName := f.GetSheetName(f.GetActiveSheetIndex())

	// 3) Додаємо дані рядок за рядком
	for r, row := range data {
		for c, cellValue := range row {
			// excel координати: (1-based)
			axis, err := excelize.CoordinatesToCellName(c+1, r+1)
			if err != nil {
				log.Fatal(err)
			}

			// Excelize вміє ставити різні типи через SetCellValue
			// (рядки/числа/float/тощо)
			if err := f.SetCellValue(sheetName, axis, cellValue); err != nil {
				log.Fatal(err)
			}
		}
	}

	// 4) Зберігаємо у файл
	if err := f.SaveAs("output.xlsx"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Готово: output.xlsx")
}
