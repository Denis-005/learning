package main

import "fmt"

func main2() {
	num := 13
	count := 0
	for i := 2; i < num; i++ {
		// Это условие того что число простое так как не делиться нацело.
		if num%2 != 0 {
			count++
		}
	}

	// Это количество делителей на которое мы поделили наше исходное число, если же наше число простое,
	// то таких делителей должно быть на 2 меньше чем наше число.
	deliteli := num - 2
	if count == deliteli {
		fmt.Println("Простое")
	} else {
		fmt.Println("Составное")
	}
}

func main() {
	num := 13

	isPrime := true
	for i := 2; i < num; i++ {
		// Это условие того что число составное так как поделилось нацело.
		if num%2 == 0 {
			isPrime = false
			break
		}
	}

	if isPrime {
		fmt.Println("Простое!")
	} else {
		fmt.Println("Составное")
	}
}
