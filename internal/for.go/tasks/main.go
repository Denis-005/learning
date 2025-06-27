package main

import (
	"fmt"
)

func main() {
	primeNum()
}

// 1. Вывести числа от 1 до 10
// Напиши программу, которая выводит числа от 1 до 10, используя цикл for.

func outputOfNum() {
	var num int
	fmt.Scan(&num)

	i := 1
	for i <= num {
		fmt.Println(i)
		i++
	}
}

// 2. Найти сумму чисел от 1 до 100
// Напиши программу, которая находит сумму всех чисел от 1 до 100, используя цикл.

func sumNum() {
	var num int
	fmt.Scan(&num)
	sum := 0

	i := 1
	for i <= num {
		sum += i
		i++
	}
	fmt.Println(sum)
}

// 3. Вывести таблицу умножения на 5
// Напиши программу, которая выводит таблицу умножения на 5 (от 5 * 1 до 5 * 10).

func outputMultiplicationTableNum() {
	var num int
	fmt.Scan(&num)

	i := 1
	for i <= num {
		fmt.Println(i * 5)
		i++
	}
}

// 4. Проверка числа на простоту
// Напиши программу, которая проверяет, является ли введённое пользователем число простым, используя цикл для проверки делителей.

func primeNum() {
	var num int
	fmt.Scan(&num)
	isPrime := true

	i := 2
	for i < num {
		if num%2 == 0 {
			isPrime = false
		}
		i++
	}

	if isPrime {
		fmt.Println("Простое")
	} else {
		fmt.Println("Составное")
	}
}

// 5. Факториал числа
// Напиши программу, которая вычисляет факториал числа, введённого пользователем, используя цикл for.

func factorialNum() {
	var num int
	fmt.Scan(&num)
	f := 1

	i := 1
	for i <= num {
		f *= i
		i++
	}
	fmt.Println(f)
}

// 6. Найти наибольший общий делитель (НОД) двух чисел
// Напиши программу, которая находит НОД двух чисел, используя алгоритм Евклида в цикле.

// 7. Обратное число.
// Дано натуральное число N. Переставьте цифры числа в обратном порядке.

func reverseNum() {
	var num int
	fmt.Scan(&num)
	//sum := 0
	lastDigit := 0

	i := 1
	for i <= num {
		lastDigit = num % 10
		//sum = lastDigit
		num /= 1
		fmt.Println(lastDigit)
	}
	//fmt.Println(sum)
}

// 8. Количество элементов, которые больше предыдущего.
// Последовательность состоит из натуральных чисел и завершается числом 0.
// Определите, сколько элементов этой последовательности больше предыдущего элемента. Числа, следующие за числом 0, считывать не нужно.

func sequenceElements() {
	var lastNum int
	fmt.Scan(&lastNum)

	counter := 0

	// Пока наше очередное число не ноль.
	var num int
	fmt.Scan(&num)
	for num != 0 {
		if num > lastNum {
			counter++
		}

		lastNum = num
		fmt.Scan(&num)
	}
	fmt.Println(counter)

}
