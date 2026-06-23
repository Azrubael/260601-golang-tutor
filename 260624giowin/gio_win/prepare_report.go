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
}

// MakeListOfCompanies - Створення списку підрозділів з нульовими даними розподілу
func MakeListOfCompanies(list []string) map[string]Distribution {
	companyDist := make(map[string]Distribution, len(list))
	for _, name := range list {
			companyDist[name] = Distribution{}
	}
	return companyDist
}

var ppd_report_list []string = []string {
	"ППД",
	"Відпустка",
	"Шпиталь",
	"СЗЧ",
	"Відрядження",
	"Загалом",
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
    if false {
        return
    }
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
	map[string]Distribution, []string,
	) {
	boReportCounter := make(map[string]Distribution, len(ppd_report_list))
	count_err := []string{}

	// manager_list := []ShortPersData{}	//0
	// c1_list := []ShortPersData{}			//1
	// c2_list := []ShortPersData{}			//2
	// c3_list := []ShortPersData{}			//3
	// c4_list := []ShortPersData{}			//4
	// vidZab_list := []ShortPersData{}	//5
	// vidZv_list := []ShortPersData{}		//6
	// vidTo_list := []ShortPersData{}		//7
	// mo_list := []ShortPersData{}			//8
	// total_list := []ShortPersData{}		//9

	// for _, key := range comp_list {
	// 		boReportCounter[key] = Distribution{}
	// }

	// for name, shpk_attr := range shpk_data {
	// 	person := ShortPersData{
	// 			Name:       name,
	// 			Department: shpk_attr.Department,
	// 			Rank:       shpk_attr.Rank,
	// 	}

	// 	dist := boReportCounter[ppd_report_list[5]]
	// 	incrementRankCount(&dist, shpk_attr.Rank)
	// 	boReportCounter[ppd_report_list[5]] = dist
	// }


	return boReportCounter, count_err
}