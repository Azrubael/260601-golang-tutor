package gio_win_test

import (
	"fmt"
	"testing"

	"github.com/Azrubael/260601-golang-tutor/260718giowin/gio_win"
)

func readShpkFile(t *testing.T, filepath string) (
	*map[string]gio_win.Person, int, error) {
	t.Helper()
	shpkPtr, err := gio_win.OpenFileXlsx("Тестовий текст TestReadShpkFile", filepath)
	if err != nil {
		return nil, 0, fmt.Errorf("Помилка читання ШПС з файлу %s:\n%v", filepath, err)
	}
	shpkDataPtr, err := gio_win.ReadShpkData(shpkPtr)
	if err != nil {
		return nil, 0, fmt.Errorf("Помилка парсинга даних ШПС в файлі %s:\n%v", shpkPtr.FilePath, err)
	}
	lenShpk := len(*shpkDataPtr)
	if lenShpk == 0 {
		return nil, 0, fmt.Errorf("Несподівано з файлу ШПС %s отримана пуста мапа\n", filepath)
	}
	fmt.Printf("З файлу %s успішно зчитано дані для %d людей:\n", filepath, len(*shpkDataPtr))
	return shpkDataPtr, lenShpk, nil
}


// TestReadShpkFile - тестування функції ReadShpkFile
func TestReadShpkFile(t *testing.T) {
	filepath := "d:/tmp/ШПС-T0320_.xlsx"
	shpkDataPtr, lenShpk, err := readShpkFile(t, filepath)
	if err != nil {
		t.Fatalf("Помилка читання %s:\n%v", filepath, err)
	}

	shpkKeys := make([]string, 0, lenShpk)
	for k := range *shpkDataPtr {
		shpkKeys = append(shpkKeys, k)
		//   fmt.Println(k)
	}

	for _, name := range shpkKeys {
		person := (*shpkDataPtr)[name]
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
	shpkDataPtr, _, err := readShpkFile(t, filepath)
	if err != nil {
		t.Fatalf("Помилка читання %s:\n%v", filepath, err)
	}

	counter, _, err_msg := gio_win.PrepareReportPPD(shpkDataPtr)
	if len(err_msg) != 0 {
		fmt.Printf("Помилка створення звіту: %v\n", err)
	}
	l := len(gio_win.PPD_report_list)
	fmt.Println("=== Report ===")
	fmt.Printf("Кількість людей в ШПС: %d\n", len(*shpkDataPtr))
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
	shpkDataPtr, _, err := readShpkFile(t, filepath)
	if err != nil {
		t.Fatalf("Помилка читання %s:\n%v", filepath, err)
	}

	boReportCounter, err_count := gio_win.PrepareReportBO(shpkDataPtr)
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

// TestPrepareVacationReport - тестування функції PrepareVacationReport
func TestPrepareVacationReport(t *testing.T) {
	filepath := "d:/tmp/ШПС-T0320_.xlsx"
	shpkDataPtr, _, err := readShpkFile(t, filepath)
	if err != nil {
		t.Fatalf("Помилка читання %s:\n%v", filepath, err)
	}

	vac1, err_count := gio_win.PrepareVacationReport(shpkDataPtr)
	if len(err_count) != 0 {
		t.Fatalf("Помилка обробки даних для звіту по відпусткам:\n%v", err)
	}

	for i, row := range *vac1 {
		switch i {
		case 0:
			fmt.Println(row[0])
		case 1:
			fmt.Printf("%-14s %-14s\t%-14s\t%s\n", row[0], row[1], row[2], row[3])
			for j := range 55 {
				if j < 54 {
					fmt.Print("-")
				} else {
					fmt.Println("-")
				}
			}
		default:
			fmt.Printf("%-14s: %-14s\t%-14s\t%s%%\n", row[0], row[1], row[2], row[3])
		}
	}
}
