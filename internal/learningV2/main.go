package main

import "fmt"

func main() {
	var first, second int
	fmt.Scan(&first, &second)

	doubleNums := make(map[int]struct{})
	for first != 0 {
		digit := first % 10

		if _, ok := doubleNums[digit]; !ok {
			doubleNums[digit] = struct{}{}
		}
	}

	for second != 0 {
		digit := second % 10

		if _, ok := doubleNums[digit]; ok {
			fmt.Print(digit, " ")
		}
	}
}
