package main

import (
	"fmt"
	"math"
)

func main() {
	viewOfTriangle()
}

// findMaxNum () - Определить какое из трех введенных пользователем чисел максимальное и вывести его на экран.
func findMaxNum() {

	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a > b && a > c {
		fmt.Println(a)
	} else if b > a && b > c {
		fmt.Println(b)
	} else {
		fmt.Println(c)
	}
}

// findAverageNum - Среди трех чисел найти среднее. Если среди чисел есть равные, вывести сообщение "Ошибка".
func findAverageNum() {

	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a > b && a < c || a < b && a > c {
		fmt.Println(a)
	} else if b > a && b < c || b > c && b < a {
		fmt.Println(b)
	} else if c > a && c < b || c < a && c > b {
		fmt.Println(c)
	} else {
		fmt.Println("Увы")
	}
}

// printOddNum - Из двух чисел с разной четностью вывести на экран нечетное число.
func printOddNum() {

	var a, b int
	fmt.Scan(&a, &b)

	if a%2 != 0 {
		fmt.Println(a)
	} else if b%2 != 0 {
		fmt.Println(b)
	} else {
		fmt.Println("Чётное числа")
	}
}

// determineMultiplicityNumbers - Вводятся два числа (большее и меньшее).
// Определить, кратно ли первое число второму, то есть делится ли первое число нацело на второе.
// Вывести на экран сообщение об этом, а также остаток от деления, если первое число не кратно второму.
// Кратное - Число, делящееся на данное целое число без остатка.
func determineMultiplicityNumbers() {

	var maxNum, minNum int
	fmt.Scan(&maxNum, &minNum)

	if maxNum%minNum == 0 {
		fmt.Println("Первое число кратно второму!")
	} else {
		fmt.Println("Первое число не кратно второму!")
		fmt.Println(maxNum % minNum)
	}
}

// determiningQuarterOnCoordinatePlane() - Требуется написать программу, определяющую по координатам точки,
// в какой четверти она находится. Координаты точки вводятся с клавиатуры.
func determiningQuarterOnCoordinatePlane() {

	var x, y int
	fmt.Scan(&x, &y)

	if x > 0 && y > 0 {
		fmt.Println("Четверть I")
	} else if x < 0 && y > 0 {
		fmt.Println("Четверть II")
	} else if x < 0 && y < 0 {
		fmt.Println("Четверть III")
	} else if x > 0 && y < 0 {
		fmt.Println("Четверть IV")
	} else {
		fmt.Println("Увы")
	}
}

// determineExistenceOfTriangleSides() - Вводятся длины трех сторон предполагаемого треугольника.
// Определить, может ли существовать треугольник с такими сторонами при условии что,
// треугольник существует только тогда, когда ни одна его сторона не превышает сумму двух других.
func determineExistenceOfTriangleSides() {

	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a < (b+c) && b < (a+c) && c < (a+b) {
		fmt.Println("Треугольник существует.")
	} else {
		fmt.Println("Треугольник не существует.")
	}
}

// doesTheCoordinatePointBelongCircle() - С клавиатуры задаются координаты точки,
// а также радиус круга, центр которого находится в начале координат.
// Определить, принадлежит ли данная точка кругу.
func doesTheCoordinatePointBelongCircle() {

	// Если выбрать точку на координатной плоскости, то можно увидеть,
	// что проекции ее координат на оси x и y являются катетами прямоугольного треугольника.
	// А гипотенуза этого прямоугольного треугольника как раз показывает расстояние от начала
	// координат до точки.
	// Таким образом, если длина гипотенузы будет не больше радиуса круга,
	// то точка будет принадлежать кругу; иначе она будет находиться за его пределами.
	// Прямоугольный треугольник, катеты c1 и c2 и гипотенуза (h)
	// Катет — одна из сторон прямоугольного треугольника,
	// образующая прямой угол. Противоположная прямому углу сторона называется гипотенузой.
	// Длину гипотенузы вычисляется по теореме Пифагора:
	// квадрат гипотенузы равен сумме квадратов катетов.
	// Откуда гипотенуза равна квадратному корню из суммы квадратов катетов.

	var x, y, r float64
	fmt.Scan(&x, &y, &r)
	hypotenuse := math.Sqrt(x * y)

	if hypotenuse <= r {
		fmt.Println("Точка принадлежит кругу!")
	} else {
		fmt.Println("Точка не принадлежит кругу")
	}
}

// calculatingAreaOfGeometricShapes() - Написать программу, которая в зависимости от выбора
// пользователя вычисляет площадь одной из трех геометрических фигур:
// прямоугольника, треугольника или круга.
func calculatingAreaOfGeometricShapes() {

	fmt.Println("Введите цифру!")
	var figure int
	fmt.Scan(&figure)

	if figure == 1 {
		var a, b float64
		fmt.Scan(&a, &b)
		square := a * b // Площадь прямоугольника равна произведению его двух сторон (длины и ширины).
		fmt.Println(square)
	} else if figure == 2 {
		var a, b, c float64
		fmt.Scan(&a, &b, &c)
		p := (a + b + c) / 2                                 // p - это полупериметр, a, b, c - длины сторон. Полупериметр равен половине периметра, то есть половине суммы сторон.
		square := math.Sqrt(p * (p - a) * (p - b) * (p - c)) // Площадь треугольника вычисляется по формуле Герона.
		fmt.Println(square)
	} else if figure == 3 {
		var r float64
		fmt.Scan(&r)
		square := math.Pi * (r * r) // Чтобы получить площадь круга по радиусу в Golang, можно использовать следующую формулу: площадь круга = π * (радиус * радиус):
		fmt.Println(square)
	} else {
		fmt.Println("Ошибка ввода")
	}
}

// leapYear() - С клавиатуры вводится год. Программа должна определять високосный это год или нет. Вывести на экран соответствующую надпись, а также количество дней в году.
func leapYear() {

	fmt.Println("Введите год!")
	var year int
	fmt.Scan(&year)

	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		fmt.Println("Это високосный год!")
		fmt.Println("366 дней в году!")
	} else {
		fmt.Println("Это невисокосный год!")
		fmt.Println("365 дней в году!")
	}
}

// enteringPassword() - Написать программу в которой пользователь будет вводить пароль дважды.
func enteringPassword() {

	fmt.Println("Введите пароль для входа!")
	var password string
	fmt.Scan(&password)

	if password == "qwerty" {
		fmt.Println("Верно!")
		fmt.Println("Введите пароль ещё раз!")

		fmt.Scan(&password)
		if password == "qwerty" {
			fmt.Println("Пароль принят!")
		} else {
			fmt.Println("Не верный пароль!")
		}
	} else {
		fmt.Println("Не верный пароль!")
	}
}

// definitionsOfNum() - Написать программу, которая определяет является число чётным или нечётным.
func definitionsOfNum() {

	var num int
	fmt.Scan(&num)

	if num%2 == 0 {
		fmt.Println("Чётное число!")
	} else if num%2 == 1 {
		fmt.Println("Нечётное число!")
	} else {
		fmt.Println("Увы")
	}
}

// ratioOfAmounts() - Написать программу, которая проверяет, что заданное четырёхзначного число
// подходит для следующего соотношения: сумма первой и последней цифр равно разности второй и третей цифр.
func ratioOfAmounts() {

	var num int
	fmt.Scan(&num)

	num1 := num / 1000
	num2 := num % 1000 / 100
	num3 := num % 100 / 10
	num4 := num % 10

	sum := num1 + num4  // Сложение 1 и 4 цифры
	sum1 := num2 - num3 // Вычитание 2 и 3 цифры

	if sum == sum1 {
		fmt.Println("Равные суммы!")
	} else {
		fmt.Println("Не равные суммы!")
	}
}

// ageLimit() - Написать программу, которая определяет возрастное ограничение.
func ageLimit() {

	fmt.Println("Введите ваш возраст для получения доступа!")
	var age int
	fmt.Scan(&age)

	if age >= 18 {
		fmt.Println("Доступ разрешён!")
	} else {
		fmt.Println("Доступ запрещён!")
	}
}

// sequenceOfNumb() - Написать программу, которая определяет, являются ли три заданных числа последовательными
// членами арифметической прогрессии.
// Арифметическая прогрессия - это когда последовательность чисел увеличивается или уменьшается
// на одно и тоже количесво чисел.
func sequenceOfNumb() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	// Нужно сравнить разность последующих цифр, она одинакова.
	if b-a == c-b { // Элемент b отличается от элемента a на такую же велечину что и элемент c отличается от элемента b.
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

// smallestNum() - Написать программу, которая определяет наименьшее из двух чисел.
func smallestNum() {

	var a, b float64
	fmt.Scan(&a, &b)
	//fmt.Println(math.Min(a, b))

	if a < b {
		fmt.Println(a)
	} else if a > b {
		fmt.Println(b)
	} else {
		fmt.Println("Увы")
	}
}

// ageGroup() - Написать программу, которая определяет какой возрастной группе относится пользователь.
func ageGroup() {

	fmt.Println("Введите возраст!")
	var age int
	fmt.Scan(&age)

	if age <= 13 {
		fmt.Println("Детсво")
	} else if age >= 14 && age <= 24 {
		fmt.Println("Молодость")
	} else if age >= 25 && age <= 59 {
		fmt.Println("Зрелость")
	} else {
		fmt.Println("Старость")
	}
}

// getMonthDays() -
func getMonthDays() {

	var monthNumber int
	fmt.Scan(&monthNumber)

	if monthNumber%2 != 0 {
		fmt.Println(31)
	} else {
		if monthNumber == 2 {
			fmt.Println(28)
		} else {
			fmt.Println(30)
		}
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
		fmt.Println("Средний вес!")
	}
}

// colorDefinitions() -
func colorDefinitions() {

	var red, blue, yellow string = "Красный", "Синий", "Жёлтый"
	var colour string
	fmt.Scan(&colour)

	if colour == blue {
		if colour == red {
			fmt.Println("Красный")
		} else {
			fmt.Println("Фиолетовый")
		}
	} else if colour == yellow {
		if colour == yellow {
			fmt.Println("Жёлтый")
		} else {
			fmt.Println("Оранжевый")
		}
	} else if colour == yellow {
		if colour == blue {
			fmt.Println("Синий")
		} else {
			fmt.Println("Зелёный")
		}
	}

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
