package gio_win

func ExperimentBO(shpk_data map[string]Person) (map[string]map[string]Distribution, []string) {
	type Distr struct {
		Offi  int
		Serg  int
		Sold  int
		Total int
	}

	type ReportData struct {
		Name       string
		Department string
		Rank       string
	}

	var bo []string = []string{
		"Відпустка",
		"Шпиталь",
		"Навчання",
		"Відрядження",
		"ВОП",
		"СЗЧ",
		"ППД",
		"Загалом",
	}

	var companies []string = []string{
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

	boReportCounter := make(map[string]map[string]Distribution, len(companies))
	count_err := []string{}
	for _, c := range companies {
		for _, b := range bo {
			boReportCounter[c][b] = Distribution{}
		}
	}

	return boReportCounter, count_err
}