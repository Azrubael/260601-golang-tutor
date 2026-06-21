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

var list_for_ppd_report []string = []string {
	"ППД",
	"Відпустка",
	"Шпиталь",
	"СЗЧ",
	"Відрядження",
	"Загалом",
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

// categorizePersonnel adds person to appropriate list and updates counter
func categorizePersonnel(
	person ShortPersData,
	list *[]ShortPersData,
	counters map[string]Distribution,
	key, rank string,
	) {
    if false {
        return
    }
    *list = append(*list, person)
    dist := counters[key]
    incrementRankCount(&dist, rank)
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
	ppdReportCounter := make(map[string]Distribution, len(list_for_ppd_report))
	count_err := []string{}

	ppd_list := []ShortPersData{}
	vac_list := []ShortPersData{}
	hosp_list := []ShortPersData{}
	szch_list := []ShortPersData{}
	asmt_list := []ShortPersData{}

	for _, key := range list_for_ppd_report {
			ppdReportCounter[key] = Distribution{}
	}

	for name, shpk_attr := range shpk_data {
		person := ShortPersData{
				Name:       name,
				Department: shpk_attr.Department,
				Rank:       shpk_attr.Rank,
		}

		dist := ppdReportCounter[list_for_ppd_report[5]]
		incrementRankCount(&dist, shpk_attr.Rank)
		ppdReportCounter[list_for_ppd_report[5]] = dist

		if shpk_attr.Assignment == "ППД" {
				categorizePersonnel(person, &ppd_list, ppdReportCounter,
					list_for_ppd_report[0], shpk_attr.Rank)
		} else if shpk_attr.Assignment != "" {
				categorizePersonnel(person, &asmt_list, ppdReportCounter,
					list_for_ppd_report[4], shpk_attr.Rank)
		}

		if shpk_attr.Vacation_now != "" && shpk_attr.Assignment == "" {
				categorizePersonnel(person, &vac_list, ppdReportCounter,
					list_for_ppd_report[1], shpk_attr.Rank)
		} else if shpk_attr.Vacation_now != "" && shpk_attr.Assignment != "" {
				err_msg := fmt.Sprintf("Потрібна перевірка актуального статусу для %s: відпустка чи відрядження?", name)
				fmt.Println(err_msg)
				count_err = append(count_err, err_msg)
		}

		if shpk_attr.Hospital != "" && shpk_attr.Assignment == "" {
				categorizePersonnel(person, &hosp_list, ppdReportCounter,
					list_for_ppd_report[2], shpk_attr.Rank)
		} else if shpk_attr.Hospital != "" && shpk_attr.Assignment != "" {
				err_msg := fmt.Sprintf("Потрібна перевірка актуального статусу для %s: відпустка чи відрядження?", name)
				fmt.Println(err_msg)
				count_err = append(count_err, err_msg)
		}

		if shpk_attr.Szch != "" && shpk_attr.Assignment == "" {
				categorizePersonnel(person, &szch_list, ppdReportCounter,
					list_for_ppd_report[3], shpk_attr.Rank)
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
func PrepareReportBO(shpk_data map[string]Person) map[string]Distribution {
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