package gio_win

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUpdateReportBO(t *testing.T) {
	shpk_filepath := "d:/tmp/ШПС-T0320_.xlsx"
	shpk, err := OpenFileXlsx("Тестовий текст TestReadShpkFile", shpk_filepath)
	if err != nil {
		t.Fatalf("Помилка читання ШПС в файлі %s: %v\n", shpk_filepath, err)
	}

	shpk_table, err := ReadShpkData(&shpk)
	if err != nil {
		t.Fatalf("Помилка читання ШПС в файлі %s: %v\n", shpk.FilePath, err)
	}

	boReportCounter, err_count := PrepareReportBO(shpk_table)
	if len(err_count) != 0 {
		t.Fatalf("Помилка обробки даних для загального розподілу підрозділу: %v\n",
		err)
	}

	title_bo := "Виберіть Excel файл загального розподілу людей"
	bo_xlsx_proto := "D:/Документи/III БАТ/склад 3 БО/260701-3бо.xlsx"
	bo_xlsx_test, err_bo := OpenFileXlsx(title_bo, bo_xlsx_proto)
	if err_bo != nil {
		msg := fmt.Sprintf("Помилка відкриття %s з даними розподілу людей",
			bo_xlsx_test.FilePath)
		t.Fatalf("%s: %v\n", msg, err_bo)
	} else {
		fmt.Printf("Прочитаний файл %s містить %s\n", bo_xlsx_test.FilePath,
		reflect.TypeOf(bo_xlsx_test))
	}


	savedFile, err_save := UpdateDistributionBO(boReportCounter, &bo_xlsx_test, "")
	if err_save != nil {
		t.Fatalf("Помилка запису оновлених даних загального розподілу до файлу %s:\n%v\n",
		savedFile, err)
	} else {
		fmt.Println("Дані успішно оновленого розподілу можна прочитати в файлі:\n",
		savedFile)
	}
}