package main

import "fmt"

func main() {
	//loop()
	test1()

	// nums := [3]int{17, 26, 34}
	// fmt.Println(nums)

	// fmt.Println(nums[0], nums[2])
	// fmt.Println(len(nums), len(nums)-1, nums[len(nums)-1])

	// fmt.Println()

	// for i, v := range nums {
	// 	fmt.Println(i, v)
	// }

	// fmt.Println()

	// for i := range nums {
	// 	fmt.Println(nums[i])
	// }

	// fmt.Println()

	// for _, v := range nums {
	// 	fmt.Println(v)
	// }

	// fmt.Println()

	// nums[1] = 66
	// fmt.Println(nums)

	// fmt.Println()

	// Анти пример того: когда мы пытаемся записать в массив числа при этом ни как необращаясь к массиву,
	// а лишь меняя локанильную переменную v, которая является просто копией (отдельной переменной) элемента массива.
	// for _, v := range nums {
	// 	v = 55
	// 	fmt.Println(v)
	// }
	// fmt.Println(nums)

	// fmt.Println()

	// for i := range nums {
	// 	nums[i] = 55
	// }
	// fmt.Println(nums)

	// fmt.Println()

	// nums = [3]int{23, 46, 78}

	// Выводы элементов массива разноми способами.

	// fmt.Println(nums)

	// for _, v := range nums {
	// 	fmt.Print(v)
	// 	fmt.Print(" ")
	// }

	// fmt.Println()
	// for i := range nums {
	// 	fmt.Print(nums[i])
	// 	fmt.Print(" ")
	// }

	// fmt.Println()
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Print(nums[i])
	// 	fmt.Print(" ")
	// }

	// i := 0
	// for i < len(nums) {
	// 	fmt.Print(nums[i])
	// 	fmt.Print(" ")
	// 	i++
	// }

}

func loop() {
	// 	var a, b, k int
	// 	fmt.Scan(&a, &b, &k)

	// 	for i := a; i <= b; i++ {
	// 		for j := i; j > k && j <= b; j++ {
	// 			if j%2 == 0 {
	// 				fmt.Print(j, " ")
	// 			}

	// 		}

	// 	}

	// 	n := 2
	// 	b := 5
	// 	k := 3

	// 	for a := n; a <= b; a++ {
	// 		count := 0
	// 		for i := 1; i <= a; i++ {
	// 			if a%i == 0 {
	// 				fmt.Println(i)
	// 				count++
	// 			}
	// 		}
	// 		if count >= k {
	// 			fmt.Print(a, " ")
	// 		}
	// 	}

	// 	var n int
	// 	fmt.Scan(&n)

	// 	count := 0
	// 	for i := 1; i <= n; i++ {
	// 		currentNum := i
	// 		for currentNum > 0 {
	// 			digit := currentNum % 10
	// 			if digit == 7 {
	// 				count++
	// 			}
	// 			currentNum /= 10
	// 		}
	// 	}

	// 	fmt.Println(count)

	// 	var n int
	// 	fmt.Scan(&n)
	// 	var maxSum, originalNum int

	// 	for i := 1; i <= n; i++ {
	// 		currentNum := i
	// 		sum := 0
	// 		for currentNum > 0 {
	// 			lastDigit := currentNum % 10
	// 			sum += lastDigit
	// 			currentNum /= 10
	// 		}

	// 		if sum > maxSum {
	// 			maxSum = sum
	// 			originalNum = i

	// 		}
	// 	}

	// 	fmt.Println(originalNum, maxSum)

	// 	// var n int
	// 	// fmt.Scan(&n)
	// 	// //f := 1

	// 	// for i := 2; i <= n; i++ {
	// 	// 	for n%i == 0 {
	// 	// 		fmt.Print(i, " ")
	// 	// 		n = n / i
	// 	// 	}
	// 	// }

	// 	// var n int
	// 	// fmt.Scan(&n)
	// 	// maxNum := 0

	// 	// for i := 1; i < n; i++ {
	// 	// 	for n%i == 0 {
	// 	// 		if i > maxNum {
	// 	// 			maxNum = i
	// 	// 		}
	// 	// 		i++
	// 	// 	}
	// 	// }

	// 	// fmt.Println(maxNum)

	// 	// var n int
	// 	// fmt.Scan(&n)

	// 	// for {
	// 	// 	fmt.Scan(&n)

	// 	// 	if n < 10 {
	// 	// 		continue
	// 	// 	} else if n > 100 {
	// 	// 		break
	// 	// 	}
	// 	// 	fmt.Println(n)
	// 	// }

	// 	//Напишите программу, которая в последовательности трехзначных чисел находит количество всех чисел сумма цифр которых равна
	// 	//8.

	// 	//var count int

	// 	for i := 100; i <= 999; i++ {
	// 		currentNum := i
	// 		sum := 0
	// 		for currentNum > 0 {
	// 			digit := currentNum % 10
	// 			sum += digit
	// 			currentNum /= 10
	// 		}

	// 		if sum == 8 {
	// 			count++
	// 		}

	// 	}

	// 	fmt.Println(count)

	// 	// Напишите программу, которая выводит все двузначные числа, которые равны удвоенному произведению своих цифр.
	// 	// В ответе запишите найденные числа через запятую без пробелов.

	// 	for i := 10; i < 100; i++ {
	// 		currentNum := i
	// 		sum := 1
	// 		for currentNum > 0 {
	// 			digit := currentNum % 10
	// 			sum *= digit
	// 			currentNum /= 10
	// 		}
	// 		sum *= 2
	// 		if sum == i {
	// 			fmt.Println(sum)
	// 		}
	// 	}
}

func test() {
	var n int
	fmt.Scan(&n)

	divider := 1000
	var lastDigit int
	isFalse := false

	for divider > 0 {
		digit := n / divider % 10
		if digit != lastDigit {
			isFalse = true
			lastDigit = digit
			fmt.Println(lastDigit, digit)
		} else if digit == lastDigit {
			break
		}

		divider /= 10
	}

	if isFalse {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}

// Совершенное число — натуральное число, равное сумме всех своих собственных делителей (то есть всех положительных делителей,
// отличных от самого числа).
// Например, 6 — это совершенное число, так как сумма его собственных делитей 1+2+3 равняется 6.

func test1() {
	// var n int = 1000
	// var sum int

	// for i := 1; i < n; i++ {
	// 	currentNum := i
	// 	sum = 0
	// 	for a := 1; a < currentNum; a++ {
	// 		if currentNum%a == 0 {
	// 			sum += a
	// 		}
	// 	}

	// 	if sum == i {
	// 		fmt.Println(sum)
	// 	}
	// }

	// var a, b, c, d, e int
	// fmt.Scan(&a, &b, &c, &d, &e)

	// for x := 0; x <= 1000; x++ {
	// 	if x-e != 0 {
	// 		if a*x*x*x+b*x*x+c*x+d == 0 {
	// 			fmt.Print(x, " ")
	// 		}
	// 	}
	// }

	// var n, maxNum int = 1, 2
	// // fmt.Scan(&n)
	// secondLargestNum := 0

	// for n != 0 {
	// 	if n > maxNum {
	// 		maxNum = n
	// 		if secondLargestNum > maxNum {
	// 			secondLargestNum = maxNum
	// 			fmt.Println(secondLargestNum)
	// 		}
	// 	}

	// 	fmt.Scan(&n)
	// }

	// var n, n1, sum, count int
	// fmt.Scan(&n, &n1)

	// for {
	// 	sum = n + n1

	// 	if sum == 8 {
	// 		count++
	// 		fmt.Println("Выпало", n, "и", n1, "в сумме", sum, "на это потребовался", count, "бросок.")
	// 		break
	// 	} else {
	// 		count++
	// 		fmt.Println("Выпало", n, "и", n1, "в сумме", sum, "бросаем еще раз.")
	// 	}
	// 	fmt.Scan(&n, &n1)
	// }

	//var n int = 5
	//fmt.Scan(&n)

	// for i := 1; i <= n; i++ {
	// 	for a := 1; a <= n; a++ {
	// 		res := a * i
	// 		fmt.Println(a, "*", i, "=", res)
	// 	}
	// }

	// arr := [10]int{1, 2, 3, 2, 1, 1, 2, 3, 2, 1}
	// arrReversed := []int{}

	// for i := len(arr) - 1; i >= 0; i-- {
	// 	fmt.Println(i)
	// 	arrReversed = append(arrReversed, arr[i])
	// }

	// if arr == [10]int(arrReversed) {
	// 	fmt.Println("Это палиндром!")
	// } else {
	// 	fmt.Println("Не палиндром!")
	// }

	// Замена символов в строке
	// Напишите функцию PrintReplaced, которая принимает строку и выводит в консоль новую строку, в которой все буквы "у" заменены на буквы "а".

	var word string
	fmt.Scan(&word)
	// Кукушка!

	for _, r := range word {
		if r == 'у' {
			r = 'a'
		}

		fmt.Print(string(r))

	}

}
