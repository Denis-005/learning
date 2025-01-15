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
