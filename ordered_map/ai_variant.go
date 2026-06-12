package main

import (
    "fmt"
)

// OrderedMap — структура з фіксованим порядком ключів
type OrderedMap struct {
    keys []string
    data map[string]interface{}
}

// NewOrderedMap створює нову OrderedMap з заданим порядком ключів
func NewOrderedMap(keys []string) *OrderedMap {
    return &OrderedMap{
        keys: keys,
        data: make(map[string]interface{}),
    }
}

// Set встановлює значення для ключа
func (om *OrderedMap) Set(key string, value interface{}) {
    // Перевірка: ключ має бути серед початково визначених
    for _, k := range om.keys {
        if k == key {
            om.data[key] = value
            return
        }
    }
    panic(fmt.Sprintf("Key '%s' is not allowed", key))
}

// Get повертає значення за ключем
func (om *OrderedMap) Get(key string) interface{} {
    return om.data[key]
}

// Iterate проходить по ключах у визначеному порядку
func (om *OrderedMap) Iterate() {
    for _, k := range om.keys {
        fmt.Printf("%s: %v\n", k, om.data[k])
    }
}

func main() {
    // Задаємо порядок ключів один раз
    keys := []string{"id", "name", "age"}
    om := NewOrderedMap(keys)

    // Заповнюємо дані
    om.Set("id", 101)
    om.Set("name", "Ivan")
    om.Set("age", 25)

    // Виводимо у визначеному порядку
    om.Iterate()
}
