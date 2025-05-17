package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//task1()
	//task2()
	//task3()
	//task4()
	//task5()
	//task6()
	//task8()

}

/*

Частотный словарь (Легкий) Напиши функцию, которая принимает строку и возвращает мапу map[rune]int,
где ключ — символ, а значение — количество его вхождений в строке.
Функция: func CharFrequency(s string) map[rune]int Вход: "hello" Выход: map[rune]int{'h': 1, 'e': 1, 'l': 2, 'o': 1}

*/

func CharFrequency(s string) map[rune]int {
	m := make(map[rune]int)
	for i := 0; i < len(s); i++ {
		r := rune(s[i])
		m[r]++
	}

	return m
}

func task1() {
	var str string
	fmt.Scan(&str)

	res := CharFrequency(str)

	first := true
	for key, val := range res {
		if !first {
			fmt.Print(", ")
		}

		fmt.Printf("'%c': %d", key, val)
		first = false
	}

}

/*

Уникальные элементы (Легкий) Функция принимает слайс целых чисел и возвращает новый слайс,
содержащий только уникальные элементы в порядке их первого появления.
Функция: func UniqueElements(nums []int) []int Вход: []int{1, 2, 2, 3, 4, 3, 5} Выход: []int{1, 2, 3, 4, 5}

*/

func UniqueElements(nums []int) []int {
	arr := []int{}
	m := make(map[int]bool)

	for _, val := range nums {
		if _, ok := m[val]; !ok {
			arr = append(arr, val)
		}

		m[val] = true
	}

	return arr

}

func task2() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	sliceStr := strings.Fields(input)
	sliceInt := []int{}

	for i := range sliceStr {
		str := sliceStr[i]

		for j := range str {
			temp, err := strconv.Atoi(string(str[j]))

			if err == nil {
				sliceInt = append(sliceInt, temp)
			}

		}

	}

	res := UniqueElements(sliceInt)
	for i := range res {
		fmt.Print(res[i])

		if i < len(res)-1 {
			fmt.Print(", ")
		}
	}

}

/*

Инвертирование мапы (Легкий–Средний) Инвертируй мапу вида map[string]string (пользователь → город) в map[string][]string (город → список пользователей).
Функция: func InvertUserCityMap(users map[string]string) map[string][]string Вход: map[string]string{"Alice": "Paris", "Bob": "London", "Charlie": "Paris"}
Выход: map[string][]string{"Paris": {"Alice", "Charlie"}, "London": {"Bob"}}

*/

func InvertUserCityMap(users map[string]string) map[string][]string {
	m := make(map[string][]string)
	for name, city := range users {
		m[city] = append(m[city], name)
	}

	return m
}

func task3() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	slice := strings.Fields(input)
	m := make(map[string]string)

	var name string
	for i, val := range slice {
		if i%2 == 0 {
			name = val
		} else {
			m[name] = val
		}
	}

	res := InvertUserCityMap(m)
	fmt.Println(res)
}

/*

Проверка анаграмм (Средний) Проверь, являются ли две строки анаграммами (содержат одинаковые символы в любом порядке, без учёта регистра и пробелов).
Функция: func AreAnagrams(a, b string) bool Вход: "Listen", "Silent" Выход: true

*/

func fillingMap(s string) map[rune]int {
	m := make(map[rune]int)
	for _, val := range s {
		m[val]++
	}

	return m

}

func AreAnagrams(a, b string) bool {
	m1 := fillingMap(a)
	m2 := fillingMap(b)

	for key := range m1 {
		if _, ok := m2[key]; !ok {
			return false
		}
	}

	return true
}

func task4() {
	var str1, str2 string
	fmt.Scan(&str1, &str2)

	str1 = strings.ToLower(strings.ReplaceAll(str1, " ", ""))
	str2 = strings.ToLower(strings.ReplaceAll(str2, " ", ""))

	if len(str1) != len(str2) {
		fmt.Println("Разная длинна строк!")
		return
	}

	res := AreAnagrams(str1, str2)
	fmt.Println(res)
}

/*

Подсчет слов (Средний) Напиши функцию, которая считает, сколько раз каждое слово встречается в тексте. Слова разделяются пробелами и знаками препинания.
Функция: func WordCount(text string) map[string]int Вход: "Hello, world! Hello world." Выход: map[string]int{"hello": 2, "world": 2}

*/

func WordCount(text string) map[string]int {
	m := make(map[string]int)

	slice := strings.Split(text, " ")
	for i := range slice {
		tempStr := strings.ToLower(slice[i])

		slice[i] = tempStr
	}

	fmt.Println(slice)

	for _, val := range slice {
		if _, ok := m[val]; ok {
			m[val]++
		}
	}

	return m
}

func task5() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	res := WordCount(input)
	fmt.Println(res)
}

/*

Проверка на повторяющиеся строки (Средний) Напиши функцию, которая проверяет, есть ли в списке строк дубликаты.
Верни true, если есть хотя бы одна строка, встречающаяся более одного раза.
Функция: func HasDuplicates(words []string) bool Вход: []string{"apple", "banana", "orange", "apple"} Выход: true Вход: []string{"go", "rust", "python"} Выход: false

*/

func HasDuplicates(words []string) bool {
	m := make(map[string]bool)
	for _, val := range words {
		if _, ok := m[val]; ok {
			return true
		}

		m[val] = true
	}

	return false

}

func task6() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	slice := strings.Split(input, " ")

	res := HasDuplicates(slice)
	fmt.Println(res)
}

/*

Пересечение двух мап (Сложный) Верни мапу с ключами, которые есть в обеих, и значениями — суммой значений из обеих мап.
Функция: func IntersectMaps(a, b map[string]int) map[string]int Вход:
a := map[string]int{"apple": 2, "banana": 3}
b := map[string]int{"banana": 1, "apple": 1, "cherry": 5}
Выход: map[string]int{"apple": 3, "banana": 4}

*/

func IntersectMaps(a, b map[string]int) map[string]int {
	m := make(map[string]int)
	// for key, val := range a {
	// 	if _, ok := b[key]; ok {
	// 		m[key] += val
	// 	}
	// }

	// for key, val := range b {
	// 	if _, ok := a[key]; ok {
	// 		m[key] += val
	// 	}
	// }

	for key, val := range a {
		if valB, ok := b[key]; ok {
			m[key] = val + valB
		}
	}

	return m

}

func task8() {
	a := map[string]int{"apple": 2, "banana": 3}
	b := map[string]int{"banana": 1, "apple": 1, "cherry": 5}

	res := IntersectMaps(a, b)
	fmt.Println(res)
}
