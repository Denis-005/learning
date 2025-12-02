package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
Мини-проекта №4 — CLI-менеджер задач

1. Добавить задачу
2. Отметить задачу как выполненную
3. Удалить задачу
4. Показать все задачи
5. Показать только выполненные / только невыполненные
6. Сбросить статус задачи
7. Выход
*/

type Task struct {
	ID          int    // Уникальный номер
	Title       string // Название задачи
	Description string // Описание задачи
	Done        bool   // Выполнена ли задача
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
		id    int
		count int
	)

	numsId := []int{}
	tasks := make(map[int]*Task)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Введите цифру из списка задач!")
		fmt.Println("1. Добавить задачу")
		fmt.Println("2. Отметить задачу как выполненную")
		fmt.Println("3. Удалить задачу")
		fmt.Println("4. Показать все задачи")
		fmt.Println("5. Показать только выполненные / только невыполненные")
		fmt.Println("6. Сбросить статус задачи")
		fmt.Println("7. Выход")

		var num int
		_, err := fmt.Scanln(&num)
		if err != nil {
			fmt.Println("Неверное значение из списка!. Повторите попытку")
			continue
		}

		switch num {
		case 1:
			count++
			numsId = append(numsId, count)

			fmt.Println("Введите название задачи!")
			scanner.Scan()
			title := scanner.Text()

			fmt.Println("Введите описание задачи!")
			scanner.Scan()
			description := scanner.Text()

			tasks[count] = &Task{
				ID:          count, // Автоматическое присвоение ID.
				Title:       title,
				Description: description,
				Done:        false, // На данный момент задачи становиться автоматически не выполненной.
			}
		case 2:
			fmt.Println("Введите номер ID задачи которую выполнили!")
			fmt.Scanln(&id)

			if task, ok := tasks[id]; ok {
				task.Done = true
				fmt.Println("Задача отмечена как выполненная.")
			} else {
				fmt.Println("Задачи с таким ID нет.")
			}

		case 3:
			fmt.Println("Введите номер ID задачи которую хотите удалить!")
			fmt.Scanln(&id)
			if _, ok := tasks[id]; ok {
				delete(tasks, id)
				numsId = removeNum(numsId, id)
				fmt.Println("Задача успешно удалена.")
			} else {
				fmt.Println("Задачи с таким ID нет.")
			}
		case 4:
			if len(tasks) == 0 {
				fmt.Println("Список задач пуст.")
				continue
			}

			for _, val := range numsId {
				if task, ok := tasks[val]; ok {
					fmt.Printf("ID - %d, Название: %s, Описание: %s, Выполнено: %t\n", val, task.Title, task.Description, task.Done)
				} else {
					fmt.Printf("Под данным ID - %d задача не найдена\n", val)
				}
			}
		case 5:
			fmt.Println("Выберите цифру из списка!")
			fmt.Println("1 - Показать выполненные задачи")
			fmt.Println("2 - Показать невыполненные задачи")

			fmt.Scanln(&num)

			switch num {
			case 1:
				isFound := false
				for _, done := range tasks {
					if done.Done {
						fmt.Println(done.Title, done.Description)
						isFound = true
					}
				}

				if !isFound {
					fmt.Println("Нет выполненных задач!")
				}

			case 2:
				for _, noDone := range tasks {
					if !noDone.Done {
						fmt.Println(noDone.Title, noDone.Description)
					}
				}
			}
		case 6:
			fmt.Println("Введите номер ID задачи у которой хотите сбросить статус")
			fmt.Scanln(&id)

			if task, ok := tasks[id]; ok {
				task.Done = false
				fmt.Println("Статус задачи успешно сброшен!")
			} else {
				fmt.Println("Задачи с таким ID нет.")
			}
		case 7:
			return
		default:
			fmt.Println("Такого числа нет в списке. Повторите попытку!")
		}
	}
}
