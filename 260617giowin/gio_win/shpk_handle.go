package gio_win

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/xuri/excelize/v2"
)

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


// ReadCellSafe - Безпечне отримання значення ячейки, з перевіркою чи вона існує
func ReadCellSafe(row []string, col int) string {
	if col < len(row) {
		return row[col]
	}
	return ""
}

// CleanName - Очистка імени від зайвих символів
func CleanName(name string) string {
	if name == "" {
		return ""
	}
	return strings.TrimSpace(strings.ReplaceAll(name, "\n", " "))
}

// IsShooter - Перевірка, чи відповідає рядок з даними підрозділу регулярному виразу для стрільців
func IsShooter(division string) bool {
	pattern := regexp.MustCompile(`^(1|2|3|4)/(1|2|3|4)/3$`)
	if pattern.MatchString(division) {
		return true
	}
	return false
}
// IsCompanyManager - Перевірка, чи відповідає рядок з даними підрозділу регулярному виразу для управління роти
func IsCompanyManager(division string) bool {
	pattern := regexp.MustCompile(`^упр\ (1|2|3|4)\/3 бо$`)
	if pattern.MatchString(division) {
		return true
	}
	return false
}

// IsVidZab - Перевірка чи відноситься військовослужбовець до відділення забезпечення
func IsVidZab(division string) bool {
	pattern := regexp.MustCompile(`^від\.заб\.\/3 бо$`)
	if pattern.MatchString(division) {
		return true
	}
	return false
}

// IsVidZv - Перевірка чи відноситься військовослужбовець до відділення зв'язку
func IsVidZv(division string) bool {
	pattern := regexp.MustCompile(`^від\.зв\./3 бо$`)
	if pattern.MatchString(division) {
		return true
	}
	return false
}

// IsVidTo - Перевірка чи відноситься військовослужбовець до відділення ехнічного обслуговування
func IsVidTo(division string) bool {
	pattern := regexp.MustCompile(`^від\.то\/3 бо$`)
	if pattern.MatchString(division) {
		return true
	}
	return false
}

// IsMp - Перевірка чи відноситься військовослужбовець до медичного пункту
func IsMp(division string) bool {
	pattern := regexp.MustCompile(`^м.п./3 бо$`)
	if pattern.MatchString(division) {
		return true
	}
	return false
}

// IsManager - Перевірка чи відноситься військовослужбовець управління частиною
func IsManager(division string) bool {
	pattern := regexp.MustCompile(`^упр 3 бо$`)
	if pattern.MatchString(division) {
		return true
	}
	return false
}

// GetPlatoonAndCompany - Визначення номера взводу та роти по типовому запису підрозділу
func GetPlatoonAndCompany(division string) (platoon, company string, err error) {
	shooterRe := regexp.MustCompile(`^(1|2|3|4)/(1|2|3|4)/3$`)
	matches := shooterRe.FindStringSubmatch(division)
	if len(matches) > 2 {
		// matches[0] is the whole match, matches[1] is platoon, matches[2] is company
		return matches[1], matches[2], nil
	} else if len(matches) == 2 {
		return "", "", fmt.Errorf("Не можу отримати номера роти та взводу по запису підрозділу: %s", division)
	}
	return "", "", nil
}

// getCompanyForManagement - Визначення номера роти по запису підрозділу для управління роти
func getCompanyForManagement(division string) (string, error) {
	pattern := regexp.MustCompile(`^упр\ (1|2|3|4)\/3.*$`)
	m := pattern.FindStringSubmatch(division)
	if len(m) >= 2 {
		return m[1], nil
	}
	return "", fmt.Errorf("Не можу отримати номер роти по запису підрозділу: %s", division)
}

// ReadShpkFile - Читання даних з ШПС в структуру даних для персоналу
func ReadShpkFile(shpk_file string) (map[string]Person, error) {
	shpk_data := make(map[string]Person)

	// Відкриття файлу з ШПС в форматі Excel
	shpk_xlsx, err_shpk := excelize.OpenFile(shpk_file)
	if err_shpk != nil {
		log.Printf("Помилка відкриття %s: %v", shpk_file, err_shpk)
		return shpk_data, err_shpk
	}

	// Отрмання таблиці даних ШПС у вигляді рядків
	shpk_rows, err_shpk := shpk_xlsx.GetRows("ШПС")
	if err_shpk != nil {
		log.Printf("Помилка зчитування змісту %s: %v", shpk_file, err_shpk)
	}

	// Заповнення структури даних персоналу змістом, пропускаючи заголовки ШПС
	for i := 2; i < len(shpk_rows) && i < 630; i++ { // index 2 = row 3
		var platoon, company, department string
		var err_platoon, err_company error
		row := shpk_rows[i]

		raw_name := ReadCellSafe(row, 8)
		if raw_name != "" {
			cleaned_name := CleanName(raw_name)
			if cleaned_name != "" {
				department = ReadCellSafe(row, 10)

				switch true {
				case IsShooter(department):
					platoon, company, err_platoon = GetPlatoonAndCompany(department)
					if err_platoon != nil {
						err_shpk = err_platoon
						log.Printf("Помилка отримання номера взводу та роти для %s: %v", cleaned_name, err_platoon)
					}
				case IsCompanyManager(department):
					company, err_company = getCompanyForManagement(department)
					if err_company != nil {
						err_shpk = err_company
						log.Printf("Помилка отримання номеру роти для %s: %v", cleaned_name, err_company)
					}
					platoon = fmt.Sprintf("упр %s/3 бо", company)
				case IsVidZab(department):
					company, platoon = "від.заб./3 бо", ""
				case IsVidZv(department):
					company, platoon = "від.зв./3 бо", ""
				case IsVidTo(department):
					company, platoon = "від.то/3 бо", ""
				case IsMp(department):
					company, platoon = "м.п./3 бо", ""
				case IsManager(department):
					company, platoon = "упр 3 бо", ""
				default:
					company, platoon = "", ""
					log.Printf("Помилка отримання номеру роти для %s", cleaned_name)
				}

				if _, ok := shpk_data[cleaned_name]; ok {
						log.Print("Для цієї персони дані вже збережено:", cleaned_name)
				} else {
					shpk_data[cleaned_name] = Person{
						Department:   department,
						Platoon:      platoon,
						Company:      company,
						Rank:         ReadCellSafe(row, 7),
						Assignment:   ReadCellSafe(row, 20),
						Hospital:     ReadCellSafe(row, 21),
						Vacation_now: ReadCellSafe(row, 23),
						Study:        ReadCellSafe(row, 25),
						Szch:         ReadCellSafe(row, 26),
						Vacation1:    ReadCellSafe(row, 29),
						Telephone:    ReadCellSafe(row, 16),
					}
				}
			}
		}
	}
	return shpk_data, err_shpk
}

// SetRowHeightXlsx - Встановлює висоту рядка в файлі excelize.File, зберігши інші властивості
func SetRowHeightXlsx(f *excelize.File, sheet string, row int, height float64, txt string) error {
	wrap_lines := len(strings.Fields(txt))
	if wrap_lines == 1 {
		return nil
	}
	required_height := (float64(wrap_lines)) * height
	err_height := f.SetRowHeight(sheet, row, required_height)

	return err_height
}

