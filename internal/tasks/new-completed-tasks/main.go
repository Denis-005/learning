package main

import (
	"fmt"
	"strings"
)

func main() {
	SumNums(2, 3)
	MaxNum(5, 10)
	IsEven(4)
	SumArray([...]int{1, 2, 3})
	ReverseStr("Go")
	Contains([]int{1, 2, 3}, 2)
	WordCount("go go gopher")
	initializingStructBook()
	initializingStructRectangle()
	Factorial(0)
	
}

/*

Задача 1. Сумма чисел

Описание:
Напиши функцию Sum, которая принимает два числа a и b и возвращает их сумму.

func Sum(a int, b int) int

Пример:

Sum(2, 3) // 5
Sum(-1, 4) // 3

*/

func SumNums(firstNum, secondNum int) {
	fmt.Println(firstNum + secondNum)
}

/*

Задача 2. Максимальное число

Описание:
Напиши функцию Max, которая возвращает большее из двух чисел.

func Max(a int, b int) int

Пример:

Max(5, 10) // 10
Max(7, 3)  // 7

*/

func MaxNum(firstNum, secondNum int) {
	fmt.Println(max(firstNum, secondNum))

}

/*

Задача 3. Чётное или нечётное

Описание:
Напиши функцию IsEven, которая возвращает true, если число чётное, и false иначе.

func IsEven(n int) bool

Пример:

IsEven(4) // true
IsEven(7) // false

*/

func IsEven(num int) {
	isEven := false
	if num%2 == 0 {
		isEven = true
	}

	fmt.Println(isEven)
}

/*

Задача 4. Сумма массива

Описание:
Напиши функцию SumArray, которая принимает массив чисел и возвращает их сумму.

func SumArray(arr []int) int

Пример:

SumArray([]int{1,2,3}) // 6
SumArray([]int{}) // 0

*/

func SumArray(arr [3]int) {
	sum := 0
	for i := range arr {
		sum += arr[i]
	}

	fmt.Println(sum)
}

/*

Задача 5. Реверс строки

Описание:
Напиши функцию Reverse, которая переворачивает строку.

func Reverse(s string) string

Пример:

Reverse("hello") // "olleh"
Reverse("Go")    // "oG"

*/

func ReverseStr(str string) {
	newStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		newStr += string(str[i])
	}

	fmt.Println(newStr)
}

/*

Задача 6. Поиск элемента

Описание:
Напиши функцию Contains, которая проверяет, есть ли элемент в срезе.

func Contains(slice []int, x int) bool

Пример:

Contains([]int{1,2,3}, 2) // true
Contains([]int{1,2,3}, 4) // false

*/

func Contains(arr []int, num int) {
	isEqually := false
	for _, val := range arr {
		if val == num {
			isEqually = true
		}
	}

	fmt.Println(isEqually)
}

/*

Задача 7. Подсчёт слов

Описание:
Напиши функцию WordCount, которая принимает строку и возвращает карту, где ключ — слово, а значение — количество повторений.

func WordCount(s string) map[string]int

Пример:

WordCount("go go gopher") // map[string]int{"go":2, "gopher":1

*/

func WordCount(str string) {
	sliceStr := strings.Split(str, " ")

	mapStr := make(map[string]int)
	for i := range sliceStr {
		mapStr[sliceStr[i]]++
	}

	fmt.Println(mapStr)
}

/*

Задача 8. Структура “Книга”

Описание:
Создай структуру Book с полями Title string, Author string, Pages int.
Напиши функцию NewBook, которая создаёт книгу и возвращает её.

type Book struct {
    Title string
    Author string
    Pages int
}

func NewBook(title, author string, pages int) Book

Пример:

b := NewBook("Go Programming", "John", 250)
fmt.Println(b.Title) // "Go Programming"

*/

type Book struct {
	Title  string
	Author string
	Pages  int
}

func NewBook(title, author string, pagew int) Book {
	book := Book{
		Title:  title,
		Author: author,
		Pages:  pagew,
	}

	return book
}

func initializingStructBook() {
	b := NewBook("Go Programming", "John", 250)
	fmt.Println(b.Title)
}

/*

Задача 9. Интерфейс “Shape”

Описание:
Создай интерфейс Shape с методом Area() float64.
Создай структуру Rectangle с полями Width и Height и реализуй метод Area.

type Shape interface {
    Area() float64
}

type Rectangle struct {
    Width, Height float64
}

Пример:

r := Rectangle{3, 4}
fmt.Println(r.Area()) // 12

*/

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func initializingStructRectangle() {
	r := Rectangle{Width: 3, Height: 4}
	fmt.Println(r.Area())
}

func Shapes(s Shape) float64 {
	return s.Area()
}

/*

Задача 10. Факториал

Описание:
Напиши функцию Factorial, которая возвращает факториал числа n.

func Factorial(n int) int

Пример:

Factorial(5) // 120
Factorial(0) // 1

*/

func Factorial(n int) {
	fact := 1

	if n > 0 {
		for i := 1; i <= n; i++ {
			fact *= i
		}
	}

	fmt.Println(fact)
}
