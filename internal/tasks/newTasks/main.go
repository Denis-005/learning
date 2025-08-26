package main

import "fmt"

func main() {
	// initializingStructCounter()
	// initializingStructBankAccount()
	// initializingStructs()
	// initializingStructStudents()
	// initializingStructsV2()
	// initializingToDoProject()

}

/*

Задача 1. Счётчик посещений
Описание: Создай структуру Counter, которая хранит количество. Добавь методы:
Increment() — увеличивает счётчик на 1
Value() int — возвращает текущее значение
Пример использования:
c := Counter{}
c.Increment()
c.Increment()
fmt.Println(c.Value()) // 2

*/

type Counter struct {
	Count int
}

func (c *Counter) Increment() {
	c.Count++
}

func (c Counter) Value() int {
	return c.Count
}

func initializingStructCounter() {
	c := Counter{}
	c.Increment()
	c.Increment()

	fmt.Println(c.Value())
}

/*

Задача 2. Банковский счёт
Описание: Создай структуру BankAccount с полями: Owner string и Balance float64. Добавь методы:
Deposit(amount float64) — добавляет деньги на счёт
Withdraw(amount float64) error — снимает деньги, если хватает, иначе возвращает ошибку
Пример:
a := BankAccount{"Alice", 100.0}
a.Deposit(50)
err := a.Withdraw(120)
fmt.Println(a.Balance) // 30
fmt.Println(err)       // nil

*/

type BankAccount struct {
	Owner   string
	Balance float64
}

func (a *BankAccount) Deposit(amount float64) {
	a.Balance += amount
}

func (a *BankAccount) Withdraw(amount float64) error {
	if a.Balance-amount >= 0 {
		a.Balance -= amount
		return nil
	} else {
		return fmt.Errorf("Не хватает денежных средств для снятия со счёта. Выберите не больше данной: %.2f денежной суммы!", a.Balance)
	}
}

func initializingStructBankAccount() {
	a := BankAccount{Owner: "Alice", Balance: 100.0}
	a.Deposit(20)
	err := a.Withdraw(125)

	fmt.Println(a.Balance)
	fmt.Println(err)
}

/*

Задача 3. Животные
Описание: Создай интерфейс Animal с методом Speak() string. Создай структуры Dog и Cat, реализующие интерфейс.
Пример:
func MakeSpeak(a Animal) {
    fmt.Println(a.Speak())
}

d := Dog{}
c := Cat{}
MakeSpeak(d) // "Woof!"
MakeSpeak(c) // "Meow!"

*/

type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow"
}

func initializingStructs() {
	d := Dog{}
	c := Cat{}

	fmt.Println(MakeSpeak(d))
	fmt.Println(MakeSpeak(c))
}

func MakeSpeak(a Animal) string {
	return a.Speak()
}

/*

Задача 4. Список студентов
Описание: Создай структуру Student с полями Name string и Grades []int. Напиши метод Average() float64, который возвращает среднюю оценку студента.
Пример:
s := Student{"Bob", []int{5, 4, 3, 5}}
fmt.Println(s.Average()) // 4.25

*/

type Student struct {
	Name   string
	Grades []int
}

func (s *Student) Average() float64 {
	count := 0.0
	totalSumGrades := 0.0
	for _, val := range s.Grades {
		totalSumGrades += float64(val)
		count++
	}

	return totalSumGrades / count
}

func initializingStructStudents() {
	s := Student{Name: "Bob", Grades: []int{5, 4, 3, 5}}
	fmt.Println(s.Average())

}

/*

Задача 5. Прямоугольники
Описание: Создай структуру Rectangle с полями Width и Height. Создай структуру Square, которая встраивает Rectangle. Напиши метод Area() float64 для обоих типов.
Пример:
r := Rectangle{3,4}
s := Square{Rectangle{5,5}}
fmt.Println(r.Area()) // 12
fmt.Println(s.Area()) // 25

*/

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Square struct {
	Rectangle
}

func (s Square) Area() float64 { // Данный метод явно показывает что когда в структуру встраивается другая структура, поля и методы встроенной структуры становятся напрямую доступными через новую структуру. Верно? И я так же понимаю что данный метод необязателен или же нет?
	return s.Width * s.Height
}

func initializingStructsV2() {
	r := Rectangle{3, 4}
	s := Square{Rectangle{Width: 5, Height: 5}}

	fmt.Println(r.Area())
	fmt.Println(s.Area())
}

/*

Задача 6. Список задач (ToDo)
Описание: Создай структуру Task с полями Title string и Done bool. Создай структуру ToDoList, которая хранит срез Tasks []Task. Добавь методы:
Add(title string) — добавляет новую задачу
Complete(index int) — отмечает задачу как выполненную
Pending() []Task — возвращает только невыполненные задачи
Пример:
list := ToDoList{}
list.Add("Учить Go")
list.Add("Прогуляться")
list.Complete(0)
fmt.Println(list.Pending()) // [{"Прогуляться", false}]

*/

type Task struct {
	Title string
	Done  bool
}

type ToDoList struct {
	Tasks []Task
}

func (l *ToDoList) Add(title string) {
	l.Tasks = append(l.Tasks, Task{
		Title: title,
		Done:  false,
	})
}

func (l *ToDoList) Complete(index int) {

	l.Tasks[index].Done = true

	// for i := range l.Tasks {  // Изначально решил снова через цикл (перебирая каждую строчку) хотя можно было просто указать на индекс (мне ещё практиковаться и практиковаться).
	// 	if i == index {
	// 		l.Tasks[i].Done = true
	// 	}
	// }
}

func (l ToDoList) Pending() []Task {
	sliceTasts := []Task{}
	for _, val := range l.Tasks {
		if !val.Done {
			sliceTasts = append(sliceTasts, val)
		}
	}

	return sliceTasts
}

func initializingToDoProject() {
	list := ToDoList{}
	list.Add("Учить Go")
	list.Add("Прогуляться")
	list.Complete(0)
	fmt.Println(list.Pending())
}
