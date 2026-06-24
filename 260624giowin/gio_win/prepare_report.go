package gio_win

import (
	"fmt"
	"strings"
)


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
	Company string
}

// MakeListOfCompanies - Створення списку підрозділів з нульовими даними розподілу
func MakeListOfCompanies(list []string) map[string]Distribution {
	companyDist := make(map[string]Distribution, len(list))
	for _, name := range list {
			companyDist[name] = Distribution{}
	}
	return companyDist
}

// ppd_report_list - Перелік категорій, за якими ведеться розподіл особового складу для звіту ППД
var ppd_report_list []string = []string {
	"ППД",
	"Відпустка",
	"Шпиталь",
	"СЗЧ",
	"Відрядження",
	"Загалом",
}

// bo_report_list - Перелік категорій, за якими ведеться розподіл особового складу по підрозділам всієї частини
var bo_report_list []string = []string {
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

// incrementRankCount - increments the appropriate rank counter in Distribution
func incrementRankCount(dist *Distribution, rank string) {
	getRankCategory := ""
	switch true {
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
}

// categorizePPD - adds person to appropriate list and updates counter
func categorizePPD(
	person ShortPersData,
	list *[]ShortPersData,
	counters map[string]Distribution,
	key string,
	) {
    *list = append(*list, person)
    dist := counters[key]
    incrementRankCount(&dist, person.Rank)
    counters[key] = dist
}

// PrepareReportPPD - Підготовка скороченого звіту для ППД
func PrepareReportPPD(
	shpk_data map[string]Person,
	) (
		map[string]Distribution,
		[][]ShortPersData,
		[]string,
		) {
	ppdReportCounter := make(map[string]Distribution, len(ppd_report_list))
	count_err := []string{}

	ppd_list := []ShortPersData{}
	vac_list := []ShortPersData{}
	hosp_list := []ShortPersData{}
	szch_list := []ShortPersData{}
	asmt_list := []ShortPersData{}

	for _, key := range ppd_report_list {
			ppdReportCounter[key] = Distribution{}
	}

	for name, shpk_attr := range shpk_data {
		person := ShortPersData{
				Name:       name,
				Department: shpk_attr.Department,
				Rank:       shpk_attr.Rank,
				Company:    shpk_attr.Company,
		}

		dist := ppdReportCounter[ppd_report_list[5]]
		incrementRankCount(&dist, person.Rank)
		ppdReportCounter[ppd_report_list[5]] = dist

		if shpk_attr.Assignment == "ППД" {
				categorizePPD(person, &ppd_list, ppdReportCounter,
					ppd_report_list[0])
		} else if shpk_attr.Assignment != "" {
				categorizePPD(person, &asmt_list, ppdReportCounter,
					ppd_report_list[4])
		}

		if shpk_attr.Vacation_now != "" && shpk_attr.Assignment == "" {
				categorizePPD(person, &vac_list, ppdReportCounter,
					ppd_report_list[1])
		} else if shpk_attr.Vacation_now != "" && shpk_attr.Assignment != "" {
				err_msg := fmt.Sprintf("Потрібна перевірка актуального статусу для %s: відпустка чи відрядження?", name)
				fmt.Println(err_msg)
				count_err = append(count_err, err_msg)
		}

		if shpk_attr.Hospital != "" && shpk_attr.Assignment == "" {
				categorizePPD(person, &hosp_list, ppdReportCounter,
					ppd_report_list[2])
		} else if shpk_attr.Hospital != "" && shpk_attr.Assignment != "" {
				err_msg := fmt.Sprintf("Потрібна перевірка актуального статусу для %s: відпустка чи відрядження?", name)
				fmt.Println(err_msg)
				count_err = append(count_err, err_msg)
		}

		if shpk_attr.Szch != "" && shpk_attr.Assignment == "" {
				categorizePPD(person, &szch_list, ppdReportCounter,
					ppd_report_list[3])
		} else if shpk_attr.Szch != "" && shpk_attr.Assignment != "" {
				err_msg := fmt.Sprintf("Потрібна перевірка актуального статусу для %s: лікування чи відрядження?", name)
				fmt.Println(err_msg)
				count_err = append(count_err, err_msg)
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
	map[string]map[string]Distribution, []string	) {

	count_err := []string{}
	boReportCounter := make(map[string]map[string]Distribution, len(comp_list))

	for _, c := range comp_list {
		boReportCounter[c] = make(map[string]Distribution, len(bo_report_list))
		for _, d := range bo_report_list {
			boReportCounter[c][d] = Distribution{}
			// fmt.Println(boReportCounter[c][d])
		}
	}

	for name, shpk_attr := range shpk_data {
		switch true {
		case name == "" || name == " ":
			err_msg := fmt.Sprintf(
				"Відсутні дані стосовно імені для особи, що має звання %s в підрозділі %s",
				shpk_attr.Rank, shpk_attr.Department)
			count_err = append(count_err, err_msg)
		case shpk_attr.Rank == "":
			err_msg := fmt.Sprintf("Для %s відсутні дані стосовно звання", name)
			count_err = append(count_err, err_msg)
		case shpk_attr.Company == "":
			err_msg := fmt.Sprintf("Для %s відсутні дані стосовно підрозділу", name)
			count_err = append(count_err, err_msg)
		}

		dist := boReportCounter[shpk_attr.Company][bo_report_list[0]]
		incrementRankCount(&dist, shpk_attr.Rank)
		boReportCounter[shpk_attr.Company][bo_report_list[0]] = dist

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
			dist := boReportCounter[shpk_attr.Company][bo_report_list[8]]
			incrementRankCount(&dist, shpk_attr.Rank)
			boReportCounter[shpk_attr.Company][bo_report_list[8]] = dist

		case shpk_attr.Szch != "":
				dist := boReportCounter[shpk_attr.Company][bo_report_list[7]]
				incrementRankCount(&dist, shpk_attr.Rank)
				boReportCounter[shpk_attr.Company][bo_report_list[7]] = dist

		case shpk_attr.Assignment == "ВОП":
			dist := boReportCounter[shpk_attr.Company][bo_report_list[6]]
			incrementRankCount(&dist, shpk_attr.Rank)
			boReportCounter[shpk_attr.Company][bo_report_list[6]] = dist

		case shpk_attr.Assignment == "КСП":
			dist := boReportCounter[shpk_attr.Company][bo_report_list[5]]
			incrementRankCount(&dist, shpk_attr.Rank)
			boReportCounter[shpk_attr.Company][bo_report_list[5]] = dist

		case shpk_attr.Assignment != "":
			dist := boReportCounter[shpk_attr.Company][bo_report_list[4]]
			incrementRankCount(&dist, shpk_attr.Rank)
			boReportCounter[shpk_attr.Company][bo_report_list[4]] = dist

		case shpk_attr.Study != "":
			dist := boReportCounter[shpk_attr.Company][bo_report_list[3]]
			incrementRankCount(&dist, shpk_attr.Rank)
			boReportCounter[shpk_attr.Company][bo_report_list[3]] = dist

		case shpk_attr.Hospital != "":
			dist := boReportCounter[shpk_attr.Company][bo_report_list[2]]
			incrementRankCount(&dist, shpk_attr.Rank)
			boReportCounter[shpk_attr.Company][bo_report_list[2]] = dist

		case shpk_attr.Vacation_now != "":
			dist := boReportCounter[shpk_attr.Company][bo_report_list[1]]
			incrementRankCount(&dist, shpk_attr.Rank)
			boReportCounter[shpk_attr.Company][bo_report_list[1]] = dist
		}
	}

	return boReportCounter, count_err
}