package gio_win_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Azrubael/260601-golang-tutor/260718giowin/gio_win"
)

// TestUpdateReportBO - Тест запису даних в новий файл після оновлення загального розподілу особового складу
func TestUpdateReportBO(t *testing.T) {
	filepath := "d:/tmp/ШПС-T0320_.xlsx"
	shpkDataPtr, _, err := readShpkFile(t, filepath)
	if err != nil {
		t.Fatalf("Помилка читання %s:\n%v", filepath, err)
	}

	boReportCounter, err_count := gio_win.PrepareReportBO(shpkDataPtr)
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

// TestSaveVacationReport - Тест запису звіту по відпусткам
func TestSaveVacationReport(t *testing.T) {
	filepath := "d:/tmp/ШПС-T0320_.xlsx"
	shpkDataPtr, _, err := readShpkFile(t, filepath)
	if err != nil {
		t.Fatalf("Помилка читання %s:\n%v", filepath, err)
	}

	vacReportPtr, err_count := gio_win.PrepareVacationReport(shpkDataPtr)
	if len(err_count) != 0 {
		t.Fatalf("Помилка обробки даних для загального розподілу підрозділу: %v\n",
		err)
	}

	vacReportFile := "d:/tmp/відпустки_тест.xlsx"
	savedFile, err_save := gio_win.SaveVacationReport(vacReportPtr, vacReportFile)
	if err_save != nil {
		t.Fatalf("Помилка запису звіту про відпустки до файлу %s:\n%v\n",
		savedFile, err)
	} else {
		fmt.Println("Звіт про відпустки можна прочитати в файлі:\n",
		savedFile)
	}
}