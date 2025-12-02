package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
Следующий мини-проект №2 — CLI-словарь. Необходимо запросить слово и его определение.
1 - добавить слово
2 - найти слово
3 - удалить слово
4 - показать все слова
0 - выход
*/

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var word, text string
	words := make(map[string]string)
	for {
		fmt.Println("Введите цифру действия!")
		fmt.Println("1 - добавить слово")
		fmt.Println("2 - найти слово")
		fmt.Println("3 - удалить слово")
		fmt.Println("4 - показать все слова")
		fmt.Println("0 - выход")

		var num int
		fmt.Scanln(&num)

		if num >= 1 && num <= 3 {
			fmt.Println("Введите слово!")
			scanner.Scan()
			word = scanner.Text()

			if num == 1 {
				fmt.Println("Введите определение слова!")
				scanner.Scan()
				text = scanner.Text()
			}
		}

		switch num {
		case 0:
			return
		case 1:
			words[word] = text
			fmt.Println("Выполнено.")
		case 2:
			if val, ok := words[word]; ok {
				fmt.Printf("Слово найдено %s %s\n", word, val)
			} else {
				fmt.Println("Такого слова в списке нет!")

			}
		case 3:
			delete(words, word)
			fmt.Println("Выполнено.")
		case 4:
			if len(words) != 0 {
				for key, val := range words {
					fmt.Printf("%s: %s\n", key, val)
				}
			} else {
				fmt.Println("Словарь пуст!")
			}
		default:
			fmt.Println("Данного значения нет в списке операций. Повторите попытку.")
		}

	}
}
