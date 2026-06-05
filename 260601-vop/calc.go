package main

import (
	"regexp"
	"strings"
)

func Divisions() []string {
	return []string{
		"упр 3 бо",
		"1", "2", "3", "4",
		"від.заб./3 бо",
		"від.то/3 бо",
		"м.п./3 бо",
		"total",
	}
}

// getPlatoon повертає номер роти як рядок ("1","2","3","4") або "" якщо не знайдено
func getPlatoon(division string) string {
	pattern := regexp.MustCompile(`^.*(1|2|3|4)\/3.*$`)
	m := pattern.FindStringSubmatch(division)
	if len(m) >= 2 {
		return m[1]
	}
	return ""
}

// personnelCounter приймає rows — зріз рядків, де кожен рядок — зріз string,
// індекси: 2 = rank, 3 = fullName, 4 = division. Повертає map[string]map[string]int
func PersonnelCounter(rows [][]string) map[string]map[string]int {
	// Ініціалізація лічильників
	new_counts := func() map[string]int {
		return map[string]int{"offi": 0, "serg": 0, "sold": 0}
	}
	divisions := Divisions()
	divisionsCounter := make(map[string]map[string]int, len(divisions))
	for _, k := range divisions {
		divisionsCounter[k] = new_counts()
	}

	// Починаємо з рядка з індексом 3 (min_row=4 в Python) — тобто пропускаємо перші 3 рядки
	for i := 3; i < len(rows); i++ {
		row := rows[i]
		// Перевіряємо наявність потрібних індексів
		if len(row) <= 4 {
			continue
		}
		rank := strings.TrimSpace(row[2])
		fullName := strings.TrimSpace(row[3])
		division := strings.TrimSpace(row[4])

		if rank == "" || fullName == "" || division == "" {
			continue
		}

		var platoon string
		switch division {
		case "упр 3 бо", "від.заб./3 бо", "від.то/3 бо", "м.п./3 бо":
			platoon = division
		default:
			platoon = getPlatoon(division)
			if platoon == "" {
				continue
			}
		}

		// Розподіл за званням
		if strings.HasSuffix(rank, "олдат") {
			divisionsCounter[platoon]["sold"] += 1
			divisionsCounter["total"]["sold"] += 1
		} else if strings.HasSuffix(rank, "ержант") {
			divisionsCounter[platoon]["serg"] += 1
			divisionsCounter["total"]["serg"] += 1
		} else {
			divisionsCounter[platoon]["offi"] += 1
			divisionsCounter["total"]["offi"] += 1
		}
	}

	return divisionsCounter
}
