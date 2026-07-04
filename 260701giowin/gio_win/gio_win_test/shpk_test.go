package gio_win_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Azrubael/260601-golang-tutor/260701giowin/gio_win"
)

// TestReadShpkFile - тестування функції ReadShpkFile
func TestReadShpkFile(t *testing.T) {
	filepath := "d:/tmp/ШПС-T0320_.xlsx"
	shpk, err := gio_win.OpenFileXlsx("Тестовий текст TestReadShpkFile", filepath)
	if err != nil {
		t.Fatalf("Помилка читання ШПС в файлі %s: %v\n", filepath, err)
	} else {
		fmt.Printf("Прочитаний файл %s містить %s\n", filepath, reflect.TypeOf(shpk))
	}

	shpk_table, err := gio_win.ReadShpkData(&shpk)
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
	shpk, err := gio_win.OpenFileXlsx("Тестовий текст TestReadShpkFile", filepath)
	if err != nil {
		t.Fatalf("Помилка читання ШПС в файлі %s: %v\n", filepath, err)
	} else {
		fmt.Printf("Прочитаний файл %s містить %s\n", filepath, reflect.TypeOf(shpk))
	}

	shpk_table, err := gio_win.ReadShpkData(&shpk)
	if err != nil {
		t.Fatalf("Помилка читання ШПС в файлі %s: %v\n", shpk.FilePath, err)
	}
	counter, _, err_msg := gio_win.PrepareReportPPD(shpk_table)
	if len(err_msg) != 0 {
		fmt.Printf("Помилка створення звіту: %v\n", err)
	}
	l := len(gio_win.PPD_report_list)
	fmt.Println("=== Report ===")
	fmt.Printf("Кількість людей в ШПС: %d\n", len(shpk_table))
	for key, d := range counter {
		fmt.Printf("%-15s: Offi=%-7d\tSerg=%-7d\tSold=%-7d\tTotal=%-7d\n",
			key, d.Offi, d.Serg, d.Sold, d.Total)
	}
	fmt.Printf("Кількість людей після обробки: %d\n",
	counter[gio_win.PPD_report_list[l-1]].Total)
}

// TestPrepareReportBO - тестування функції PrepareReportBO
func TestPrepareReportBO(t *testing.T) {
	filepath := "d:/tmp/ШПС-T0320_.xlsx"
	shpk, err := gio_win.OpenFileXlsx("Тестовий текст TestReadShpkFile", filepath)
	if err != nil {
		t.Fatalf("Помилка читання ШПС в файлі %s:\n%v", filepath, err)
	} else {
		fmt.Printf("Прочитаний файл %s містить %s\n", filepath, reflect.TypeOf(shpk))
	}
	shpk_table, err := gio_win.ReadShpkData(&shpk)
	if err != nil {
		t.Fatalf("Помилка читання ШПС в файлі %s:\n%v", shpk.FilePath, err)
	}
	boReportCounter, err_count := gio_win.PrepareReportBO(shpk_table)
	if len(err_count) != 0 {
		t.Fatalf("Помилка обробки даних для загального розподілу підрозділу:\n%v", err)
	}

	for i := range gio_win.BO_report_list {
		fmt.Println("=== Report ===", gio_win.BO_report_list[i])
		for _, c := range gio_win.COMP_list {
			fmt.Println(c, "\t", boReportCounter[c][gio_win.BO_report_list[i]])
		}
	}
}

// TestPrepareVacationReport1
func TestPrepareVacationReport1(t *testing.T) {
		filepath := "d:/tmp/ШПС-T0320_.xlsx"
	shpk, err := gio_win.OpenFileXlsx("Тестовий текст TestReadShpkFile", filepath)
	if err != nil {
		t.Fatalf("Помилка читання ШПС в файлі %s: %v\n", filepath, err)
	} else {
		fmt.Printf("Прочитаний файл %s містить %s\n", filepath, reflect.TypeOf(shpk))
	}
	shpk_table, err := gio_win.ReadShpkData(&shpk)
	if err != nil {
		t.Fatalf("Помилка читання ШПС в файлі %s: %v\n", shpk.FilePath, err)
	}

	vac1, err_count := gio_win.PrepareVacationReport1(shpk_table)
	if len(err_count) != 0 {
		t.Fatalf("Помилка обробки даних для звіту по відпусткам:\n%v", err)
	}

	for _, row := range vac1 {
		fmt.Printf("%-14s: %-14s\t%-14s\t%s%%\n", row[0], row[1], row[2], row[3])
	}
}
