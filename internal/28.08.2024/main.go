package main

import "fmt"

func main() {
	colorDefinitions()
}

// sumOfPositiveNum - Написать программу, которая считывает три числа и подсчитывет сумму только пололожительных чисел.
func sumOfPositiveNum() {

	var num1, num2, num3 int
	fmt.Scan(&num1, &num2, &num3)

	if num1 >= 0 && num2 < 0 && num3 >= 0 {
		fmt.Println(num1 + num3)
	} else if num1 < 0 && num2 >= 0 && num3 >= 0 {
		fmt.Println(num2 + num3)
	} else if num1 >= 0 && num2 >= 0 && num3 < 0 {
		fmt.Println(num1 + num2)
	} else {
		fmt.Println("Увы!")
	}
}

// beautifulInteger - Написать программу, которая
func beautifulInteger() {

	var fourDigitNum int
	fmt.Scan(&fourDigitNum)

	if fourDigitNum%7 == 0 || fourDigitNum%17 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func viewOfTriangle() {
	var num1, num2, num3 int
	fmt.Scan(&num1, &num2, &num3)

	// Равнобедренный треугольник — это треугольник, у которого две стороны одинаковой длины, а третья больше или меньше их.
	// Две равные стороны называются боковыми, третья — основанием.
	if num1 == num2 && num1 != num3 && num2 != num3 {
		fmt.Println("Равнобедренный!")
	} else if num1 == num2 && num1 == num3 && num2 == num3 { // В равностороннем треугольнике все три стороны имеют одинаковую длину.
		fmt.Println("Равностороний!")
	} else if num1 != num2 && num1 != num3 && num2 != num3 { // В разностороннем треугольнике все стороны имеют разную длину.
		fmt.Println("Разносторонний!")
	} else {
		fmt.Println("Увы")
	}
}

// averageNum - Написать программу, которая находит из 3 различных чисел среднее по величине.
func averageNum() {

	var num1, num2, num3 int
	fmt.Scan(&num1, &num2, &num3)

	if num1 < num2 && num2 < num3 || num1 > num2 && num2 > num3 {
		fmt.Println(num2)
	} else if num1 > num2 && num1 < num3 || num1 < num2 && num1 > num3 {
		fmt.Println(num1)
	} else if num1 > num3 && num3 > num2 || num1 < num3 && num3 < num2 {
		fmt.Println(num3)
	} else {
		fmt.Println("Увы")
	}
}

// getMonthDays - Написать программу, которая считывает одно целое число месяца и выводит количество дней в этом месяце.
func getMonthDays() {
	var day int
	fmt.Scan(&day)

	if day == 2 {
		fmt.Println(28)
	} else if day%2 != 0 && day <= 7 || day%2 == 0 && day >= 8 && day <= 12 {
		fmt.Println(31)
	} else if day%2 == 0 && day >= 4 && day <= 6 || day%2 != 0 && day >= 9 && day <= 11 {
		fmt.Println(30)
	} else {
		fmt.Println("Увы")
	}
}

// weightCategory() - Написать программу, которая определяет весовую категорию.
func weightCategory() {
	var weight int
	fmt.Scan(&weight)

	if weight < 60 {
		fmt.Println("Лёгкий вес!")
	} else if weight >= 60 && weight < 64 {
		fmt.Println("Первый полусредний вес!")
	} else if weight >= 64 && weight < 69 {
		fmt.Println("Полусредний вес!")
	} else {
		fmt.Println("Cредний вес!")
	}
}

// calculator -
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
		fmt.Println("Увы")
	}
}

// colorDefinitions() -
func colorDefinitions() {

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
		fmt.Println("Фиолетовый!")
	} else if colour1 != blue && colour1 != yellow || colour2 != blue && colour2 != yellow {
		fmt.Println("Красный")
	} else if colour1 == red && colour2 == yellow {
		fmt.Println("Оранжевый")
	} else if colour1 == yellow && colour2 == yellow {
		fmt.Println("Жёлтый")
	} else if colour1 == blue && colour2 == yellow {
		fmt.Println("Зелёный")
	} else if colour1 == blue && colour2 == blue {
		fmt.Println("Синий")
	}
}

func intersectionOfSegments() {
	var a1, b1, a2, b2 int
	fmt.Scan(&a1)
	fmt.Scan(&b1)
	fmt.Scan(&a2)
	fmt.Scan(&b2)

	// В первом случае у нас b1 < a2 получается что отрезки пересечение неимеют и тут нужно вывести пустое множество.
	// Во втором случае b2 < a1 получается что и тут отрезки не пересекаются и тут нужно вывести пустое множество.
	if b1 < a2 || b2 < a1 {
		fmt.Println("Пустое множество!")
	} else if b1 == a2 { // В таком случае пересечением являеться только одна точка.
		fmt.Println(b1)
	} else if b2 == a1 {
		fmt.Println(a1)
	} else if a2 > a1 && b1 < b2 { // В этом случае пересечением у нас является отрезок от a2 до b1
		fmt.Println(a2, b1)
	}
}
