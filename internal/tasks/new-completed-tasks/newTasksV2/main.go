package main

import (
	"fmt"
	"strings"
)

func main() {
	// initializingStore()
	// initializingLibrary()
	// initializingStructWordCounter()
	// initializingShapes()
	// initializingStructTeacher()
	// initializingStructChatRoom()
}

/*

Задача 1. Магазин
Описание: Создай структуру Product с полями: Name string, Price float64. Создай структуру Store, которая хранит срез Products []Product. Добавь методы:
AddProduct(p Product) — добавляет товар в магазин
TotalValue() float64 — возвращает суммарную стоимость всех товаров
Пример:
store := Store{}
store.AddProduct(Product{"Book", 10})
store.AddProduct(Product{"Pen", 2})
fmt.Println(store.TotalValue()) // 12

*/

type Product struct {
	Name  string
	Prise float64
}

type Store struct {
	Product []Product
}

func (s *Store) AddProduct(p Product) {
	s.Product = append(s.Product, Product{
		Name:  p.Name,
		Prise: p.Prise,
	})
}

func (s Store) TotalValue() float64 {
	var totalSumProduct float64
	for _, val := range s.Product {
		totalSumProduct += val.Prise
	}

	return totalSumProduct
}

func initializingStore() {
	store := Store{}
	store.AddProduct(Product{Name: "Book", Prise: 10})
	store.AddProduct(Product{Name: "Pen", Prise: 2})

	fmt.Println(store.TotalValue())
}

/*

Задача 2. Библиотека
Описание: Создай структуру Book с полями: Title string и Author string. Создай структуру Library с картой Books map[string]Book, где ключ — название книги. Добавь методы:
AddBook(b Book) — добавляет книгу
FindBook(title string) (Book, bool) — ищет книгу по названию
Пример:
lib := Library{}
lib.AddBook(Book{"Go Basics", "Alice"})
b, ok := lib.FindBook("Go Basics")
fmt.Println(b.Author, ok) // Alice true

*/

type Book struct {
	Title  string
	Author string
}

type Library struct {
	Books map[string]Book
}

func (l *Library) AddBook(book Book) {
	l.Books[book.Author] = book
}

func (l Library) FindBook(title string) (Book, bool) {
	book := Book{}
	isFoundBook := false
	for key, val := range l.Books {
		if val.Title == title {
			book = Book{
				Title:  val.Title,
				Author: key,
			}

			isFoundBook = true
		}
	}

	return book, isFoundBook
}

func initializingLibrary() {
	lib := Library{Books: make(map[string]Book)}
	lib.AddBook(Book{Title: "Go Basics", Author: "Alice"})

	b, ok := lib.FindBook("Go Basics")
	fmt.Println(b.Author, ok)
}

/*

Задача 3. Счётчик слов
Описание: Создай структуру WordCounter с картой Counts map[string]int. Добавь метод AddText(text string), который считает слова в тексте и обновляет карту.
Пример:
wc := WordCounter{}
wc.AddText("go go gopher")
fmt.Println(wc.Counts["go"]) // 2
fmt.Println(wc.Counts["gopher"]) // 1

*/

type WordCounter struct {
	Counts map[string]int
}

func (wc *WordCounter) AddText(text string) {
	sliceStr := strings.Split(text, " ")
	for _, val := range sliceStr {
		wc.Counts[val]++
	}

}

func initializingStructWordCounter() {
	wc := WordCounter{Counts: make(map[string]int)}
	wc.AddText("go go gopher")

	fmt.Println(wc.Counts["go"])
	fmt.Println(wc.Counts["gopher"])
}

/*

Задача 4. Геометрия
Описание: Создай интерфейс Shape с методом Area() float64. Создай структуры Circle и Rectangle, которые реализуют этот интерфейс. Создай структуру Canvas с срезом Shapes []Shape. Добавь метод TotalArea() float64, который возвращает суммарную площадь всех фигур.
Пример:
canvas := Canvas{}
canvas.Shapes = append(canvas.Shapes, Circle{Radius:2})
canvas.Shapes = append(canvas.Shapes, Rectangle{Width:3, Height:4})
fmt.Println(canvas.TotalArea()) // 12 + 12.566 ~= 24.566

*/

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Canvas struct {
	Shapes []Shape
}

func (s Canvas) TotalArea() float64 {
	totlaSumArea := 0.0
	for _, val := range s.Shapes { // Хочу обратить внимание на данный метод так как впервые использую интерфейс в роли слайса. Можно будет обсудить.
		totlaSumArea += val.Area()
	}

	return totlaSumArea
}

func initializingShapes() {
	canvas := Canvas{}
	canvas.Shapes = append(canvas.Shapes, Circle{Radius: 2})
	canvas.Shapes = append(canvas.Shapes, Rectangle{Width: 3, Height: 4})

	fmt.Println(canvas.TotalArea())
}

/*

Задача 5. Учитель и студенты
Описание: Создай структуру Student с полями Name string и Grade float64. Создай структуру Teacher с полем Students []Student. Добавь методы:
AddStudent(s Student) — добавляет студента
AverageGrade() float64 — возвращает средний балл всех студентов
TopStudent() Student — возвращает студента с наивысшим баллом
Пример:
t := Teacher{}
t.AddStudent(Student{"Bob", 4.5})
t.AddStudent(Student{"Alice", 5.0})
fmt.Println(t.AverageGrade()) // 4.75
fmt.Println(t.TopStudent().Name) // Alice

*/

type Student struct {
	Name  string
	Grade float64
}

type Teacher struct {
	Students []Student
}

func (t *Teacher) AddStudent(student Student) {
	t.Students = append(t.Students, Student{
		Name:  student.Name,
		Grade: student.Grade,
	})
}

func (t Teacher) AverageGrade() float64 {
	countStudents := 0.0
	totalSumGradeStudents := 0.0

	for _, val := range t.Students {
		totalSumGradeStudents += val.Grade
		countStudents++
	}

	return totalSumGradeStudents / countStudents
}

func (t Teacher) TopStudent() Student {
	nameStudent := Student{}
	maxGradeStudent := 0.0

	for _, val := range t.Students {
		if val.Grade > maxGradeStudent {
			maxGradeStudent = val.Grade
			nameStudent = val // Возвращаем студента целиком.
		}
	}

	/*
		top := Student{}
		for _, val := range t.Students {  // Отправли код счату GPT на проверку и он предложил вот такой вариант работы метода TopStudent().
			if val.Grade > top.Grade {
				top = val
			}
		}

		return top

	*/

	return nameStudent
}

func initializingStructTeacher() {
	t := Teacher{}
	t.AddStudent(Student{Name: "Bob", Grade: 4.5})
	t.AddStudent(Student{Name: "Alice", Grade: 5.0})

	fmt.Println(t.AverageGrade())
	fmt.Println(t.TopStudent().Name)
}

/*

Задача 6. Чат
Описание: Создай структуру Message с полями User string и Text string. Создай структуру ChatRoom с срезом Messages []Message. Добавь методы:
SendMessage(user, text string) — добавляет новое сообщение
History() []Message — возвращает все сообщения
HistoryByUser(user string) []Message — возвращает все сообщения одного пользователя
Пример:
chat := ChatRoom{}
chat.SendMessage("Alice", "Hello")
chat.SendMessage("Bob", "Hi")
fmt.Println(len(chat.History())) // 2
fmt.Println(len(chat.HistoryByUser("Alice"))) // 1

*/

type Message struct {
	User, Text string
}

type ChatRoom struct {
	Messages []Message
}

func (c *ChatRoom) SendMessage(user, text string) {
	c.Messages = append(c.Messages, Message{
		User: user,
		Text: text,
	})
}

func (c ChatRoom) History() []Message {
	return c.Messages
}

func (c ChatRoom) HistoryByUser(user string) []Message {
	allUserMessages := []Message{}
	for _, val := range c.Messages {
		if val.User == user {
			allUserMessages = append(allUserMessages, val)
		}
	}

	return allUserMessages
}

func initializingStructChatRoom() {
	chat := ChatRoom{}
	chat.SendMessage("Alice", "Hello")
	chat.SendMessage("Bob", "Hi")

	fmt.Println(len(chat.History()))
	fmt.Println(len(chat.HistoryByUser("Alice")))

}
