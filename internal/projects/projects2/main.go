// 1 - вариант: Перезапись полей через копию структуры.

package main

import "fmt"

/*
   Мини-проект №3 — CLI-библиотека книг
   Ты должен реализовать программу, которая позволяет управлять коллекцией книг.
   У книги есть:
   Title — название,
   Author — автор,
   Year — год издания.

   1. Добавить книгу
   2. Найти книгу по названию
   3. Удалить книгу
   4. Показать все книги
   5. Обновить данные книги (название, автора или год)
   0. Выход
*/

type Book struct {
	Title  string
	Author string
	Year   int
}

func removeNum(slice []int, num int) []int {
	for i, val := range slice {
		if val == num {
			return append(slice[:i], slice[i+1:]...)
		}
	}

	return slice
}

func main() {
	var (
		numBook int
		title   string
		author  string
		year    int
	)

	numBooks := []int{}
	books := make(map[int]Book)
	for {
		fmt.Println("Введите цифру действия из списка!")
		fmt.Println("1. Добавить книгу")
		fmt.Println("2. Найти книгу по названию")
		fmt.Println("3. Удалить книгу")
		fmt.Println("4. Показать все книги")
		fmt.Println("5. Обновить данные книги (название, автора или год)")
		fmt.Println("0. Выход")

		var num int
		fmt.Scan(&num)

		switch num {
		case 0:
			return
		case 1:
			fmt.Println("Введите номер книги!")
			fmt.Scan(&numBook)

			if _, ok := books[numBook]; ok {
				fmt.Println("Книга под таким номером уже существует!")
				continue
			}

			numBooks = append(numBooks, numBook)

			fmt.Println("Введите название книги!")
			fmt.Scan(&title)

			fmt.Println("Введите автора книги!")
			fmt.Scan(&author)

			fmt.Println("Введите год издания книги!")
			fmt.Scan(&year)

			books[numBook] = Book{
				Title:  title,
				Author: author,
				Year:   year,
			}
		case 2:
			fmt.Println("Введите название книги!")
			fmt.Scan(&title)

			for _, titleBook := range books {
				if titleBook.Title == title {
					fmt.Printf("Книга найдена: %s\n", title)
				}
			}
		case 3:
			fmt.Println("Введите номер книги!")
			fmt.Scan(&numBook)

			delete(books, numBook)
			numBooks = removeNum(numBooks, numBook)

			fmt.Println("Выполнено!")
		case 4:
			fmt.Println("Список всех книг!")
			for _, val := range numBooks {
				fmt.Printf("%d: %s %s %d\n", val, books[val].Title, books[val].Author, books[val].Year)
			}
		case 5:
			fmt.Println("Введите номер книги")
			fmt.Scan(&numBook)

			if _, ok := books[numBook]; ok {
				fmt.Println("Введите номер из списка того что хотите обновить!")
				fmt.Println("1 - Название")
				fmt.Println("2 - Автора")
				fmt.Println("3 - Год издания")
				fmt.Println("4 - Всю книгу")

				var num int
				fmt.Scan(&num)

				switch num {
				case 1:
					fmt.Println("Введите новое название книги!")
					fmt.Scan(&title)

					book := books[numBook]
					book.Title = title
					books[numBook] = book
				case 2:
					fmt.Println("Введите нового автора книги!")
					fmt.Scan(&author)

					book := books[numBook]
					book.Author = author
					books[numBook] = book
				case 3:
					fmt.Println("Введите новый год издания книги!")
					fmt.Scan(&year)

					book := books[numBook]
					book.Year = year
					books[numBook] = book
				case 4:
					fmt.Println("Обновите всю книгу: название, автора, год издания")
					fmt.Scan(&title, &author, &year)

					books[numBook] = Book{
						Title:  title,
						Author: author,
						Year:   year,
					}
				default:
					fmt.Println("Неверное значение из списка операции. Повторите попытку!")
				}
			} else {
				fmt.Println("Книги под таким номером нет!")
			}
		default:
			fmt.Println("Неверное значение из списка операции. Повторите попытку!")
		}

	}
}

// 2 вариант: Работаем с оригинальной структурой через указатель.

/*
package main

import "fmt"

//Мини-проект №3 — CLI-библиотека книг
//Ты должен реализовать программу, которая позволяет управлять коллекцией книг.
//У книги есть:
//Title — название,
//Author — автор,
//Year — год издания.

//1. Добавить книгу
//2. Найти книгу по названию
//3. Удалить книгу
//4. Показать все книги
//5. Обновить данные книги (название, автора или год)
//0. Выход

type Book struct {
	Title  string
	Author string
	Year   int
}

func removeNum(slice []int, num int) []int {
	for i, val := range slice {
		if val == num {
			return append(slice[:i], slice[i+1:]...)
		}
	}

	return slice
}

func main() {
	var (
		numBook int
		title   string
		author  string
		year    int
	)

	numBooks := []int{}
	books := make(map[int]*Book)
	for {
		fmt.Println("Введите цифру действия из списка!")
		fmt.Println("1. Добавить книгу")
		fmt.Println("2. Найти книгу по названию")
		fmt.Println("3. Удалить книгу")
		fmt.Println("4. Показать все книги")
		fmt.Println("5. Обновить данные книги (название, автора или год)")
		fmt.Println("0. Выход")

		var num int
		fmt.Scan(&num)

		switch num {
		case 0:
			return
		case 1:
			fmt.Println("Введите номер книги!")
			fmt.Scan(&numBook)

			if _, ok := books[numBook]; ok {
				fmt.Println("Книга под таким номером уже существует!")
				continue
			}

			numBooks = append(numBooks, numBook)

			fmt.Println("Введите название книги!")
			fmt.Scan(&title)

			fmt.Println("Введите автора книги!")
			fmt.Scan(&author)

			fmt.Println("Введите год издания книги!")
			fmt.Scan(&year)

			books[numBook] = &Book{
				Title:  title,
				Author: author,
				Year:   year,
			}
		case 2:
			fmt.Println("Введите название книги!")
			fmt.Scan(&title)

			for _, titleBook := range books {
				if titleBook.Title == title {
					fmt.Printf("Книга найдена: %s\n", title)
				}
			}
		case 3:
			fmt.Println("Введите номер книги!")
			fmt.Scan(&numBook)

			delete(books, numBook)
			numBooks = removeNum(numBooks, numBook)

			fmt.Println("Выполнено!")
		case 4:
			if len(numBooks) > 0 {
				fmt.Println("Список всех книг!")
				for _, val := range numBooks {
					fmt.Printf("%d: %s %s %d\n", val, books[val].Title, books[val].Author, books[val].Year)
				}
			} else {
				fmt.Println("Список пуст!")
			}
		case 5:
			fmt.Println("Введите номер книги")
			fmt.Scan(&numBook)

			if _, ok := books[numBook]; ok {
				fmt.Println("Введите номер из списка того что хотите обновить!")
				fmt.Println("1 - Название")
				fmt.Println("2 - Автора")
				fmt.Println("3 - Год издания")
				fmt.Println("4 - Всю книгу")

				var num int
				fmt.Scan(&num)

				switch num {
				case 1:
					fmt.Println("Введите новое название книги!")
					fmt.Scan(&title)

					books[numBook].Title = title
				case 2:
					fmt.Println("Введите нового автора книги!")
					fmt.Scan(&author)

					books[numBook].Author = author
				case 3:
					fmt.Println("Введите новый год издания книги!")
					fmt.Scan(&year)

					books[numBook].Year = year
				case 4:
					fmt.Println("Обновите всю книгу: название, автора, год издания")
					fmt.Scan(&title, &author, &year)

					books[numBook] = &Book{
						Title:  title,
						Author: author,
						Year:   year,
					}
				default:
					fmt.Println("Неверное значение из списка операции. Повторите попытку!")
				}
			} else {
				fmt.Println("Книги под таким номером нет!")
			}
		default:
			fmt.Println("Неверное значение из списка операции. Повторите попытку!")
		}

	}
}
*/
