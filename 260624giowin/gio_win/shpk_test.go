package gio_win

import (
	"fmt"
	"testing"

	"github.com/xuri/excelize/v2"
)

// TestReadShpkFile - тестування функції ReadShpkFile
func TestReadShpkFile(t *testing.T) {
    path := "d:/tmp/ШПС-T0320_.xlsx"
    shpk, err := ReadShpkData(path)
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
    shpk, err := ReadShpkData(path)
    if err != nil {
        t.Fatalf("Помилка читання ШПС в файлі %s: %v\n", path, err)
    }

    counter, _, err_msg := PrepareReportPPD(shpk)
    if len(err_msg) != 0 {
        fmt.Printf("Помилка створення звіту: %v\n", err)
    }
    l := len(ppd_report_list)
    fmt.Println("=== Report ===")
    fmt.Printf("Кількість людей в ШПС: %d\n", len(shpk))
	for key, d := range counter {
			fmt.Printf("%q: Offi=%d Serg=%d Sold=%d Total=%d\n",
					key, d.Offi, d.Serg, d.Sold, d.Total)
	}
    fmt.Printf("Кількість людей після обробки: %d\n", counter[ppd_report_list[l-1]].Total)
    // fmt.Println(report)

}

func TestPrepareReportBO(t *testing.T, ) {
	default_path := "d:/tmp/ШПС-T0320_.xlsx"
	shpk, err := ReadShpkData(default_path)
	if err != nil {
			t.Fatalf("Помилка читання ШПС в файлі %s: %v\n", default_path, err)
	}
	boReportCounter, err_count := PrepareReportBO(shpk)
	if len(err_count) != 0 {
			t.Fatalf("Помилка обробки даних для загального розподілу підрозділу: %v\n", err)
	}

	for i:= 0; i < len(bo_report_list); i++ {
		fmt.Println("=== Report ===", bo_report_list[i])
		for _, c := range comp_list {
			fmt.Println(c, "\t", boReportCounter[c][bo_report_list[i]])
		}
	}
}

// TestExperimentBO - тестування функції ExperimentOpenXlsx в файлі tmp.go
func TestExperimentOpenXlsx(t *testing.T) {
    title := "Виберіть Excel файл"
    filterPairs := []string{
		"Excel files (*.xlsx)", "*.xlsx",
		"All files (*.*)", "*.*",
	}

	xlsx_file, err_xlsx := ExperimentOpenXlsx(title, filterPairs)
	if err_xlsx != nil {
		t.Fatalf("Помилка спроби відкриття фійлу *.xlsx за допомогою функції ExperimentOpenXlsx: %v\n", err_xlsx)
	}

	if xlsx_file != "" {
		xlsx_data, err_shpk := excelize.OpenFile(xlsx_file)
		if err_shpk != nil {
			t.Fatalf("Помилка відкриття %s: %v", xlsx_file, err_shpk)
		}
		fmt.Printf("Успішно відкрито файл %s, кількість аркушів: %d\n", xlsx_file, len(xlsx_data.GetSheetList()))
	} else {
		fmt.Println("Користувач скасував вибір файлу *.xlsx за допомогою функції ExperimentOpenXlsx")
	}
}

// TestExperimentBO - тестування функції ExperimentOpenXlsx в файлі tmp.go
func TestOpenFileXlsx(t *testing.T) {
    title := "Виберіть Excel файл"

	xlsx_data, shpk_file_path, err_xlsx := OpenFileXlsx(title, "d:/tmp/ШПС-T0320.xlsx")
	if err_xlsx != nil {
		t.Fatalf("Помилка спроби відкриття фійлу %s за допомогою функції OpenFileXlsx: %v\n",
        shpk_file_path, err_xlsx)
	}

	if xlsx_data != nil {
		fmt.Printf("Успішно відкрито файл %s, кількість аркушів: %d\n", shpk_file_path, len(xlsx_data.GetSheetList()))
	} else {
		fmt.Println("Користувач скасував вибір файлу *.xlsx за допомогою функції ExperimentOpenXlsx")
	}
}