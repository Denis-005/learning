package tasks

import "fmt"

// Замечательное число.
// Дано натуральное число N .Определите является ли данное число замечательным.
// Число называется замечательным, если оно делится на сумму своих цифр без остатка.
func wonderfulNum() {
	var n int
	fmt.Scan(&n)
	t := n
	d := 0
	sum := 0
	isNum := true

	i := 0
	for i < n {
		d = n % 10
		sum += d
		n /= 10
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

// Обратное число.
// Дано натуральное число N. Переставьте цифры числа в обратном порядке.
func reverseNum() {
	var n int
	fmt.Scan(&n)
	d := 0
	s := 0

	i := 1
	for i <= n {
		d = n % 10
		s = d
		if s <= n {
			fmt.Print(s)
		}
		n /= 10
	}
}
