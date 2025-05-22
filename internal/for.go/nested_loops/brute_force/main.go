package main

import (
	"fmt"
)

func main() {
	newNumber()
}

func newNumber() {
	num := 9673

	divider := 1000

	for divider > 0 {
		fmt.Println(num / divider % 10)
		divider /= 10

	}

	// for divider > 0 {
	//     digit := num / divider % 10
	//     res = res * 10 + digit
	//     divider /= 10
	// }

	// num := 5060
	// divider := 1

	// digitAmount := len(strconv.Itoa(num))
	// for i := 1; i < digitAmount; i++ {
	// 	divider *= 10
	// }

	// res := num / divider % 10
	// divider /= 10

	// for divider > 0 {
	// 	digit := num / divider % 10
	// 	if digit != 5 && digit != 7 {
	// 		res = res*10 + digit
	// 	}

	// 	divider /= 10
	// }

	// fmt.Println(res)

}

func from10to99() {
	for i := 10; i <= 99; i++ {
		secondDigit := i % 10
		firstDigit := i / 10
		res := secondDigit * firstDigit * 2

		if i == res {
			fmt.Println(i, res)
		}

	}
}

func from100to999() {
	counter := 0

	for i := 100; i <= 999; i++ {
		thirdDigit := i % 10
		secondDigit := i / 10 % 10
		firstDigit := i / 100
		sum := secondDigit + firstDigit + thirdDigit

		if sum == 8 {
			counter++
		}
	}

	fmt.Println(counter)
}

func from100to999V2() {
	counter := 0

	for i := 100; i <= 999; i++ {
		currentNum := i
		sumDigit := 0

		for currentNum > 0 {
			lastDigit := currentNum % 10
			sumDigit += lastDigit
			currentNum /= 10
		}

		if sumDigit == 8 {
			counter++
		}
	}

	fmt.Println(counter)
}

// Совершенное число — натуральное число, равное сумме всех своих собственных делителей
// (то есть всех положительных делителей, отличных от самого числа).
// Например, 6 — это совершенное число, так как сумма его собственных делитей
// 1+2+3 равняется 6.
// Напишите программу, которая будет искать совершенные числа.

func perfectNum() {
	for i := 1; i < 100000000; i++ {
		sum := 0

		for a := 1; a < i; a++ {
			if i%a == 0 {
				sum += a
			}
		}

		if sum == i {
			fmt.Println(i)
		}
	}

}

// Сложное кубическое уравнение
// По данным целым числам a,b,c,d и e, найдите количество целых решений уравнения
// (ax3 + bx2 + cx + d) / ( x − e ) = 0 на отрезке [0,1000].

func complexCubicEquation() {
	// var a, b, c, d, e int
	// fmt.Scan(&a, &b, &c, &d, &e)
	// sum := 0

	// for x := 0; x <= 1000; x++ {
	// 	if (a*x*x*x+b*x*x+c*x+d) == 0 && (x-e) != 0 {
	// 		sum++
	// 		fmt.Println(sum)
	// 	}
	// }

	// fmt.Println(sum)

	var num, max int = 100, 1
	var number int

	for i := 1; i < num; i++ {
		fmt.Scan(&number)

		if number > 0 {
			if number > max {
				max = number
			}
		}

	}

	fmt.Println(max)

}

func secondMax() {
	var num, max int = 100, 1
	var number, secondLargest int

	for i := 1; i < num; i++ {
		fmt.Scan(&number)

		if number > 0 {

			if number > max {
				max = number
			}

			if secondLargest < number {
				secondLargest = max
			}
			fmt.Println(secondLargest)

		}
	}

	fmt.Println(max)
}

// package main

// import "fmt"

// func main() {
//     var n int
//     fmt.Scan(&n)

//     var num,counter,minNum,minNum1 int

//     for i := 1; i <= n; i++ {
//         fmt.Scan(&num)

//         if num < minNum {
//             minNum1 = num
//         } else if minNum1 == minNum {
//                 counter++
//         }

//      minNum = num

//     }

//     fmt.Println(counter)
// }
