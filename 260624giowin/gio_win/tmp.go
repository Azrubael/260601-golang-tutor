package gio_win

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

// ExperimentBO - Експериментальна функція для створення структури звіту по підрозділах частини
func ExperimentBO(shpk_data map[string]Person) (map[string]map[string]Distribution, []string) {

	type ReportData struct {
		Name       string
		Department string
		Rank       string
	}

	var bo []string = []string{
		"Відпустка",
		"Шпиталь",
		"Навчання",
		"Відрядження",
		"ВОП",
		"СЗЧ",
		"ППД",
		"Загалом",
	}

	var companies []string = []string{
		"упр 3 бо",
		"1",
		"2",
		"3",
		"4",
		"від.зв./3 бо",
		"від.заб./3 бо",
		"від.то/3 бо",
		"м.п./3 бо",
		"підсумок",
	}

	boReportCounter := make(map[string]map[string]Distribution, len(companies))
	count_err := []string{}
	for _, c := range companies {
		for _, b := range bo {
			boReportCounter[c][b] = Distribution{}
		}
	}

	return boReportCounter, count_err
}

func ExperimentOpenXlsx(title string, filterPairs []string) (string, error) {
	// filterPairs: [ "Text files (*.txt)", "*.txt", "All files (*.*)", "*.*" ]
	const (
		OFN_FILEMUSTEXIST = 0x00001000
		OFN_PATHMUSTEXIST = 0x00000800
		MAX_PATH          = 260
	)

	type openFileNameW struct {
		lStructSize       uint32
		hwndOwner         uintptr
		hInstance         uintptr
		lpstrFilter       *uint16
		lpstrCustomFilter *uint16
		nMaxCustFilter    uint32
		nFilterIndex      uint32
		lpstrFile         *uint16
		nMaxFile          uint32
		lpstrFileTitle    *uint16
		nMaxFileTitle    uint32
		lpstrInitialDir   *uint16
		lpstrTitle        *uint16
		Flags             uint32
		nFileOffset       uint16
		nFileExtension    uint16
		lpstrDefExt       *uint16
		lCustData         uintptr
		lpfnHook          uintptr
		lpTemplateName   *uint16
	}
	modComdlg32 := windows.NewLazySystemDLL("comdlg32.dll")
	procGetOpenFileNameW := modComdlg32.NewProc("GetOpenFileNameW")
	// COMDLG32 очікує рядок: "desc\0pattern\0desc\0pattern\0\0"
	var filters16 []uint16

	for i := 0; i+1 < len(filterPairs); i += 2 {
		a, err := windows.UTF16FromString(filterPairs[i])
		if err != nil {
			return "", err
		}
		filters16 = append(filters16, a...)
		filters16 = append(filters16, 0)

		b, err := windows.UTF16FromString(filterPairs[i+1])
		if err != nil {
			return "", err
		}
		filters16 = append(filters16, b...)
		filters16 = append(filters16, 0)
	}
	filters16 = append(filters16, 0)

	title16, err := windows.UTF16FromString(title)
	if err != nil {
		return "", err
	}

	fileBuf := make([]uint16, MAX_PATH)
	ofn := openFileNameW{
		lStructSize:  uint32(unsafe.Sizeof(openFileNameW{})),
		// hwndOwner/hInstance/other unused fields: 0
		lpstrFilter:  &filters16[0],
		lpstrFile:    &fileBuf[0],
		nMaxFile:     MAX_PATH,
		lpstrTitle:   &title16[0],
		Flags:        OFN_FILEMUSTEXIST | OFN_PATHMUSTEXIST,
		nFilterIndex: 1,
	}

	ret, _, callErr := procGetOpenFileNameW.Call(uintptr(unsafe.Pointer(&ofn)))
	if ret == 0 {
		// скасовано
		if callErr == windows.ERROR_SUCCESS {
			return "", nil
		}
		// помилка
		if callErr != nil && callErr != windows.ERROR_SUCCESS {
			return "", callErr
		}
		return "", nil
	}

	// перетворюємо UTF-16 до Go рядка (до першого 0)
	n := 0
	for n < len(fileBuf) && fileBuf[n] != 0 {
		n++
	}
	path := syscall.UTF16ToString(fileBuf[:n])
	return path, nil
}