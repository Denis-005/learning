package main

import "fmt"

func main() {
	a := []int{3, 4, 6, 8, 2, 9, 1, 5, 3, 1, 2}
	fmt.Println(len(a))

	for i := 0; i < len(a); i++ {
		if i%2 == 0 && i+1 < len(a) {
			a[i+1], a[i] = a[i], a[i+1]
		}
	}

	fmt.Println(a)
}
