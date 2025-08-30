package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	initializationArrays()
	removingDuplicateNum()
}

/*

Поиск пересечения массивов
Напишите программу, которая находит пересечение двух массивов чисел и выводит только уникальные числа, входящие в оба массива.

*/

func conversionSlice(str []string) []int {
	slice := []int{}
	for i := range str {
		num, _ := strconv.Atoi(str[i])
		slice = append(slice, num)

	}

	return slice
}

func initializationArrays() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	i := scanner.Text()

	tempSliceStr1 := strings.Split(i, " ")
	firstSliceInt := conversionSlice(tempSliceStr1)

	scanner.Scan()
	n := scanner.Text()

	tempSliceStr2 := strings.Split(n, " ")
	secondSliceInt := conversionSlice(tempSliceStr2)

	m := make(map[int]bool)
	for _, v := range firstSliceInt {
		m[v] = true
	}

	res := []int{}
	for _, v := range secondSliceInt {
		if m[v] {
			res = append(res, v)
		}
	}

	for i := range res {
		fmt.Print(res[i], " ")
	}
}

/*

Удаление дубликатов чисел
Напишите программу, которая удаляет все дублирующиеся числа из массива и выводит только уникальные числа в порядке возрастания.

*/

func removingDuplicateNum() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	tempSlice := strings.Split(input, " ")

	slice := []int{}
	for i := range tempSlice {
		num, _ := strconv.Atoi(tempSlice[i])
		slice = append(slice, num)
	}

	m := make(map[int]struct{})
	for _, v := range slice {
		m[v] = struct{}{}
	}

	newSlice := []int{}
	for i := range m {
		newSlice = append(newSlice, i)
	}

	sort.Ints(newSlice)

	for i, v := range newSlice {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
}
