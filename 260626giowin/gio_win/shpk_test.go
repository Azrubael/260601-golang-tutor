package gio_win

import (
	"fmt"
	// "path/filepath"
	"reflect"
	"testing"
)

// TestOpenShpkFile - тестування функції OpenShpkFile
// func TestOpenShpkFile(t *testing.T) {
// 	filepath := ""
// 	shpk, err := OpenFileXlsx("Тестовий текст TestOpenShpkFile", filepath)
// 	if err != nil {
// 		t.Fatalf("Помилка читання ШПС в файлі %s: %v\n", filepath, err)
// 	} else {
// 		fmt.Printf("Прочитаний файл %s містить %s\n", filepath, reflect.TypeOf(shpk))
// 	}

// }

// TestReadShpkFile - тестування функції ReadShpkFile
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

// TestPrepareReportPPD - тестування функції PrepareReportPPD
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

// TestPrepareReportBO - тестування функції PrepareReportBO
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

