package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
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
	//task7()
	//task8()
	//task9()
	//task10()
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

	isCommaNeeded := false
	for key, val := range res {
		if isCommaNeeded {
			fmt.Print(", ")
		}

		fmt.Printf("'%c': %d", key, val)
		isCommaNeeded = true
	}
}

/*

Уникальные элементы (Легкий) Функция принимает слайс целых чисел и возвращает новый слайс,
содержащий только уникальные элементы в порядке их первого появления.
Функция: func UniqueElements(nums []int) []int Вход: []int{1, 2, 2, 3, 4, 3, 5} Выход: []int{1, 2, 3, 4, 5}

*/

func getUniqueNums(nums []int) []int {
	res := []int{}
	uniqueNums := make(map[int]struct{})

	for _, val := range nums {
		if _, ok := uniqueNums[val]; !ok {
			res = append(res, val)
		}

		uniqueNums[val] = struct{}{}
	}

	// m := make(map[int]bool)

	// for _, val := range nums {
	// 	if _, ok := m[val]; !ok {
	// 		res = append(res, val)
	// 	}

	// 	m[val] = true
	// }

	return res

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
			if err != nil {
				log.Fatal(err)
			}

			sliceInt = append(sliceInt, temp)

		}

	}

	res := getUniqueNums(sliceInt)
	for i := range res {
		fmt.Print(res[i])

		if i < len(res)-1 {
			fmt.Print(", ")
		}
	}
}

func task2_V2() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(os.Stdin)
	numbers := make([]int, n)

	for i := range n {
		_, err := fmt.Fscan(reader, &numbers[i])
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println(numbers)

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

	if len(a) != len(b) {
		return false
	}

	for key, val := range m1 {
		if _, ok := m2[key]; !ok {
			if val != m2[key] {
				return false
			}
		}
	}

	return true
}

func task4() {
	var str1, str2 string
	fmt.Scan(&str1, &str2)

	str1 = strings.ToLower(strings.ReplaceAll(str1, " ", ""))
	str2 = strings.ToLower(strings.ReplaceAll(str2, " ", ""))

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
	var re = regexp.MustCompile(`[[:punct:]]`)

	for i := range slice {
		tempStr := strings.ToLower(slice[i])
		tempStr = re.ReplaceAllString(tempStr, "")

		slice[i] = tempStr
	}

	for _, val := range slice {
		m[val]++
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
	seenWords := make(map[string]struct{})
	for _, val := range words {
		if _, ok := seenWords[val]; ok {
			return true
		}

		seenWords[val] = struct{}{}
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

Обратный словарь с приоритетом (Средний–Сложный) Построй мапу map[int][]string по исходной map[string]int, группируя задачи по приоритету.
Внутри каждой группы отсортируй задачи по алфавиту. Функция:
func GroupTasksByPriority(tasks map[string]int) map[int][]string Вход: map[string]int{"TaskA": 2, "TaskB": 1, "TaskC": 2}
Выход: map[int][]string{1: {"TaskB"}, 2: {"TaskA", "TaskC"}}

*/

func GroupTasksByPriority(tasks map[string]int) map[int][]string {
	m := make(map[int][]string)
	for key, val := range tasks {
		m[val] = append(m[val], key)
		sort.Strings(m[val])
	}

	return m
}

func task7() {
	m := map[string]int{"TaskA": 2, "TaskB": 1, "TaskC": 2}

	res := GroupTasksByPriority(m)
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

/*

Рейтинг фильмов (Сложный) Рассчитай среднюю оценку каждого фильма на основе мапы оценок.
Функция: func AverageRatings(ratings map[string][]int) map[string]float64 Вход: map[string][]int{"Inception":
{5, 4, 5}, "Avatar": {4, 3}} Выход: map[string]float64{"Inception": 4.67, "Avatar": 3.5}

*/

func AverageRatings(ratings map[string][]int) map[string]float64 {
	m := make(map[string]float64)

	for key, slice := range ratings {
		var sum float64

		for i := range slice {
			sum += float64(slice[i])
		}

		m[key] = sum / float64(len(slice))
	}

	return m
}

func task9() {
	m := map[string][]int{"Inception": {5, 4, 5}, "Avatar": {4, 3}}

	isValueFirst := false
	res := AverageRatings(m)
	for key, val := range res {
		if isValueFirst {
			fmt.Print(", ")
		}

		fmt.Printf("%q: %.2f", key, val)
		isValueFirst = true
	}

}

/*

Самое частое слово (Сложный) Верни слово, которое чаще всего встречается в строке.
Игнорируй регистр. В случае равенства — верни любое из них. Функция: func MostFrequentWord(text string) string
Вход: "The sun is the sun and the moon is the moon" Выход: "the"

*/

func MostFrequentWord(text string) string {
	slice := strings.Split(text, " ")

	 resStr := ""
	 maxVal := 0
	m := make(map[string]int)
	for _, val := range slice {
		m[val]++

		for key, val := range m {
			if val > maxVal {
				resStr = key
				maxVal = val
			}
		}
	}

	return resStr
}

func task10() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	res := MostFrequentWord(strings.ToLower(input))
	fmt.Println(res)
}
