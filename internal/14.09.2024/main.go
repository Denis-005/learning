package main

import (
	"fmt"
)

func main() {
	numOfBanknotes()
}

func printMaximumNum() {
	var num1, num2 int
	fmt.Scan(&num1, &num2)

	if num1 < num2 {
		fmt.Println(num2)
	} else if num1 > num2 {
		fmt.Println(num1)
	} else {
		fmt.Println("Равны!")
	}
}

func outputMaxnumOfThree() {
	var num1, num2, num3 int
	fmt.Scan(&num1, &num2, &num3)

	if num1 > num2 && num1 > num3 {
		fmt.Println(num1)
	} else if num1 < num2 && num2 > num3 {
		fmt.Println(num2)
	} else if num1 < num3 && num2 < num3 {
		fmt.Println(num3)
	} else {
		fmt.Println("Равны!")
	}
}

func checkingTheNum() {
	var number int
	fmt.Scan(&number)

	if number > 0 {
		fmt.Println(number, "положительное число!")
	} else if number < 0 {
		fmt.Println(number, "отрицательное число!")
	} else {
		fmt.Println("Нулевое число!")
	}
}

func divisibleInteger() {
	var number int
	fmt.Scan(&number)

	if number%5 == 0 && number%11 == 0 {
		fmt.Println("Число делится на 5 и 11")
	} else {
		fmt.Println("Делится но с отатком!")
	}
}

func parityTheNum() {
	var number int
	fmt.Scan(&number)

	if number%2 == 0 {
		fmt.Println("Чётное число!")
	} else if number%2 != 0 {
		fmt.Println("Нечётное число!")
	} else {
		fmt.Println("Увы!")
	}
}

func checkingTheLeapYear() {
	var year int
	fmt.Scan(&year)

	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		fmt.Println(year, "год - високосный.")
	} else {
		fmt.Println(year, "год - невисокосный.")
	}
}

func characterDefinitions() {
	var symbol string
	fmt.Scan(&symbol)
	if symbol >= "a" && symbol <= "z" || symbol >= "A" && symbol <= "Z" {
		fmt.Println(symbol, "является алфавитом")
	} else {
		fmt.Printf("Не является алфавитом")
	}

	switch symbol {
	case "a", "b", "c", "d", "u", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "w", "v", "x", "y", "z":
		fmt.Println(symbol, "является алфавитом")
	case "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "V", "U", "W", "X", "Y", "Z":
		fmt.Println(symbol, "является алфавитом")
	default:
		fmt.Printf("Не является алфавитом")
	}
}

func characterDefinitions1() {
	fmt.Println("Введите символ!")
	var symbol string

	fmt.Scan(&symbol)
	if symbol >= "a" && symbol <= "z" || symbol >= "A" && symbol <= "Z" {
		switch symbol {
		case "a", "u", "e", "i", "o", "A", "U", "E", "I", "O":
			fmt.Println(symbol, " это гласная!")
		default:
			fmt.Println(symbol, "'это согласная!")
		}
	}
	fmt.Println("Нет такого символа алфавита!")
}

func checkingTheSymbol() {
	fmt.Println("Введите символ!")
	var symbol string

	fmt.Scan(&symbol)
	if symbol >= "a" && symbol <= "z" || symbol >= "A" && symbol <= "Z" {
		fmt.Println(symbol, "это символ алфавита!")
	} else if symbol >= "0" && symbol <= "9" {
		fmt.Println(symbol, "является цифрой!")
	} else {
		fmt.Println(symbol, "является специальным символом!")
	}
}

func checkingTheSymbol1() {
	fmt.Println("Введите символ!")
	var symbol string

	fmt.Scan(&symbol)
	if symbol >= "a" && symbol <= "z" {
		fmt.Println(symbol, "это строчный символ!")
	} else if symbol >= "A" && symbol <= "Z" {
		fmt.Println(symbol, "это заглавный символ!")
	} else {
		fmt.Println("Увы!")
	}
}

func dayOfWeek() {
	fmt.Println("Введите день недели!")
	var day int
	fmt.Scan(&day)

	switch day {
	case 1:
		fmt.Println("Monday!")
	case 2:
		fmt.Println("Tuesday!")
	case 3:
		fmt.Println("Wednesday!")
	case 4:
		fmt.Println("Thursday!")
	case 5:
		fmt.Println("Friday!")
	case 6:
		fmt.Println("Saturday!")
	case 7:
		fmt.Println("Sunday!")
	default:
		fmt.Println("Invalid Input! Please enter week number between 1-7!")
	}
}

func getMonthDays() {
	fmt.Println("Введите номер месяца!")
	var month int
	fmt.Scan(&month)

	if month == 2 {
		fmt.Println("28 days!")
	} else if month >= 1 && month <= 7 && month%2 != 0 || month >= 8 && month <= 12 && month%2 == 0 {
		fmt.Println("31 days!")
	} else {
		fmt.Println("30 days!")
	}
}

func numOfBanknotes() {
	var sum int
	fmt.Scan(&sum)

	if sum >= 500 {
		res500 := sum / 500
		sum = sum - (res500 * 500)
		fmt.Println(sum)

		// if sum >= 200 {
		// 	res200 := sum / 200
		// 	sum -= res200 * 200
		// 	fmt.Println(sum)
		// }

		// if sum >= 100 {
		// 	res100 := sum / 100
		// 	sum -= res100 * 100
		// 	fmt.Println(sum)
		// }

		// if sum >= 50 {
		// 	res50 := sum / 50
		// 	sum -= res50 * 50
		// 	fmt.Println(sum)
		// }

		// if sum >= 20 {
		// 	res20 := sum / 20
		// 	sum -= res20 * 20
		// 	fmt.Println(sum)
		// }

		// if sum >= 10 {
		// 	res10 := sum / 10
		// 	sum -= res10 * 10
		// 	fmt.Println(sum)
		// }

		// if sum >= 5 {
		// 	res5 := sum / 5
		// 	sum -= res5 * 5
		// 	fmt.Println(sum)
		// }

		// if sum >= 2 {
		// 	res2 := sum / 2
		// 	sum -= res2 * 2
		// 	fmt.Println(sum)
		// }

		// if sum >= 1 {
		// 	sum *= 1
		// 	fmt.Println(sum)
		// }
	} else {
		fmt.Println("Увы!")
	}
}

func authorization() {
	fmt.Println("Введите пароль!")
	var password string
	fmt.Scan(&password)

	if password == "qwer" {
		fmt.Println("Введите пароль ещё раз!")

		var repeatPass string
		fmt.Scan(&repeatPass)

		if repeatPass == "qwer" {
			fmt.Println("Пароль верный!")
		} else {
			fmt.Println("Пароль не верный!")
		}
	} else {
		fmt.Println("Пароль не верный, повторите попытку!")
	}
}
