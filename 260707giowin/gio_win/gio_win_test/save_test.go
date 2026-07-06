package gio_win_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Azrubael/260601-golang-tutor/260707giowin/gio_win"
)

// TestUpdateReportBO - Тест запису даних в новий файл після оновлення загального розподілу особового складу
func TestUpdateReportBO(t *testing.T) {
	shpk_filepath := "d:/tmp/ШПС-T0320_.xlsx"
	shpkPtr, err := gio_win.OpenFileXlsx("Тестовий текст TestReadShpkFile", shpk_filepath)
	if err != nil {
		t.Fatalf("Помилка читання ШПС в файлі %s: %v\n", shpk_filepath, err)
	}

	shpk_table, err := gio_win.ReadShpkData(shpkPtr)
	if err != nil {
		t.Fatalf("Помилка читання ШПС в файлі %s: %v\n", shpkPtr.FilePath, err)
	}

	boReportCounter, err_count := gio_win.PrepareReportBO(shpk_table)
	if len(err_count) != 0 {
		t.Fatalf("Помилка обробки даних для загального розподілу підрозділу: %v\n",
		err)
	}

	title_bo := "Виберіть Excel файл загального розподілу людей"
	bo_xlsx_proto := "D:/Документи/III БАТ/склад 3 БО/260701-3бо.xlsx"
	boXlsxPtr, err_bo := gio_win.OpenFileXlsx(title_bo, bo_xlsx_proto)
	if err_bo != nil {
		msg := fmt.Sprintf("Помилка відкриття %s з даними розподілу людей",
			boXlsxPtr.FilePath)
		t.Fatalf("%s: %v\n", msg, err_bo)
	} else {
		fmt.Printf("Прочитаний файл %s містить %s\n", boXlsxPtr.FilePath,
		reflect.TypeOf(*boXlsxPtr))
	}

	savedFile, err_save := gio_win.UpdateDistributionBO(boReportCounter, boXlsxPtr, "")
	if err_save != nil {
		t.Fatalf("Помилка запису оновлених даних загального розподілу до файлу %s:\n%v\n",
		savedFile, err)
	} else {
		fmt.Println("Дані успішно оновленого розподілу можна прочитати в файлі:\n",
		savedFile)
	}
}