package gio_win

import (
	"fmt"
	"strings"
	"time"
)

// makeListOfCompanies - Створення списку підрозділів з нульовими даними розподілу int(0)
func makeListOfCompanies(list []string) map[string]Distribution {
	companyDist := make(map[string]Distribution, len(list))
	for _, name := range list {
		companyDist[name] = Distribution{}
	}
	return companyDist
}

// incrementRankCount -інкрементує відповідні лічильники в структурі Disribution
func incrementRankCount(dist *Distribution, rank string) *Distribution {
	getRankCategory := ""
	switch {
	case strings.HasSuffix(rank, "олдат"):
		getRankCategory = "Sold"
	case strings.HasSuffix(rank, "ержант"):
		getRankCategory = "Serg"
	default:
		getRankCategory = "Offi"
	}
	switch getRankCategory {
	case "Sold":
		dist.Sold++
	case "Serg":
		dist.Serg++
	case "Offi":
		dist.Offi++
	}
	dist.Total++

	return dist
}

// categorizePPD - додає персону до відповідного списку і збільшує лічильники
func categorizePPD(
	person ShortPersData,
	list *[]ShortPersData,
	counter *Distribution,
	) *Distribution {
	*list = append(*list, person)
	return incrementRankCount(counter, person.Rank)
}

// PrepareReportPPD - Підготовка скороченого звіту для ППД
func PrepareReportPPD(shpk_data map[string]Person) (
	map[string]Distribution, [][]ShortPersData, []string) {

	ppdReportCounter := makeListOfCompanies(PPD_report_list)
	count_err := []string{}

	ppd_list := []ShortPersData{}
	vac_list := []ShortPersData{}
	hosp_list := []ShortPersData{}
	szch_list := []ShortPersData{}
	asmt_list := []ShortPersData{}

	if len(shpk_data) == 0 {
		count_err = append(count_err, "Потрібні дані ШПК не зчитано з файлу.")
		return ppdReportCounter, [][]ShortPersData{}, count_err
	}

	var aux Distribution
	for name, shpk_attr := range shpk_data {
		person := ShortPersData{
			Name:       name,
			Department: shpk_attr.Department,
			Rank:       shpk_attr.Rank,
			Company:    shpk_attr.Company,
		}

		// dist - допоміжна змінна для рахунку спискової кількості
		aux = ppdReportCounter[PPD_report_list[5]]
		ppdReportCounter[PPD_report_list[5]] = *(incrementRankCount(&aux, person.Rank))

		switch {
		case shpk_attr.Szch != "" && shpk_attr.Assignment != "":
			err_msg := fmt.Sprintf("Потрібна перевірка актуального статусу для %s: лікування чи відрядження?", name)
			fmt.Println(err_msg)
			count_err = append(count_err, err_msg)

		case shpk_attr.Vacation_now != "" && shpk_attr.Assignment != "":
			err_msg := fmt.Sprintf("Потрібна перевірка актуального статусу для %s: відпустка чи відрядження?", name)
			fmt.Println(err_msg)
			count_err = append(count_err, err_msg)

		case shpk_attr.Hospital != "" && shpk_attr.Assignment != "":
			err_msg := fmt.Sprintf("Потрібна перевірка актуального статусу для %s: відпустка чи відрядження?", name)
			fmt.Println(err_msg)
			count_err = append(count_err, err_msg)

		case shpk_attr.Assignment == "ППД":
			aux = ppdReportCounter[PPD_report_list[0]]
			ppdReportCounter[PPD_report_list[0]] = *(categorizePPD(person, &ppd_list, &aux))

		case  shpk_attr.Assignment != "":
			aux = ppdReportCounter[PPD_report_list[4]]
			ppdReportCounter[PPD_report_list[4]] = *(categorizePPD(person, &asmt_list, &aux))

		case shpk_attr.Vacation_now != "" && shpk_attr.Assignment == "":
			aux = ppdReportCounter[PPD_report_list[1]]
			ppdReportCounter[PPD_report_list[1]] = *(categorizePPD(person, &vac_list, &aux))

		case shpk_attr.Hospital != "" && shpk_attr.Assignment == "":
			aux = ppdReportCounter[PPD_report_list[2]]
			ppdReportCounter[PPD_report_list[2]] = *(categorizePPD(person, &hosp_list, &aux))

		case shpk_attr.Szch != "" && shpk_attr.Assignment == "":
			aux = ppdReportCounter[PPD_report_list[3]]
			ppdReportCounter[PPD_report_list[3]] = *(categorizePPD(person, &szch_list, &aux))
		}
	}

	reportList := [][]ShortPersData{
		ppd_list,
		vac_list,
		hosp_list,
		szch_list,
		asmt_list}
	return ppdReportCounter, reportList, count_err
}

// PrepareReportBO - Підготовка розгорнутого звіту по всьому підрозділу
func PrepareReportBO(shpk_data map[string]Person) (
	map[string]map[string]Distribution, []string) {

	count_err := []string{}
	boReportCounter := make(map[string]map[string]Distribution, len(COMP_list))

	// Заповнюємо boReportCounter нулями
	for _, c := range COMP_list {
		boReportCounter[c] = makeListOfCompanies(BO_report_list)
	}

	for name, shpk_attr := range shpk_data {
		switch true {
		case name == "" || name == " ":
			err_msg := fmt.Sprintf(
				"Відсутні дані стосовно імені для особи, що має звання %s в підрозділі %s",
				shpk_attr.Rank, shpk_attr.Department)
			count_err = append(count_err, err_msg)
		case shpk_attr.Rank == ""|| shpk_attr.Rank == " ":
			err_msg := fmt.Sprintf("Для %s відсутні дані стосовно звання", name)
			count_err = append(count_err, err_msg)
		case shpk_attr.Company == ""|| shpk_attr.Company == " ":
			err_msg := fmt.Sprintf("Для %s відсутні дані стосовно підрозділу", name)
			count_err = append(count_err, err_msg)
		}

		dist := boReportCounter[shpk_attr.Company][BO_report_list[0]]
		incrementRankCount(&dist, shpk_attr.Rank)
		boReportCounter[shpk_attr.Company][BO_report_list[0]] = dist

		switch true {
		case shpk_attr.Szch != "" && (shpk_attr.Assignment != "" ||
			shpk_attr.Hospital != "" || shpk_attr.Vacation_now != "" || shpk_attr.Study != ""):
			err_msg := fmt.Sprintf("Для %s одночасно є дані про СЗЧ і про наявність!", name)
			count_err = append(count_err, err_msg)

		case shpk_attr.Vacation_now != "" && shpk_attr.Assignment != "":
			err_msg := fmt.Sprintf("Для %s одночасно є дані про відрядження і про відпустку!", name)
			count_err = append(count_err, err_msg)

		case shpk_attr.Hospital != "" && (shpk_attr.Assignment != "" || shpk_attr.Vacation_now != "" || shpk_attr.Study != ""):
			err_msg := fmt.Sprintf("Для %s одночасно є дані про лікування і про наявність!", name)
			count_err = append(count_err, err_msg)

		case shpk_attr.Assignment == "ППД":
			dist := boReportCounter[shpk_attr.Company][BO_report_list[8]]
			incrementRankCount(&dist, shpk_attr.Rank)
			boReportCounter[shpk_attr.Company][BO_report_list[8]] = dist

		case shpk_attr.Szch != "":
			dist := boReportCounter[shpk_attr.Company][BO_report_list[7]]
			incrementRankCount(&dist, shpk_attr.Rank)
			boReportCounter[shpk_attr.Company][BO_report_list[7]] = dist

		case shpk_attr.Assignment == "ВОП":
			dist := boReportCounter[shpk_attr.Company][BO_report_list[6]]
			incrementRankCount(&dist, shpk_attr.Rank)
			boReportCounter[shpk_attr.Company][BO_report_list[6]] = dist

		case shpk_attr.Assignment == "КСП":
			dist := boReportCounter[shpk_attr.Company][BO_report_list[5]]
			incrementRankCount(&dist, shpk_attr.Rank)
			boReportCounter[shpk_attr.Company][BO_report_list[5]] = dist

		case shpk_attr.Assignment != "":
			dist := boReportCounter[shpk_attr.Company][BO_report_list[4]]
			incrementRankCount(&dist, shpk_attr.Rank)
			boReportCounter[shpk_attr.Company][BO_report_list[4]] = dist

		case shpk_attr.Study != "":
			dist := boReportCounter[shpk_attr.Company][BO_report_list[3]]
			incrementRankCount(&dist, shpk_attr.Rank)
			boReportCounter[shpk_attr.Company][BO_report_list[3]] = dist

		case shpk_attr.Hospital != "":
			dist := boReportCounter[shpk_attr.Company][BO_report_list[2]]
			incrementRankCount(&dist, shpk_attr.Rank)
			boReportCounter[shpk_attr.Company][BO_report_list[2]] = dist

		case shpk_attr.Vacation_now != "":
			dist := boReportCounter[shpk_attr.Company][BO_report_list[1]]
			incrementRankCount(&dist, shpk_attr.Rank)
			boReportCounter[shpk_attr.Company][BO_report_list[1]] = dist
		}
	}

	// Визначення підсумкових даних за призначеннями по званням
	for _, brl := range BO_report_list {
		offi, serg, sold := 0, 0, 0
		for _, comp := range COMP_list {
		if comp == "підсумок" { continue }
			el := boReportCounter[comp][brl]
			offi += el.Offi
			serg += el.Serg
			sold += el.Sold
		}
		boReportCounter["підсумок"][brl] = Distribution{
			Offi : offi,
			Serg : serg,
			Sold : sold,
		}
	}

	return boReportCounter, count_err
}

// PrepareVacationReport1 - Підготовка даних для звіту стосовно відгуляних перших частин щорічної відпустки
func PrepareVacationReport1(shpk_data map[string]Person) (
	VacReport1 [][]string, count_err []string) {

	// totalCounter - Для визначення спискової кількості людей
	totalCounter := makeListOfCompanies(COMP_list)
	// vac1ReportCounter - для звіту по відгуляним відпусткам 1 черги
	vac1ReportCounter := makeListOfCompanies(COMP_list)

	var aux Distribution
	for name, shpk_attr := range shpk_data {
		person := ShortPersData{
			Name:       name,
			Department: shpk_attr.Department,
			Rank:       shpk_attr.Rank,
			Company:    shpk_attr.Company,
		}

		// Рахунок  загальної спискової кількості
		aux = totalCounter[COMP_list[9]]
		totalCounter[COMP_list[9]] = *(incrementRankCount(&aux, person.Rank))

		// Рахунок спискової кількості по підрозділам
		aux = totalCounter[person.Company]
		totalCounter[person.Company] = *(incrementRankCount(&aux, person.Rank))

		// Рахунок тих, что відгуляв першу частину відпустки
		if shpk_attr.Vacation1 != "" {
			aux = vac1ReportCounter[person.Company]
			vac1ReportCounter[person.Company] = *(incrementRankCount(&aux, person.Rank))
		}
	}

	now := time.Now()
	dateTime := now.Format("02.01.2006")
	firstLineText := fmt.Sprintf("Кількість особового складу 3бо, що відгуляла першу частину щорічної відпустки станом на %v", dateTime)
	VacReport1 = append(VacReport1, []string{firstLineText, "", "", ""})
	VacReport1 = append(VacReport1, []string{"Підрозділ", "За списком", "Відгуляли", "Процент"})

	for _, cl := range COMP_list{
		fmt.Println(cl)
		t := totalCounter[cl]
		v := vac1ReportCounter[cl]
		colB := fmt.Sprintf("%d (%d-%d-%d)", t.Total, t.Offi, t.Serg, t.Sold)
		colC := fmt.Sprintf("%d (%d-%d-%d)", v.Total, v.Offi, v.Serg, v.Sold)
		colD := "100.00 "
		if t.Total > 0 {
			percent := 100.0 * float32(v.Total) / float32(t.Total)
			colD = fmt.Sprintf("%.2f ", percent)
		}
		VacReport1 = append(VacReport1, []string{cl, colB, colC, colD})
	}

	return VacReport1, count_err
}