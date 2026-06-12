package main

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

type Distribution struct {
	Offi int
	Serg int
	Sold int
	Total int
}

var list_of_companies []string = []string{
	"упр 3 бо",
	"1",
	"2",
	"3",
	"4",
	"від.зв./3 бо",
	"від.заб./3 бо",
	"від.то/3 бо",
	"м.п./3 бо",
	}

func MakeListOfCompanies() map[string]Distribution {
	companyDist := make(map[string]Distribution, len(list_of_companies))
	for _, name := range list_of_companies {
			companyDist[name] = Distribution{}
	}
	return companyDist
}

func CleanName(name string) string {
	// Очистка імени від зайвих символів
	if name == "" {
		return ""
	}
	return strings.TrimSpace(strings.ReplaceAll(name, "\n", " "))
}

func IsShooter(division string) bool {
	// Перевірка, чи відповідає рядок з даними підрозділу регулярному виразу для стрільців
	pattern := regexp.MustCompile(`^(1|2|3|4)/(1|2|3|4)/3$`)
	if pattern.MatchString(division) {
		return true
	}
	return false
}
func IsCompanyManager(division string) bool {
	// Перевірка, чи відповідає рядок з даними підрозділу регулярному виразу для управління роти
	pattern := regexp.MustCompile(`^упр\ (1|2|3|4)\/3 бо$`)
	if pattern.MatchString(division) {
		return true
	}
	return false
}

func IsVidZab(division string) bool {
	// Перевірка чи відноситься військовослужбовець до відділення забезпечення
	pattern := regexp.MustCompile(`^від\.заб\.\/3 бо$`)
	if pattern.MatchString(division) {
		return true
	}
	return false
}

func IsVidZv(division string) bool {
	// Перевірка чи відноситься військовослужбовець до відділення зв'язку
	pattern := regexp.MustCompile(`^від\.зв\./3 бо$`)
	if pattern.MatchString(division) {
		return true
	}
	return false
}

func IsVidTo(division string) bool {
	// Перевірка чи відноситься військовослужбовець до відділення ехнічного обслуговування
	pattern := regexp.MustCompile(`^від\.то\/3 бо$`)
	if pattern.MatchString(division) {
		return true
	}
	return false
}

func IsMp(division string) bool {
	// Перевірка чи відноситься військовослужбовець до медичного пункту
	pattern := regexp.MustCompile(`^м.п./3 бо$`)
	if pattern.MatchString(division) {
		return true
	}
	return false
}

func IsManager(division string) bool {
	// Перевірка чи відноситься військовослужбовець управління частиною
	pattern := regexp.MustCompile(`^упр 3 бо$`)
	if pattern.MatchString(division) {
		return true
	}
	return false
}

func GetPlatoonAndCompany(division string) (platoon, company string, err error) {
	// Визначення номера взводу та роти по типовому запису підрозділу
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

func getCompanyForManagement(division string) (string, error) {
	pattern := regexp.MustCompile(`^упр\ (1|2|3|4)\/3.*$`)
	m := pattern.FindStringSubmatch(division)
	if len(m) >= 2 {
		return m[1], nil
	}
	return "", fmt.Errorf("Не можу отримати номер роти по запису підрозділу: %s", division)
}

func ReadShpkFile(shpk_file string) (map[string]Person, error) {
	// Структура даних для персоналу
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

		if len(row) > 16 && row[8] != "" {
			cleaned_name := CleanName(row[8])
			if cleaned_name != "" {
				department = row[10]

				if IsShooter(department) {
					platoon, company, err_platoon = GetPlatoonAndCompany(department)
					if err_platoon != nil {
						err_shpk = err_platoon
						log.Printf("Помилка отримання номера взводу та роти для %s: %v", cleaned_name, err_platoon)
					}
				} else if IsCompanyManager(department) {
					company, err_company = getCompanyForManagement(department)
					if err_company != nil {
						err_shpk = err_company
						log.Printf("Помилка отримання номеру роти для %s: %v", cleaned_name, err_company)
					}
					platoon = fmt.Sprintf("упр %s/3 бо", company)
				} else if IsVidZab(department) {
					company, platoon = "від.заб./3 бо", ""
				} else if IsVidZv(department) {
					company, platoon = "від.зв./3 бо", ""
				} else if IsVidTo(department) {
					company, platoon = "від.то/3 бо", ""
				} else if IsMp(department) {
					company, platoon = "м.п./3 бо", ""
				} else if IsManager(department) {
					company, platoon = "упр 3 бо", ""
				} else {
					company, platoon = "", ""
					log.Printf("Помилка отримання номеру роти для %s", cleaned_name)
				}

				shpk_data[cleaned_name] = Person{
					Department:   department,
					Platoon:      platoon,
					Company:      company,
					Rank:         row[7],
					Assignment:   row[20],
					Hospital:     row[21],
					Vacation_now: row[23],
					Study:        row[25],
					Szch:         row[26],
					Vacation1:    row[29],
					Telephone:    row[16],
				}
			}
		}
	}
	return shpk_data, err_shpk
}

func SetRowHeightXlsx(f *excelize.File, sheet string, row int, height float64, txt string) error {
	// SetRowHeightXlsx встановлює висоту рядка, зберігши інші властивості
	wrap_lines := len(strings.Fields(txt))
	if wrap_lines == 1 {
		return nil
	}
	required_height := (float64(wrap_lines)) * height
	err_height := f.SetRowHeight(sheet, row, required_height)

	return err_height
}
