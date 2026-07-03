package gio_win

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"github.com/xuri/excelize/v2"
)

// C та D - скорочення для layout.Context та layout.Dimensions
type C = layout.Context
type D = layout.Dimensions

// Оголошення інтерфейсу для даних, що мають бути записані в *.xlsx файл
type cellValue interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 |
		~string
}

// BtnState - структура для зберігання стану кнопок та прапорців
type BtnState struct {
	file_btn, action_btn, help_btn,
	shpk_btn, proto_distrib_btn,
	prep_shpk_btn, prep_ppd_btn, refresh_distrib_btn, save_vacation_btn *widget.Clickable
	open_file, open_action, open_help,
	define_shpk, define_distrib,
	prepare_shpk, prepare_ppd, refresh_distrib, save_vacation bool
}

// Distribution - Структура для зберігання даних про розподіл особового складу в підрозділах по категоріях
type Distribution struct {
	Offi  int
	Serg  int
	Sold  int
	Total int
}

// ShortPersData - Структура для зберігання скорочених даних про військовослужбовця (для звіту ППД)
type ShortPersData struct {
	Name       string
	Department string
	Rank       string
	Company    string
}

// Person - Структура для зберігання даних ШПК про військовослужбовця
type Person struct {
	Department   string // Підрозділ
	Platoon      string // Взвод
	Company      string // Рота
	Rank         string // Звання
	Assignment   string // Відрядження
	Hospital     string // Шпиталь
	Vacation_now string // Поточна відпустка
	Study        string // Навчання
	Szch         string // СЗЧ
	Vacation1    string // Чи був у І частині щорічної відпустки
	Telephone    string // Телефон
}

// openFileNameW - Структура для виклику діалогового вікна відкриття файлу у Windows, структуру міняти не можна, бо вона відповідає C-структурі у Windows API
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
	nMaxFileTitle     uint32
	lpstrInitialDir   *uint16
	lpstrTitle        *uint16
	Flags             uint32
	nFileOffset       uint16
	nFileExtension    uint16
	lpstrDefExt       *uint16
	lCustData         uintptr
	lpfnHook          uintptr
	lpTemplateName    *uint16
}

type xlsxData struct {
	Data     *excelize.File
	FilePath string
}

var (
	// SHPK_XLSX, BO_XLSX - Змінні для зберігання зчитаних даних ШПК та розподілу в форматі Excel
	SHPK_XLSX, BO_XLSX xlsxData
	// SHPK_DATA - Підготовлені дані ШПК в форматі словника
	SHPK_DATA map[string]Person
	// PPD_COUNTER - Підготовлені дані для звіту ППД
	PPD_COUNTER map[string]Distribution
	// PPD_LIST - Списки людей по локаціям для звіту ППД
	PPD_LIST [][]ShortPersData
	// BO_COUNTER - Підготовлені дані для заповнення головної таблиці розподілу особового складу
	BO_COUNTER map[string]map[string]Distribution
)

// ppd_report_list - Перелік категорій, за якими ведеться розподіл особового складу для звіту ППД
var ppd_report_list []string = []string{
	"ППД",
	"Відпустка",
	"Шпиталь",
	"СЗЧ",
	"Відрядження",
	"Загалом",
}

// bo_report_list - Перелік категорій, за якими ведеться розподіл особового складу по підрозділам всієї частини
var bo_report_list []string = []string{
	"Загалом",
	"Відпустка",
	"Шпиталь",
	"Навчання",
	"Відрядження",
	"КСП",
	"ВОП",
	"СЗЧ",
	"ППД",
}

// comp_list - Впорядкований список підрозділів
var comp_list []string = []string{
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
