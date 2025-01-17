package tasks

import "fmt"

// 1. Вывести числа от 1 до 10.
// Напиши программу, которая выводит числа от 1 до 10, используя цикл for.

func printSequenceOfNum() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

// 2. Найти сумму чисел от 1 до 100.
// Напиши программу, которая находит сумму всех чисел от 1 до 100, используя цикл.

func printSumOfNum() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
		fmt.Println(sum)
	}
}

// 3. Вывести таблицу умножения на 5.
// Напиши программу, которая выводит таблицу умножения на 5 (от 5 * 1 до 5 * 10).

func printMultiplicationTable() {
	res := 0
	for i := 1; i <= 10; i++ {
		res = i * 5
		fmt.Println(res)
	}
}

// 4. Проверка числа на простоту.
// Напиши программу, которая проверяет, является ли введённое пользователем число простым, используя цикл для проверки делителей.
func primeNum() {
	var num int
	fmt.Scan(&num)
	isPrime := true

	for i := 2; i < num; i++ {
		if num%2 == 0 {
			isPrime = false
			break
		}
	}

	if isPrime {
		fmt.Println("Число простое!")
	} else {
		fmt.Println("Число составное!")
	}
}

// 8. Обратный порядок цифр числа
// Напиши программу, которая выводит цифры введённого числа в обратном порядке. Например, для числа 1234 программа должна вывести 4321.

func printReverseSequenceOfNum() {
	for i := 4; i >= 1; i-- {
		fmt.Println(i)
	}
}
