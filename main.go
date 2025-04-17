package main

import (
	"fmt"
	"learning/internol/structs"
)

func main() {
	n := structs.Human{
		Name:   "Denis",
		Age:    25,
		Height: 175.0,
	}

	fmt.Println(n)

}
