package main

import (
	"log"
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
	
// MakeListOfCompanies - Створення списку підрозділів з нульовими даними розподілу
func MakeListOfCompanies() map[string]Distribution {
	companyDist := make(map[string]Distribution, len(list_of_companies))
	for _, name := range list_of_companies {
			companyDist[name] = Distribution{}
	}
	return companyDist
}
