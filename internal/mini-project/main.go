package main

import "fmt"

func main() {
	m := map[string]float64{
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

	slice := []string{" "}
	for key := range m {
		slice = append(slice, key)
	}

	fmt.Println("Добро пожадлвать в конвертер валют!")
	fmt.Println("Доступные валюты для конвертации:")

	for i := 0; i < len(slice); i++ {
		if i != 0 {
			fmt.Printf("%d. %s\n", i, slice[i])
		}	
	}

	check := true
	var sum float64
	for {
		if check {
			fmt.Print("Введите сумму в USD: ")
			fmt.Scan(&sum)
		} 

		if sum > 0 {
			fmt.Println("Выберите номер валюты для конвертации из списка выше:")
			var index int
			fmt.Scan(&index)

			isTrue := false
			for i := range slice {
				if i == index {
					isTrue = true
				}
			}

			if isTrue {
				num := 0.0
				var currency string
				for i, val := range slice {
					if i == index {
						num = sum * m[val]
						currency = val
					}
				}

				fmt.Printf("%.2f %s = %.2f %s", sum, currency, num, currency)
				break

			} else {
				fmt.Println("Неправильный выбор валюты!")
				check = false
			}

		} else {
			fmt.Println("Сумма должна превышать 0!")
			continue
		}
	}
}
