package main

import "fmt"

var currencyPrices map[string]float64

func init() {
	currencyPrices = map[string]float64{
		"USD": 1.0,    // Доллар США
		"EUR": 0.88,   // Евро
		"RUB": 79.08,  // Российский рубль
		"JPY": 143.18, // Японская иена
		"CNY": 7.25,   // Китайский юань
		"GBP": 0.74,   // Британский фунт
		"KZT": 512.83, // Казахстанский тенге
		"TRY": 39.26,  // Турецкая лира
		"INR": 85.7,   // Индийская рупия
		"BRL": 5.71,   // Бразильский реал
		"AUD": 1.55,   // Австралийский доллар
		"CAD": 1.37,   // Канадский доллар
		"CHF": 0.82,   // Швейцарский франк
		"SEK": 9.53,   // Шведская крона
		"NOK": 10.11,  // Норвежская крона
	}
}

type List struct {
	Name   string
	Number int
}

func initSlice() ([16]List, map[int]string) {
	l := [16]List{}
	itemsByKey := map[int]string{}

	count := 1
	for name, _ := range currencyPrices {
		if count > len(l)-1 {
			break
		}

		itemsByKey[count] = name
		l[count] = List{name, count}
		count++
	}

	return l, itemsByKey
}

func main() {
	fmt.Println("Добро пожаловать в конвертер валют!")
	fmt.Println("Доступные валюты для конвертации:")

	list, listMap := initSlice()
	for i, val := range list {
		if i != 0 {
			fmt.Printf("%d. %s\n", val.Number, val.Name)
		}
	}

	check := true
	for {
		var sum float64
		if check {
			fmt.Print("Введите сумму в USD: ")
			fmt.Scanln(&sum)
		}

		if sum > 0 {
			fmt.Println("Выберите номер валюты для конвертации из списка выше:")
			var numCurrency int
			fmt.Scan(&numCurrency)

			currencyName, ok := listMap[numCurrency]
			if !ok {
				fmt.Println("Неправильный выбор валюты!")
				continue
			}

			price := (currencyPrices[currencyName]) * sum
			fmt.Printf("%.2f %s = %.2f %s", sum, currencyName, price, currencyName)
			break
		}

		fmt.Println("Сумма должна превышать 0!")
		continue
	}
}
