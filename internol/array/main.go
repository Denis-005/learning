package main

import "fmt"

func main() {
	var size int
	fmt.Scan(&size)
	slice := make([]int, size)
	isFalse := false

	for i := 0; i < size; i++ {
		fmt.Scan(&slice[i])
	}

	for i := 1; i < len(slice)-1; i++ {
		if slice[i] > 0 && slice[i-1] > 0 || slice[i] > 0 && slice[i+1] > 0 {
			isFalse = true
		} else if slice[i] < 0 && slice[i-1] < 0 || slice[i] < 0 && slice[i+1] < 0 {
			isFalse = true
		} else {
			continue
		}
	}

	if isFalse {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}

}

// Циклический сдвиг вправо.

// Напишите программу, которая циклически сдвигает элементы массива вправо.
// Например, если элементы нумеруются, начиная с нуля, то 0-й элемент становится 1-м,
// 1-й становится 2-м, ..., последний становится 0-м, то есть массив [3,5,7,9]
// превращается в массив [9,3,5,7].

func test() {
	var size int
	fmt.Scan(&size)
	slice := make([]int, size)

	for i := 1; i < size; i++ {
		fmt.Scan(&slice[i])
	}

	for i := 0; i < len(slice); i++ {
		fmt.Scan(&slice[i])

		fmt.Print(slice[i], " ")
	}

}
