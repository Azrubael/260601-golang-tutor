package gio_win

import (
	"fmt"
	"testing"
)


func TestReadShpkFile_BadPath(t *testing.T) {
    _, err := ReadShpkFile("d:/tmp/not-exist.xlsx")
    if err == nil {
        t.Fatalf("Очікувана помилка читання файлу, якого немає, але насправді він є!")
    }
}

func TestReadShpkFile_Success(t *testing.T) {
    _, err := ReadShpkFile("d:/tmp/ШПС-T0320.xlsx")
    if err != nil {
        t.Fatalf("Для тестування потрібен файл 'd:/tmp/ШПС-T0320_.xlsx', але його немає!")
    }
}

// TestReadShpkFile - тестування функції ReadShpkFile
func TestReadShpkFile(t *testing.T) {
    path := "d:/tmp/ШПС-T0320_.xlsx"
    shpk, err := ReadShpkFile(path)
    if err != nil {
        t.Fatalf("Помилка читання ШПС в файлі %s: %v\n", path, err)
    }

    lenShpk := len(shpk)
    if lenShpk == 0 {
        t.Fatalf("Несподівано з файлу %s отримана пуста мапа\n", path)
    } else {
      fmt.Printf("З файлу %s успішно зчитано дані для %d людей:\n", path, len(shpk))
    }

    shpkKeys := make([]string, 0, lenShpk)
    for k := range shpk {
      shpkKeys = append(shpkKeys, k)
    //   fmt.Println(k)
    }

		for _, name := range shpkKeys {
      person := shpk[name]
			if person.Department == "" {
        fmt.Printf("Для %s відсутні відсутні дані стосовно підрозділу", name )
      }
      if person.Rank == "" {
        fmt.Printf("Для %s відсутні відсутні дані стосовно підрозділу", name )
        }
    }
}

func TestPrepareReportPPD(t *testing.T, ) {
    path := "d:/tmp/ШПС-T0320_.xlsx"
    shpk, err := ReadShpkFile(path)
    if err != nil {
        t.Fatalf("Помилка читання ШПС в файлі %s: %v\n", path, err)
    }

    counter, _, err_msg := PrepareReportPPD(shpk)
    if err_msg != nil {
        fmt.Printf("Помилка створення звіту: %v\n", err)
    }
    l := len(list_for_ppd_report)
    fmt.Println("=== Report ===")
    fmt.Printf("Кількість людей в ШПС: %d\n", len(shpk))
	for key, d := range counter {
			fmt.Printf("%q: Offi=%d Serg=%d Sold=%d Total=%d\n",
					key, d.Offi, d.Serg, d.Sold, d.Total)
	}
    fmt.Printf("Кількість людей після обробки: %d\n", counter[list_for_ppd_report[l-1]].Total)
    // fmt.Println(report)

}