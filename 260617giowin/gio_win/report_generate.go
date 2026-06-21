package gio_win

import (
	"fmt"
	"strings"
)

// CreateReportPPD - Створення скороченого звіту для ППД
func CreateReportPPD(shpk_data map[string]Person) (map[string]Distribution, [][]ShortPersData, []string) {

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