package main

import "fmt"

func main() {
	// i := 1
	// for ; ; i++ {
	// 	if i <= 10 {
	// 		fmt.Println(i)
	// 	} else {
	// 		break
	// 	}
	// }

	// i := 1
	// for {
	// 	if i > 10 {
	// 		break
	// 	}

	// 	fmt.Println(i)
	// 	i++
	// }

	// for i := range 10 {
	// 	fmt.Println(i)
	// }
	outputOfInverseBinaryNum()
}

// Напиши программу, которая выводит числа от 1 до 10, используя цикл for.
func printSequenceOfNum() {

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

}

// Напиши программу, которая находит сумму всех чисел от 1 до 100, используя цикл.
func printSumOfNum() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
		fmt.Println(sum)
	}
}

// Напиши программу, которая выводит таблицу умножения на 5 (от 5 * 1 до 5 * 10).
func printMultiplicationTable() {
	res := 0
	for i := 1; i <= 10; i++ {
		res = i * 5
		fmt.Println(res)
	}
}

// Напиши программу, которая выводит цифры введённого числа в обратном порядке. Например, для числа 1234 программа должна вывести 4321.
func printReverseSequenceOfNum() {
	for i := 4; i >= 1; i-- {
		fmt.Println(i)
	}
}

// Напиши программу, которая проверяет, является ли введённое пользователем число простым, используя цикл для проверки делителей.
func printPrimeNum() {

}

func for1() {
	for i := 10; i <= 90; i += 10 {
		if i > 70 {
			break
		}

		if i%2 == 1 {
			continue
		}

		fmt.Println(i)
	}

}

// func printPrimeNum() {
// 	num := 13
// 	for i := 2; i < num; i++ {
// 		fmt.Println(i)
// 	}

// }

// 5>> Напиши функцию, которая принимает целое положительное число и возвращает true, если число является простым,
// и false в противном случае. Простое число — это число больше 1, которое делится только на 1 и на само себя.

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

// Замечательное число.
// Дано натуральное число - N. Определите является ли данное число замечательным.
// Число называется замечательным, если оно делится на сумму своих цифр без остатка.

func wonderfulNum() {
	var n int
	fmt.Scan(&n)
	t := n
	//fmt.Println(t)
	d := 0
	sum := 0
	isNum := true

	i := 0
	for i < n {
		d = n % 10
		//fmt.Println(d)
		sum += d
		//fmt.Println(sum)
		n /= 10
		//fmt.Println(n)
	}

	t %= sum
	fmt.Println(t)

	if t%sum != 0 {
		isNum = false
	}

	if isNum {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func reverseNum() {
	var n int
	fmt.Scan(&n)
	d := 0
	s := 0

	i := 1
	for i <= n {
		d = n % 10
		fmt.Println(d)
		s = d
		fmt.Println(s)
		if s <= n {
			fmt.Print(s)
		}
		n /= 10
	}
}

func outputOfInverseBinaryNum() {
	var n int
	fmt.Scan(&n)

	for n > 0 {
		fmt.Print(n % 2)
		n /= 2
	}
}
