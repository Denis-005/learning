package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	res, err := getStudentType("555555")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}

// outputOfEachCharacter - Вывести каждый символ на новой строке. (Итерироваться по строке)
func outputOfEachCharacter() {
	symbol := "478"

	for i := 0; i < len(symbol); i++ {
		sim := string(symbol[i])

		//fmt.Println(sim)
		num, err := strconv.Atoi(sim)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(num)
	}

}

func performanceStudent() {
	// var estimation int = 54434534           // Объявили и проинициализировали переменную estimation.
	// var s string = strconv.Itoa(estimation) // Конвертируем переменную estimation в строку.
	// res := 0                                // Создали накопительную переменную.

	// for i := 0; i < len(s); i++ { // При помощи цикла for проходимся по каждому символу строки.
	// 	str := (string(s[i])) // Присвоили в переменную str символы строки s.
	// }
	// num, err := strconv.Atoi(str) // Создаём переменную num и присваиваем сконвертируемые в тип int символы строки str.
	// if err != nil {               // Обрабатываем ошибки при помощи условной конструкции
	// 	fmt.Println(err)
	// 	return // return — это оператор в языке Go, который используется для возврата результата из функции.
	// }

	// res += num // При помощи составного оператора присваивания складываем все значения и присваиваем в накопительную переменную.
	// markAmount := float32(len(s))

	// average := float32(res) / markAmount // Присваиваем в переменную res1 вычисление средней арифметической.
	// fmt.Println(average)

	// 	if res1 >= 3 && res1 <= 3.7 { // Через условную конструкцию узнаём успеваемость студента.
	// 		fmt.Println("Троешник")
	// 	} else if res1 >= 3.8 && res1 < 4.5 {
	// 		fmt.Println("Хорошист")
	// 	} else if res1 > 4.5 {
	// 		fmt.Println("Отличник")
	// 	} else {
	// 		fmt.Println("Увы!")
	// 	}
}

func getStudentType(marks string) (string, error) {
	markSum := 0
	for i := 0; i < len(marks); i++ {
		symbol := (string(marks[i]))

		mark, err := strconv.Atoi(symbol)
		if err != nil { // если ошибка
			return "", err
		}

		markSum += mark
	}

	markAmount := float32(len(marks))
	average := float32(markSum) / markAmount // Присваиваем в переменную res1 вычисление средней арифметической.

	if average >= 3 && average <= 3.7 { // Через условную конструкцию узнаём успеваемость студента.
		return "Троешник", nil
	} else if average >= 3.8 && average < 4.5 {
		return "Хорошист", nil
	}

	return "Отличник", nil
}

func str2() {
	str := "Hello"
	symbol1 := rune(str[0])
	symbol2 := rune(str[2])
	symbol3 := rune(str[4])
	fmt.Printf("%c", symbol1) // Printf — это функция в языке Go, которая используется для печати форматированной строки в консоль.
	fmt.Printf("%c", symbol2)
	fmt.Printf("%c", symbol3)

	res := strings.Split(str, "") // strings.Split — это функция в Golang, которая используется для разделения строки на части с использованием определённого разделителя.
	fmt.Println(res[0])           // Здесь разделителем является пустая строка.
	fmt.Println(res[2])
	fmt.Println(res[4])
}

func printlengthOfString() {
	str := "height"       // Объявили и проинициализировали переменную str.
	fmt.Println(len(str)) // При помощи фукции len вывели количество символов в строке.

	if len(str) > 5 { // В блоке if указываем диапазон длины строки len(str).
		res := str[0:3]          // В переменную res присваиваем первые три символа строки.
		res1 := str[len(str)-3:] //В переменную res1 присваиваем последние три символа строки.
		fmt.Println(res, res1)   // Вывод переменных.
	} else {
		for i := 0; i <= len(str); i++ { // В блоке else при помощи цикла for выводим первый символ строки столько раз,
			fmt.Println(string(str[0])) // какова будет длина строки.
		}
	}
}

func convertNumber() {
	// strconv.Atoi(). Это самый простой способ конвертации строки в целое число.
	// Функция принимает строку в качестве аргумента и возвращает целое число или ошибку,
	// если строка не представляет собой допустимое целое число.
	// Важно помнить, что при преобразовании чисел могут возникать ошибки,
	// например, если число нельзя преобразовать в нужный тип, или если строка содержит неправильный формат числа.
	// Поэтому при преобразовании чисел в Go рекомендуется обрабатывать ошибки, чтобы избежать непредвиденных ошибок выполнения программы.
	s := "123fgdgdf"
	var num, err = strconv.Atoi(s)
	if err == nil {
		// преобразование прошло успешно
		fmt.Println(num + 4)
	} else {
		// возникла ошибка во время преобразования
		fmt.Println("Ошибка:", err)
	}
}

func convertNumber1() {
	symbol := "478"

	for i := 0; i < len(symbol); i++ {
		sim := string(symbol[i])

		//fmt.Println(sim)
		num, err := strconv.Atoi(sim)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(num)
	}
}

func number1() {

	num := 123456
	num1 := num / 1000 % 100
	num2 := num / 10 % 100
	//num3 := num % 10000 / 100
	num4 := num / 100000
	num5 := num % 10
	fmt.Println(num1, num2, num4, num5)

	var s string = strconv.Itoa(num4)
	var s1 string = strconv.Itoa(num5)
	s3 := s + s1
	//fmt.Print(s3)

	var num6, err = strconv.Atoi(s3)
	if err == nil {
		fmt.Println(num6)
	} else {
		fmt.Println("Ошибка!")
	}

	res := (num1 + num2) / num6
	fmt.Println(res)
}

func number2() {
	var num = 789354
	num1 := num / 1000 % 100
	num2 := num / 10 % 100
	num3 := num / 100000
	nun4 := num % 10
	fmt.Println(num1, num2, num3, nun4)

	var str string = strconv.Itoa(num3)
	var str1 string = strconv.Itoa(nun4)
	s3 := str + str1
	//fmt.Print(s3)

	var num5, err = strconv.Atoi(s3)
	if err == nil {
		fmt.Println(num5)
	} else {
		fmt.Println("Увы")
	}

	res := (num1 + num2) / num5
	fmt.Println(res)
}

func srt1() {
	var s = "Hello"
	res := ""

	for i := 0; i < 3; i++ {
		res += s + ","
	}

	fmt.Println(res[:len(res)-1])
}
