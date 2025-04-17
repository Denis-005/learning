package main

import (
	"fmt"
	"math"
)

func main() {
	multiplicityCondition()
	//isPrime()
	//isPrimeV2()
}

func divisorsOnSegment() {
	var a, b, k int
	fmt.Scan(&a, &b, &k)

	for num := a; num <= b; num++ {
		counter := 0

		for i := 1; i <= num; i++ {
			if num%i == 0 {
				counter++
			}
		}

		if counter >= k {
			fmt.Print(num)
			fmt.Print(" ")
		}
	}

}

func test() {
	num := 10
	counter := 0

	for i := 1; i <= num; i++ {
		if num%i == 0 {
			counter++
		}
	}
	fmt.Println(counter)
}

// Количество цифр 7.
// Посчитать, сколько раз встречается цифра 7 в последовательности чисел от 1 до N, включая N.
func robot() {
	var num int
	fmt.Scan(&num)
	sum := 0

	for i := 1; i <= num; i++ {
		currentNum := i
		digit := 1

		for digit <= currentNum {
			res := currentNum % 10
			if res == 7 {
				sum++
			}

			currentNum /= 10
		}
	}
	fmt.Println(sum)
}

// Максимальная сумма цифр.
// Среди натуральных чисел от 1 до N (включая N), найдите число, наибольшее по сумме цифр.
// Вывести это число и сумму его цифр.
func test1() {
	var num int
	fmt.Scan(&num)
	sum := 0
	max := 0

	for i := 1; i < num; i++ {
		currentNum := i

		sum = 0
		for currentNum > 0 {
			lastDigit := currentNum % 10
			sum += lastDigit

			if sum > max {
				max = sum
				fmt.Println(max)
			}
			currentNum /= 10
		}
	}
	fmt.Println(max)
}

// Разложение на простые множители*
// Вывести представление целого числа N в виде произведения простых чисел.
// Выведите число, если оно одно, или список чисел в порядке не убывания (в порядке возрастания) через пробел.

func isPrime() {
	num := 2147483647
	isPrime := true

	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}

	fmt.Println(isPrime)

}

func isPrimeV2() {
	num := 14
	isPrime := true

	limiter := int(math.Ceil(math.Sqrt(float64(num))))

	for i := 2; i <= limiter; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}

	fmt.Println(isPrime)

}

func decomposingNumIntoPrimeFactor() {
	// var num int
	// fmt.Scan(&num)
	num := 338724

	a := 2 // Потенциальное простое число.

	for num > 1 {
		isPrime := true
		limiter := int(math.Ceil(math.Sqrt(float64(a))))

		for i := 2; i < limiter; i++ {
			if a%i == 0 {
				isPrime = false
				break
			}
		}

		// Если НЕ простое то переходим к следующему числу.
		if !isPrime {
			a++
			continue
		}

		// Если мы дошли сюда, то переменная "a" - точно простое.

		// Если изначальное число НЕ делится нацело на простое,
		// то переходим к следующему простому числу.
		if num%a != 0 {
			a++
			continue
		}

		num /= a
		fmt.Print(a, " ")
		a = 2
	}
}

// decomposingNumIntoPrimeFactorV2 - В этот версии мы не проверяем на простоту.
func decomposingNumIntoPrimeFactorV2() {
	var num int
	fmt.Scan(&num)

	a := 2

	for num > 1 {

		if num%a != 0 {
			a++
			continue
		}

		num /= a
		fmt.Print(a, " ")
		a = 2
	}
}

func decomposingNumIntoPrimeFactorV3() {
	var num int
	fmt.Scan(&num)

	for i := 2; i <= num; i++ {
		for num%i == 0 {
			num /= i
			fmt.Print(i, " ")
		}
	}
}

func multiplicityCondition() {

	var n, n1, sum, count int
	fmt.Scan(&n, &n1)

	for {
		sum = n + n1

		if sum == 8 {
			count++
			fmt.Println("Выпало", n, "и", n1, "в сумме", sum, "на это потребовался", count, "бросок.")
			break
		} else {
			count++
			fmt.Println("Выпало", n, "и", n1, "в сумме", sum, "бросаем еще раз.")
		}
		fmt.Scan(&n, &n1)
	}
}
