package main

import (
	"fmt"
	"strconv"
	//	"strings"
)

func main() {

}

// weightCategory - Написать программу, которая определяет весовую категорию.
func weightCategory() {
	// var weight int
	// fmt.Scan(&weight)

	// if weight < 60 {
	// 	fmt.Println("Лёгкий вес!")
	// } else if weight >= 60 && weight < 64 {
	// 	fmt.Println("Первый полусредний вес!")
	// } else if weight >= 64 && weight < 69 {
	// 	fmt.Println("Полусредний вес!")
	// } else {
	// 	fmt.Println("Средний вес!")
	// }

}

func calculator() {
	var num1, num2 int
	var sign string
	fmt.Scan(&num1, &num2, &sign)

	if sign == "+" {
		fmt.Println(num1 + num2)
	} else if sign == "-" {
		fmt.Println(num1 - num2)
	} else if sign == "*" {
		fmt.Println(num1 * num2)
	} else if sign == "/" {
		if num2 == 0 {
			fmt.Println("На ноль делить нельзя!")
		} else {
			fmt.Println(num1 / num2)
		}
	} else {
		fmt.Println("Неверная операция")
	}
}

func colorDefinitions() {
	var colour1, colour2 string
	fmt.Scan(&colour1, &colour2)

	if colour1 == colour2 {
		if colour1 == colour2 {
			fmt.Println("Красный")
		} else {
			fmt.Println("Фиолетовый")
		}
	} else if colour1 == colour2 {
		if colour1 == colour2 {
			fmt.Println("Жёлтый")
		} else {
			fmt.Println("Оранжевый")
		}
	} else if colour1 == colour2 {
		if colour1 == colour2 {
			fmt.Println("Синий")
		} else {
			fmt.Println("Зелёный")
		}
	}
}

func colorDefinitions2() {
	var colour1, colour2 string
	var red, blue, yellow = "Красный", "Синий", "Жёлтый"

	fmt.Scan(&colour1)
	if colour1 != red && colour1 != blue && colour1 != yellow {
		fmt.Println("Ошибка: 1 цвет не правильно!")
		return
	}

	fmt.Scan(&colour2)
	if colour2 != red && colour2 != blue && colour2 != yellow {
		fmt.Println("Ошибка: 2 цвет не правильно!")
		return
	}

	if colour1 == red && colour2 == blue {
		fmt.Println("Фиолетовый")
	} else if colour1 == red && colour2 == yellow {
		fmt.Println("Оранжевый")
	} else if colour1 == blue && colour2 == yellow {
		fmt.Println("Зелёный")
	}

}

func stringTest() {
	str := "12345678"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n", str[i])

	}

}

func test5() {
	var str string

	for i := 0; i < 4; i++ {
		fmt.Scan(&str)

		runStr := []rune(str)
		for j := range len(runStr) {
			if runStr[i] == rune('0') {
				fmt.Println(i, j)
				return
			}
		}

	}

}

func test6() {
	var str string
	zero := 1
	for i := 0; i < 4; i++ {
		fmt.Scan(&str)
		divider, count := 10000000, 0
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Ошибка!")
		}

		for divider > 0 {
			digit := num / divider % 10
			count++
			if digit == 0 {
				zero = digit
				break
			}

			divider /= 10
		}

		if zero == 0 {
			fmt.Println(i, count-1)
			break
		}

	}
}

func str5() {

	str := "Hello"
	runStr := []rune(str)
	for i := range len(runStr) {
		fmt.Println(string(runStr[i]))

	}

}
