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

