package tasks

import "fmt"

// 1. Вывести числа от 1 до 10.
// Напиши программу, которая выводит числа от 1 до 10, используя цикл for.

func printSequenceOfNum() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}
