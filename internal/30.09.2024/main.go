package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func main() {
	DescribeNumberV2()
}

func outputOfMinimumNominalValue() {
	var amt int
	fmt.Scan(&amt)

	switch {
	case amt >= 500:
		note500 := amt / 500
		amt -= note500 * 500
		fmt.Println("500 =", note500)
		fallthrough
	case amt >= 100:
		note100 := amt / 100
		amt -= note100 * 100
		fmt.Println("100 =", note100)
		fallthrough
	case amt >= 50:
		note50 := amt / 50
		amt -= note50 * 50
		fmt.Println("50 =", note50)
		fallthrough
	case amt >= 20:
		note20 := amt / 20
		amt -= note20 * 20
		fmt.Println("20 =", note20)
		fallthrough
	case amt >= 10:
		note10 := amt / 10
		amt -= note10 * 10
		fmt.Println("10 =", note10)
		fallthrough
	case amt >= 5:
		note5 := amt / 5
		amt -= note5 * 5
		fmt.Println("5 =", note5)
		fallthrough
	case amt >= 2:
		note2 := amt / 2
		amt -= note2 * 2
		fmt.Println("2 =", note2)
		fallthrough
	case amt >= 1:
		note1 := amt / 1
		amt -= note1 * 1
		fmt.Println("1 =", note1)
	default:
		fmt.Println("Увы")
	}
}

func checkingThreeCornersTriangle() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a > 0 && b > 0 && c > 0 {
		fmt.Println(a, b, c)
	} else {
		fmt.Println("Увы")
	}
	sum := a + b + c

	// Треугольник считается допустимым тогда и только тогда, когда сумма его углов равна 180 °.
	if sum == 180 {
		fmt.Println("Треугольник допустим")
	} else {
		fmt.Println("Треугольник не допустим")
	}
}

func checkingThreeSidesTriangle() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a < b+c && b < a+c && c < a+b {
		fmt.Println("Треугольник допустим")
	} else {
		fmt.Println("Треугольник не допустим")
	}
}

func propertiesOfTriangle() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a == b && a == c {
		fmt.Println("Равносторонний треугольник.")
	} else if a == b || a == c || b == c {
		fmt.Println("Равнобедренный треугольник.")
	} else {
		fmt.Println("Разносторонний треугольник.")
	}
}

func findingAllRootsOfQuadraticEquation() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	discriminant := b*b - b*a*c

	if discriminant > 0 {
		root1 := float64((-b + math.Sqrt(discriminant)) / (2 * a))
		root2 := float64((-b - math.Sqrt(discriminant)) / (2 * a))
		fmt.Println(root1, root2)
	} else if discriminant == 0 {
		root1 := -b / (2 * a)
		root2 := root1
		fmt.Println(root1, root2)
	} else if discriminant < 0 {
		rootCommon := -b / (2 * a)
		temp := math.Sqrt(-discriminant) / (2 * a)

		root1 := complex(rootCommon, temp)
		root2 := cmplx.Conj(root1)
		fmt.Println(root1, root2)
	}
}

func profitCalculation() {
	var cp, sp int // sp цена продажи, а cp - себестоимость.
	fmt.Scan(&cp, &sp)

	if sp > cp {
		profit := sp - cp // Прибыль
		fmt.Println("Прибыль:", profit)
	} else if cp > sp {
		loss := cp - sp // Убытки
		fmt.Println("Убытки:", loss)
	} else {
		fmt.Println("Нет прибыли и нет убытков")
	}
}

func enteringStudentGrades() {
	var a, b, c, d, e float64
	fmt.Scan(&a, &b, &c, &d, &e)

	percent := (a + b + c + d + e) / 5.0
	fmt.Println("Процент =", percent)

	if percent >= 90 {
		fmt.Println("Оценка A")
	} else if percent >= 80 {
		fmt.Println("Оценка B")
	} else if percent >= 70 {
		fmt.Println("Оценка C")
	} else if percent >= 60 {
		fmt.Println("Оценка D")
	} else if percent >= 40 {
		fmt.Println("Оценка E")
	} else {
		fmt.Println("Оценка F")
	}
}

func payrollCalculation() {
	var basicSalary float64
	fmt.Scan(&basicSalary)

	var da float64
	var hra float64

	if basicSalary <= 10000 {
		da = basicSalary * 0.8
		hra = basicSalary * 0.2
	} else if basicSalary > 10000 && basicSalary <= 20000 {
		da = basicSalary * 0.9
		hra = basicSalary * 0.25
	} else if basicSalary > 20000 {
		da = basicSalary * 0.95
		hra = basicSalary * 0.3
	}

	fmt.Println(basicSalary + da + hra)
}

func calculationOfElectricityBill() {
	fmt.Println("Введите общее количество потребленных единиц")
	var num float64
	fmt.Scan(&num)

	if num <= 50 {
		res := num * 0.50
		sum := res * 0.20
		totalRes := res + sum
		fmt.Println("Счет за электроэнергию =", totalRes)
	} else if num > 50 && num <= 100 {
		res := 25 + (num-50)*0.75
		sum := res * 0.20
		totalRes := res + sum
		fmt.Println("Счет за электроэнергию =", totalRes)
	} else if num > 100 && num <= 250 {
		res := 100 + ((num - 150) * 1.20)
		sum := res * 0.20
		totalRes := res + sum
		fmt.Println("Счет за электроэнергию =", totalRes)
	} else {
		res := 220 + ((num - 250) * 1.50)
		sum := res * 0.20
		totalRes := res + sum
		fmt.Println("Счет за электроэнергию =", totalRes)

	}
}

// Напиши функцию, которая принимает целое положительное число и возвращает true,
// если число является простым, и false в противном случае.
// Простое число — это число больше 1, которое делится только на 1 и на само себя.

func IsPrime() {
	// var num int
	// fmt.Scan
}

// Напиши функцию, которая принимает целое число и возвращает строку с описанием:
// "Положительное четное", если число больше нуля и делится на 2;
// "Положительное нечетное", если число больше нуля и не делится на 2;
// "Отрицательное четное", если число меньше нуля и делится на 2;
// "Отрицательное нечетное", если число меньше нуля и не делится на 2;
// "Ноль", если число равно нулю.

func DescribeNumber() {
	var num int
	fmt.Scan(&num)

	if num > 0 && num%2 == 0 {
		fmt.Println("Положительное четное")
	} else if num > 0 && num%2 != 0 {
		fmt.Println("Положительное нечетное")
	} else if num < 0 && num%2 == 0 {
		fmt.Println("Отрицательное четное")
	} else if num < 0 && num%2 != 0 {
		fmt.Println("Отрицательное нечетное")
	} else {
		fmt.Println("Ноль")
	}
}

func DescribeNumberV2() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}
