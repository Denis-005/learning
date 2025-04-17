package main

import "fmt"

type Book struct {
	Title string
	Pages int
	// Author struct {
	// Name string
	// Age int
	// }

	Author Author
}

type Author struct {
	Name string
	Age  int
}

type Movie struct {
	Title string
	Year  int
	Genre string
}

type Student struct {
	Name  string
	Grade int
}

type House struct {
	Number int
	Date   int
}

type Task struct {
	Title string
	Done  bool
}

type User struct {
	Name string
	Age  int
}

type Product struct {
	Appellation string
	Price       int
}

func task1() {
	var b Book
	b.Title = "Denis"
	b.Pages = 250

	fmt.Println(b)

	//////////////////////////////////////////////////

	x := Book{Title: "РђР»РіРѕСЂРёС‚РјС‹", Pages: 250}
	fmt.Println(x)
	fmt.Println(x.Title, x.Pages)

	//////////////////////////////////////////////////

	var p Book
	fmt.Scan(&p.Title)
	fmt.Scan(&p.Pages)

	fmt.Printf("Title: %s, Pages: %d", p.Title, p.Pages)
}

func task2() {
	m := Movie{Title: "Golang", Year: 2015, Genre: "РЈР¶Р°СЃС‹"}
	fmt.Print(m)
}

func task3() {
	book := [3]Book{
		// {"Marvel", 2012},
		// {"Marvel", 2015},
		// {"Marvel", 2018},
	}

	fmt.Println(book)

	for i := 0; i < len(book); i++ {
		fmt.Println(book[i])
	}

	//var
}

func task4() {
	//size := 3
	//filter := make([]Movie, size)

	filter := []Movie{
		{"T", 25, "y"},
	}

	for _, v := range filter {
		fmt.Scan(&v.Title)
		fmt.Scan(&v.Year)
		fmt.Scan(&v.Genre)

		filter = append(filter, v)
	}

	fmt.Println(filter)

	size := 5
	h := make([]House, size)
	for i := range h {
		fmt.Scan(&h[i].Number, &h[i].Date)
	}

	for j := range size {
		fmt.Println(h[j])
	}

	var street []House
	for i := 0; i < 5; i++ {
		var number int
		var date int
		fmt.Scan(&number, &date)

		h := House{
			Number: number,
			Date:   date,
		}

		street = append(street, h)
	}

	for _, h := range street {
		fmt.Println(h)
	}

}

func task5() {

	p := Movie{}
	fmt.Scan(&p.Title, &p.Year, &p.Genre)

	filter := []Movie{}
	filter = append(filter, p)

	for _, m := range filter {
		fmt.Scan(&m.Title, &m.Year, &m.Genre)

		if m.Year > 2010 {
			filter = append(filter, m)
		}

	}

	fmt.Println(filter)

}

func task6() {
	students := []Student{
		{"Denis", 5},
		{"Bob", 4},
	}

	sum := 0
	for _, s := range students {
		sum += s.Grade
	}

	fmt.Println("Average:", float64(sum)/float64(len(students)))

}

func task7() {
	// b := Book{
	// Title: "Golang",
	// Pages: 310,
	// Author: struct {
	// Name string
	// Age int
	// }{
	// Name: "Alen",
	// Age: 45,
	// },
	// }

	b := Book{
		Title: "Golang",
		Pages: 310,
		Author: Author{
			Name: "Alen",
			Age:  45,
		},
	}

	fmt.Println(b)

}

func removeTaskByIndex(slice []Task, s int) []Task {
	return append(slice[:s], slice[s+1:]...)
}

func task8() {
	tasks := []Task{
		{Title: "Решить задачу", Done: true},
		{Title: "РќР°РїРёСЃР°С‚СЊ РєРѕРґ:", Done: false},
		{Title: "РЎРѕР·РґР°С‚СЊ СЃС‚СЂСѓРєС‚СЂСѓ Task:", Done: true},
		{Title: "РЎРѕР·РґР°С‚СЊ РјРµС‚РѕРґ СЃС‚СЂСѓРєС‚СѓСЂС‹:", Done: false},
	}

	for i, v := range tasks {
		if v.Done {
			tasks = removeTaskByIndex(tasks, i)
		}
	}

	fmt.Println(tasks)

}

func (u User) SayHello() {
	fmt.Println("РўРІРѕС‘ РёРјСЏ:", u.Name)
}

func (u User) IsAdult() bool {
	return u.Age >= 18
}

func (product Product) Print() {
	fmt.Printf("Наименование: %s, Цена: %d", product.Appellation, product.Price)
}

func (product Product) DiscountedPrice(percent int) float64 {
	return float64(product.Price) * (1 - float64(percent)/100)
}

func main() {
	//task1()
	//task2()
	//task3()
	//task4()
	//task5()
	//task7()
	task8()

	// u := User{Name: "Den"}
	// u.SayHello()

	// age := User{Age: 18}
	// res := age.IsAdult()
	// fmt.Println(res)

	// product := Product{Appellation: "РҐР»РµР±", Price: 45}
	// product.Print()

	// product := Product{Appellation: "РҐР»РµР±", Price: 45}
	// res := product.DiscountedPrice(product.Price)
	// fmt.Println(res)

}
