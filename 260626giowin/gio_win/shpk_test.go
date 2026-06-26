package gio_win

import (
	"fmt"
	// "path/filepath"
	"reflect"
	"testing"
)

// TestOpenShpkFile - тестування функції ReadShpkFile
func TestOpenShpkFile(t *testing.T) {
	filepath := "d:/tmp/ШПС-T0320_.xlsx"
	shpk, err := OpenFileXlsx("Тестовий текст TestOpenShpkFile", filepath)
	if err != nil {
		t.Fatalf("Помилка читання ШПС в файлі %s: %v\n", filepath, err)
	} else {
		fmt.Printf("Прочитаний файл %s містить %s\n", filepath, reflect.TypeOf(shpk))
	}

}

// TestOpenShpkFile - тестування функції ReadShpkFile
func TestReadShpkFile(t *testing.T) {
	filepath := "d:/tmp/ШПС-T0320_.xlsx"
	shpk, err := OpenFileXlsx("Тестовий текст TestReadShpkFile", filepath)
	if err != nil {
		t.Fatalf("Помилка читання ШПС в файлі %s: %v\n", filepath, err)
	} else {
		fmt.Printf("Прочитаний файл %s містить %s\n", filepath, reflect.TypeOf(shpk))
	}

	shpk_table, err := ReadShpkData(shpk)
	if err != nil {
		t.Fatalf("Помилка читання ШПС в файлі %s: %v\n", shpk.FilePath, err)
	}
	lenShpk := len(shpk_table)
	if lenShpk == 0 {
		t.Fatalf("Несподівано з файлу %s отримана пуста мапа\n", filepath)
	} else {
		fmt.Printf("З файлу %s успішно зчитано дані для %d людей:\n", filepath, len(shpk_table))
	}

	shpkKeys := make([]string, 0, lenShpk)
	for k := range shpk_table {
		shpkKeys = append(shpkKeys, k)
		//   fmt.Println(k)
	}

	for _, name := range shpkKeys {
		person := shpk_table[name]
		if person.Department == "" {
			fmt.Printf("Для %s відсутні відсутні дані стосовно підрозділу", name)
		}
		if person.Rank == "" {
			fmt.Printf("Для %s відсутні відсутні дані стосовно підрозділу", name)
		}
	}
}

func TestPrepareReportPPD(t *testing.T) {
	filepath := "d:/tmp/ШПС-T0320_.xlsx"
	shpk, err := OpenFileXlsx("Тестовий текст TestReadShpkFile", filepath)
	if err != nil {
		t.Fatalf("Помилка читання ШПС в файлі %s: %v\n", filepath, err)
	} else {
		fmt.Printf("Прочитаний файл %s містить %s\n", filepath, reflect.TypeOf(shpk))
	}

	shpk_table, err := ReadShpkData(shpk)
	if err != nil {
		t.Fatalf("Помилка читання ШПС в файлі %s: %v\n", shpk.FilePath, err)
	}
	counter, _, err_msg := PrepareReportPPD(shpk_table)
	if len(err_msg) != 0 {
		fmt.Printf("Помилка створення звіту: %v\n", err)
	}
	l := len(ppd_report_list)
	fmt.Println("=== Report ===")
	fmt.Printf("Кількість людей в ШПС: %d\n", len(shpk_table))
	for key, d := range counter {
		fmt.Printf("%q: Offi=%d Serg=%d Sold=%d Total=%d\n",
			key, d.Offi, d.Serg, d.Sold, d.Total)
	}
	fmt.Printf("Кількість людей після обробки: %d\n", counter[ppd_report_list[l-1]].Total)
	// fmt.Println(report)

}

func TestPrepareReportBO(t *testing.T) {
	filepath := "d:/tmp/ШПС-T0320_.xlsx"
	shpk, err := OpenFileXlsx("Тестовий текст TestReadShpkFile", filepath)
	if err != nil {
		t.Fatalf("Помилка читання ШПС в файлі %s: %v\n", filepath, err)
	} else {
		fmt.Printf("Прочитаний файл %s містить %s\n", filepath, reflect.TypeOf(shpk))
	}
	shpk_table, err := ReadShpkData(shpk)
	if err != nil {
		t.Fatalf("Помилка читання ШПС в файлі %s: %v\n", shpk.FilePath, err)
	}
	boReportCounter, err_count := PrepareReportBO(shpk_table)
	if len(err_count) != 0 {
		t.Fatalf("Помилка обробки даних для загального розподілу підрозділу: %v\n", err)
	}

	for i := 0; i < len(bo_report_list); i++ {
		fmt.Println("=== Report ===", bo_report_list[i])
		for _, c := range comp_list {
			fmt.Println(c, "\t", boReportCounter[c][bo_report_list[i]])
		}
	}
}

// TestExperimentBO - тестування функції ExperimentOpenXlsx в файлі tmp.go
func TestOpenFileXlsx(t *testing.T) {
	title := "Виберіть Excel файл"

	shpk, err_xlsx := OpenFileXlsx(title, "d:/tmp/ШПС-T0320.xlsx")
	if err_xlsx != nil {
		t.Fatalf("Помилка спроби відкриття фійлу %s за допомогою функції OpenFileXlsx: %v\n",
			shpk.FilePath, err_xlsx)
	}

	if shpk.Data != nil {
		fmt.Printf("Успішно відкрито файл %s, кількість аркушів: %d\n", shpk.FilePath, len(shpk.Data.GetSheetList()))
	} else {
		fmt.Println("Користувач скасував вибір файлу *.xlsx за допомогою функції ExperimentOpenXlsx")
	}
}
