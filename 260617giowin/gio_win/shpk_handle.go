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

type Distribution struct {
	Offi int
	Serg int
	Sold int
	Total int
}

type ShortPersData struct {
	Name string
	Department string
	Rank string
}

// list_of_companies - Впорядкований список підрозділів
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
	"підсумок",
}

var list_for_ppd_report []string = []string {
	"ППД",
	"Відпустка",
	"Шпиталь",
	"СЗЧ",
	"Відрядження",
	"Загалом",
}


// MakeListOfCompanies - Створення списку підрозділів з нульовими даними розподілу
func MakeListOfCompanies(list []string) map[string]Distribution {
	companyDist := make(map[string]Distribution, len(list))
	for _, name := range list {
			companyDist[name] = Distribution{}
	}
	return companyDist
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

// CreateReportPPD - Створення скороченого звіту для ППД
func CreateReportPPD(shpk_data map[string]Person) /*map[string]Distribution*/ (map[string]Distribution, [][]ShortPersData, []string) {

	ppdReportCounter := make(map[string]Distribution, len(list_for_ppd_report))
	ppd_count := ppdReportCounter[list_for_ppd_report[0]]
	vac_count := ppdReportCounter[list_for_ppd_report[1]]
	hosp_count := ppdReportCounter[list_for_ppd_report[2]]
	szch_count := ppdReportCounter[list_for_ppd_report[3]]
	asmt_count := ppdReportCounter[list_for_ppd_report[4]]
	total_count := ppdReportCounter[list_for_ppd_report[5]]
	count_err := []string{}

	ppd_list := []ShortPersData{}
	vac_list := []ShortPersData{}
	hosp_list := []ShortPersData{}
	szch_list := []ShortPersData{}
	asmt_list := []ShortPersData{}

	for name, shpk_attr := range shpk_data {
		rank := shpk_attr.Rank
		department := shpk_attr.Department

		if strings.HasSuffix(rank, "олдат") {
			total_count.Sold++
		} else if strings.HasSuffix(rank, "ержант") {
			total_count.Serg++
		} else {
			total_count.Offi++
		}
		total_count.Total++


		if shpk_attr.Assignment == "ППД" {
			ppd_list = append(ppd_list, ShortPersData{
				Name:        name,
				Department:  department,
				Rank:        rank,
			})
			if strings.HasSuffix(rank, "олдат") {
				ppd_count.Sold++
			} else if strings.HasSuffix(rank, "ержант") {
				ppd_count.Serg++
			} else {
				ppd_count.Offi++
			}
			ppd_count.Total++
			} else if shpk_attr.Assignment != "" {
				asmt_list = append(asmt_list, ShortPersData{
					Name:        name,
					Department:  department,
					Rank:        rank,
				})
				if strings.HasSuffix(rank, "олдат") {
					asmt_count.Sold++
				} else if strings.HasSuffix(rank, "ержант") {
					asmt_count.Serg++
				} else {
					asmt_count.Offi++
				}
				asmt_count.Total++
			}

			if shpk_attr.Vacation_now != "" && shpk_attr.Assignment == ""{
				vac_list = append(vac_list, ShortPersData{
					Name:        name,
					Department:  department,
					Rank:        rank,
				})
				if strings.HasSuffix(rank, "олдат") {
					vac_count.Sold++
				} else if strings.HasSuffix(rank, "ержант") {
					vac_count.Serg++
				} else {
					vac_count.Offi++
				}
				vac_count.Total++
		} else if shpk_attr.Vacation_now != "" && shpk_attr.Assignment != ""{
			err_msg := fmt.Sprintf("Потрібна перевірка актуального статусу для %s: відпустка чи відрядження?", name)
			fmt.Println(err_msg)
			count_err = append(count_err, "\n", err_msg)
		}

		if shpk_attr.Hospital != "" && shpk_attr.Assignment == ""{
			hosp_list = append(hosp_list, ShortPersData{
				Name:        name,
				Department:  department,
				Rank:        rank,
			})
			if strings.HasSuffix(rank, "олдат") {
				hosp_count.Sold++
			} else if strings.HasSuffix(rank, "ержант") {
				hosp_count.Serg++
			} else {
				hosp_count.Offi++
			}
			hosp_count.Total++
		} else if shpk_attr.Hospital != "" && shpk_attr.Assignment != ""{
			err_msg := fmt.Sprintf("Потрібна перевірка актуального статусу для %s: відпустка чи відрядження?", name)
			fmt.Println(err_msg)
			count_err = append(count_err, "\n", err_msg)
		}

		if shpk_attr.Szch != "" && shpk_attr.Assignment == ""{
			szch_list = append(szch_list, ShortPersData{
				Name:        name,
				Department:  department,
				Rank:        rank,
			})
			if strings.HasSuffix(rank, "олдат") {
				szch_count.Sold++
			} else if strings.HasSuffix(rank, "ержант") {
				szch_count.Serg++
			} else {
				szch_count.Offi++
			}
			szch_count.Total++
		} else if shpk_attr.Szch != "" && shpk_attr.Assignment != ""{
			err_msg := fmt.Sprintf("Потрібна перевірка актуального статусу для %s: лікування чи відрядження?", name)
			fmt.Println(err_msg)
			count_err = append(count_err, "\n", err_msg)
		}
	}

	reportList := [][]ShortPersData{ppd_list, vac_list, hosp_list, szch_list, asmt_list}

	ppdReportCounter[list_for_ppd_report[0]] = ppd_count
	ppdReportCounter[list_for_ppd_report[1]] = vac_count
	ppdReportCounter[list_for_ppd_report[2]] = hosp_count
	ppdReportCounter[list_for_ppd_report[3]] = szch_count
	ppdReportCounter[list_for_ppd_report[4]] = asmt_count
	ppdReportCounter[list_for_ppd_report[5]] = total_count

	return ppdReportCounter, reportList, count_err
}


// CreateReportBO - Створення розгорнутого звіту по всьому підрозділу
func CreateReportBO(shpk_data map[string]Person) map[string]Distribution {
	// compDistr := MakeListOfCompanies(list_of_companies)
	// manager_dist := compDistr["упр 3 бо"]
	// c1_dist := compDistr["1"]
	// c2_dist := compDistr["2"]
	// c3_dist := compDistr["3"]
	// c4_dist := compDistr["4"]
	// zv_dist := compDistr["від.зв./3 бо"]
	// zab_dist := compDistr["від.заб./3 бо"]
	// to_dist := compDistr["від.то/3 бо"]
	// mp_dist := compDistr["м.п./3 бо"]
	// total_dist := compDistr["підсумок"]
	boReportCounter := make(map[string]Distribution, len(list_of_companies))
	return boReportCounter
}