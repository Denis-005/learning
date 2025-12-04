package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//countElements()
	//initMap()
	//createMaps()

}

/*
Задача 1: Подсчёт частоты
Условие:
Реализуйте функцию WordCount(s []string) map[string]int, которая возвращает мапу,
в которой ключ — это слово, а значение — сколько раз оно встречается в слайсе.

Input:  ["go", "java", "go", "python"]
Output: map[go:2 java:1 python:1]
*/

func WordCount(s []string) map[string]int {
	dictionary := make(map[string]int)
	for _, word := range s {
		dictionary[word]++
	}

	return dictionary
}

func countElements() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	words := strings.Fields(scanner.Text())

	res := WordCount(words)
	fmt.Println(res)
}

/*
Задача 2: Инвертирование мапы
Условие:
Напишите функцию InvertMap(m map[string]string) map[string][]string, которая инвертирует мапу. Если значение повторяется, собирайте ключи в слайс.

Input:  map[a:x b:y c:x]
Output: map[x:[a c] y:[b]]
*/

func InvertMap(d map[string]string) map[string][]string {
	dataMap := make(map[string][]string)
	for k, v := range d {
		dataMap[v] = append(dataMap[v], k)
	}

	return dataMap
}

func initMap() {
	dictionary := make(map[string]string)
	for i := 0; i < 3; i++ {
		var key, val string
		fmt.Scan(&key, &val)

		dictionary[key] = val
	}

	res := InvertMap(dictionary)
	fmt.Println(res)
}

/*
Задача 3: Находятся ли мапы равными
Условие:
Реализуйте функцию MapsEqual(a, b map[string]int) bool, которая проверяет, равны ли две мапы (одинаковый набор ключей и соответствующие значения).

Input:  a = map[x:1 y:2], b = map[y:2 x:1]
Output: true

Input:  a = map[a:1], b = map[a:2]
Output: false

Input:  a = map[], b = map[]
Output: true
*/

func MapsEqual(a, b map[string]int) bool {
	isMatch := true
	for k, v := range a {
		if val, ok := b[k]; !ok || val != v {
			isMatch = false
		}
	}

	return isMatch
}

func createMaps() {
	firstMap := make(map[string]int)
	secondMap := make(map[string]int)

	for i := 0; i < 10; i++ {
		var key string
		fmt.Scan(&key)

		if key == "exit" {
			break
		}

		firstMap[key]++
	}

	for i := 0; i < 10; i++ {
		var key string
		fmt.Scan(&key)

		if key == "exit" {
			break
		}

		secondMap[key]++
	}

	if len(firstMap) != len(secondMap) {
		fmt.Println(false)
		return
	}

	res := MapsEqual(firstMap, secondMap)
	fmt.Println(res)

}
