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
func PrepareReportPPD(shpkDataPtr *map[string]Person) (
	map[string]Distribution, [][]ShortPersData, []string) {

	ppdReportCounter := makeListOfCompanies(PPD_report_list)
	countErr := []string{}

	ppd_list := []ShortPersData{}
	vac_list := []ShortPersData{}
	hosp_list := []ShortPersData{}
	szch_list := []ShortPersData{}
	asmt_list := []ShortPersData{}

	if len(*shpkDataPtr) == 0 {
		countErr = append(countErr, "Потрібні дані ШПК не зчитано з файлу.")
		return ppdReportCounter, [][]ShortPersData{}, countErr
	}

	var aux Distribution
	for name, shpkAttr := range *shpkDataPtr {
		person := ShortPersData{
			Name:       name,
			Department: shpkAttr.Department,
			Rank:       shpkAttr.Rank,
			Company:    shpkAttr.Company,
		}

		// dist - допоміжна змінна для рахунку спискової кількості
		aux = ppdReportCounter[PPD_report_list[5]]
		ppdReportCounter[PPD_report_list[5]] = *(incrementRankCount(&aux, person.Rank))

		switch {
		case shpkAttr.Szch != "" && shpkAttr.Assignment != "":
			errMsg := fmt.Sprintf("Потрібна перевірка актуального статусу для %s: лікування чи відрядження?", name)
			fmt.Println(errMsg)
			countErr = append(countErr, errMsg)

		case shpkAttr.Vacation_now != "" && shpkAttr.Assignment != "":
			errMsg := fmt.Sprintf("Потрібна перевірка актуального статусу для %s: відпустка чи відрядження?", name)
			fmt.Println(errMsg)
			countErr = append(countErr, errMsg)

		case shpkAttr.Hospital != "" && shpkAttr.Assignment != "":
			errMsg := fmt.Sprintf("Потрібна перевірка актуального статусу для %s: відпустка чи відрядження?", name)
			fmt.Println(errMsg)
			countErr = append(countErr, errMsg)

		case shpkAttr.Assignment == "ППД":
			aux = ppdReportCounter[PPD_report_list[0]]
			ppdReportCounter[PPD_report_list[0]] = *(categorizePPD(person, &ppd_list, &aux))

		case  shpkAttr.Assignment != "":
			aux = ppdReportCounter[PPD_report_list[4]]
			ppdReportCounter[PPD_report_list[4]] = *(categorizePPD(person, &asmt_list, &aux))

		case shpkAttr.Vacation_now != "" && shpkAttr.Assignment == "":
			aux = ppdReportCounter[PPD_report_list[1]]
			ppdReportCounter[PPD_report_list[1]] = *(categorizePPD(person, &vac_list, &aux))

		case shpkAttr.Hospital != "" && shpkAttr.Assignment == "":
			aux = ppdReportCounter[PPD_report_list[2]]
			ppdReportCounter[PPD_report_list[2]] = *(categorizePPD(person, &hosp_list, &aux))

		case shpkAttr.Szch != "" && shpkAttr.Assignment == "":
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
	return ppdReportCounter, reportList, countErr
}

// PrepareReportBO - Підготовка розгорнутого звіту по всьому підрозділу
func PrepareReportBO(shpkDataPtr *map[string]Person) (
	map[string]map[string]Distribution, []string) {

	countErr := []string{}
	boReportCounter := make(map[string]map[string]Distribution, len(COMP_list))

	// Заповнюємо boReportCounter нулями
	for _, c := range COMP_list {
		boReportCounter[c] = makeListOfCompanies(BO_report_list)
	}

	for name, shpkAttr := range *shpkDataPtr {
		switch true {
		case name == "" || name == " ":
			errMsg := fmt.Sprintf(
				"Відсутні дані стосовно імені для особи, що має звання %s в підрозділі %s",
				shpkAttr.Rank, shpkAttr.Department)
			countErr = append(countErr, errMsg)
		case shpkAttr.Rank == ""|| shpkAttr.Rank == " ":
			errMsg := fmt.Sprintf("Для %s відсутні дані стосовно звання", name)
			countErr = append(countErr, errMsg)
		case shpkAttr.Company == ""|| shpkAttr.Company == " ":
			errMsg := fmt.Sprintf("Для %s відсутні дані стосовно підрозділу", name)
			countErr = append(countErr, errMsg)
		}

		dist := boReportCounter[shpkAttr.Company][BO_report_list[0]]
		incrementRankCount(&dist, shpkAttr.Rank)
		boReportCounter[shpkAttr.Company][BO_report_list[0]] = dist

		switch true {
		case shpkAttr.Szch != "" && (shpkAttr.Assignment != "" ||
			shpkAttr.Hospital != "" || shpkAttr.Vacation_now != "" || shpkAttr.Study != ""):
			errMsg := fmt.Sprintf("Для %s одночасно є дані про СЗЧ і про наявність!", name)
			countErr = append(countErr, errMsg)

		case shpkAttr.Vacation_now != "" && shpkAttr.Assignment != "":
			errMsg := fmt.Sprintf("Для %s одночасно є дані про відрядження і про відпустку!", name)
			countErr = append(countErr, errMsg)

		case shpkAttr.Hospital != "" && (shpkAttr.Assignment != "" || shpkAttr.Vacation_now != "" || shpkAttr.Study != ""):
			errMsg := fmt.Sprintf("Для %s одночасно є дані про лікування і про наявність!", name)
			countErr = append(countErr, errMsg)

		case shpkAttr.Assignment == "ППД":
			dist := boReportCounter[shpkAttr.Company][BO_report_list[8]]
			incrementRankCount(&dist, shpkAttr.Rank)
			boReportCounter[shpkAttr.Company][BO_report_list[8]] = dist

		case shpkAttr.Szch != "":
			dist := boReportCounter[shpkAttr.Company][BO_report_list[7]]
			incrementRankCount(&dist, shpkAttr.Rank)
			boReportCounter[shpkAttr.Company][BO_report_list[7]] = dist

		case shpkAttr.Assignment == "ВОП":
			dist := boReportCounter[shpkAttr.Company][BO_report_list[6]]
			incrementRankCount(&dist, shpkAttr.Rank)
			boReportCounter[shpkAttr.Company][BO_report_list[6]] = dist

		case shpkAttr.Assignment == "КСП":
			dist := boReportCounter[shpkAttr.Company][BO_report_list[5]]
			incrementRankCount(&dist, shpkAttr.Rank)
			boReportCounter[shpkAttr.Company][BO_report_list[5]] = dist

		case shpkAttr.Assignment != "":
			dist := boReportCounter[shpkAttr.Company][BO_report_list[4]]
			incrementRankCount(&dist, shpkAttr.Rank)
			boReportCounter[shpkAttr.Company][BO_report_list[4]] = dist

		case shpkAttr.Study != "":
			dist := boReportCounter[shpkAttr.Company][BO_report_list[3]]
			incrementRankCount(&dist, shpkAttr.Rank)
			boReportCounter[shpkAttr.Company][BO_report_list[3]] = dist

		case shpkAttr.Hospital != "":
			dist := boReportCounter[shpkAttr.Company][BO_report_list[2]]
			incrementRankCount(&dist, shpkAttr.Rank)
			boReportCounter[shpkAttr.Company][BO_report_list[2]] = dist

		case shpkAttr.Vacation_now != "":
			dist := boReportCounter[shpkAttr.Company][BO_report_list[1]]
			incrementRankCount(&dist, shpkAttr.Rank)
			boReportCounter[shpkAttr.Company][BO_report_list[1]] = dist
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

	return boReportCounter, countErr
}

// PrepareVacationReport1 - Підготовка даних для звіту стосовно відгуляних перших частин щорічної відпустки
func PrepareVacationReport1(shpkDataPtr *map[string]Person) (
	*[][]string, []string) {

	VacReport1 := [][]string{}
	countErr := []string{}
	// totalCounter - Для визначення спискової кількості людей
	totalCounter := makeListOfCompanies(COMP_list)
	// vac1ReportCounter - для звіту по відгуляним відпусткам 1 черги
	vac1ReportCounter := makeListOfCompanies(COMP_list)

	var aux Distribution
	for name, shpkAttr := range *shpkDataPtr {
		person := ShortPersData{
			Name:       name,
			Department: shpkAttr.Department,
			Rank:       shpkAttr.Rank,
			Company:    shpkAttr.Company,
		}

		// Рахунок  загальної спискової кількості
		aux = totalCounter[COMP_list[9]]
		totalCounter[COMP_list[9]] = *(incrementRankCount(&aux, person.Rank))

		// Рахунок спискової кількості по підрозділам
		aux = totalCounter[person.Company]
		totalCounter[person.Company] = *(incrementRankCount(&aux, person.Rank))

		// Рахунок тих, что відгуляв першу частину відпустки
		if shpkAttr.Vacation1 != "" {
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

	return &VacReport1, countErr
}