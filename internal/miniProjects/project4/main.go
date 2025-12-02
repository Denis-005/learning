// 1 - версия мини-проекта.

/*
   Мини-проект №5 — CLI-магазин (инвентарь товаров)

   1. Добавить товар
   2. Найти товар по названию
   3. Обновить данные товара (название, цену или количество)
   4. Удалить товар
   5. Показать все товары
   6. Показать товары дешевле заданной цены
   7. Показать товары, которых меньше определённого количества
   0. Выход
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

type Shop struct {
	ID       int
	Name     string
	Price    float64
	Quantity int
}

func removeNum(nums []int, num int) []int {
	for i, val := range nums {
		if val == num {
			return append(nums[:i], nums[i+1:]...)
		}
	}

	return nums
}

func main() {
	var (
		id              int
		count           int
		num             int
		price           float64
		quantity        int
		isRepeatAttempt = true
		numFromList     int
	)

	numsID := []int{}
	shop := make(map[int]*Shop)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if isRepeatAttempt {
			fmt.Println("Введите цифру из списка операции.")
			fmt.Println("1. Добавить товар")
			fmt.Println("2. Найти товар по названию")
			fmt.Println("3. Обновить данные товара (название, цену или количество)")
			fmt.Println("4. Удалить товар")
			fmt.Println("5. Показать все товары")
			fmt.Println("6. Показать товары дешевле заданной цены")
			fmt.Println("7. Показать товары, которых меньше определённого количества")
			fmt.Println("0. Выход")

			_, err := fmt.Scanln(&num)
			if err != nil {
				fmt.Println("Неверный ввод. Повторите попытку!")
				continue
			}
		}

		isFound := false
		switch num {
		case 0:
			return
		case 1:
			count++
			numsID = append(numsID, count)

			fmt.Println("Введите название товара!")
			scanner.Scan()
			name := scanner.Text()

			fmt.Println("Введите цену данного товара! Цена товара не должна быть отрицательным n < 0")
			fmt.Scan(&price)

			if price < 0 {
				fmt.Println("Цена должна быть больше нуля!")
				continue
			}

			fmt.Println("Введите количество одного товара! Количество одного товара не должна быть отрицательным n < 0")
			fmt.Scan(&quantity)

			if quantity < 0 {
				fmt.Println("Количество не может быть отрицательным!")
				continue
			}

			shop[count] = &Shop{
				count,
				name,
				price,
				quantity,
			}

			fmt.Println("Товар добавлен!")
		case 2:
			fmt.Println("Введите название товара!")
			scanner.Scan()
			name := scanner.Text()

			for _, nameProduct := range shop {
				if nameProduct.Name == name {
					fmt.Printf("Найден: %s\n", nameProduct.Name)
					isFound = true
				}
			}

			if !isFound {
				fmt.Println("Под таким название товар не найден. Повторите попытку!")
				isRepeatAttempt = false
				continue
			}

			isRepeatAttempt = true
		case 3:
			fmt.Println("Введите номер ID товара!")
			fmt.Scanln(&id)

			if _, ok := shop[id]; !ok {
				fmt.Println("Товар с таким ID не найден!")
				isRepeatAttempt = false
				continue
			}

			fmt.Println("Введите цифру операции из списка!")
			fmt.Println("1. Изменить название товара")
			fmt.Println("2. Изменить цену товара")
			fmt.Println("3. Изменить количество одного товара")

			_, err := fmt.Scanln(&numFromList)
			if err != nil {
				fmt.Println("Неверный ввод. Повторите попытку!")
			}

			switch numFromList {
			case 1:
				fmt.Println("Введите новое нзавание товара!")
				scanner.Scan()
				name := scanner.Text()

				shop[id].Name = name
			case 2:
				fmt.Println("Введите новую цену товара!")
				fmt.Scanln(&price)

				shop[id].Price = price
			case 3:
				fmt.Println("Введите новое количество одного товара!")
				fmt.Scanln(&quantity)

				shop[id].Quantity = quantity
			default:
				fmt.Println("Такого значения нет в списке. Повторите попытку!")
				isRepeatAttempt = false
				continue
			}

			isRepeatAttempt = true
		case 4:
			fmt.Println("Введите номер ID товара!")
			fmt.Scanln(&id)

			if _, ok := shop[id]; ok {
				delete(shop, id)
				numsID = removeNum(numsID, id)
				fmt.Println("Удаление завершено!")

			} else {
				fmt.Println("Под таким ID нет товара в списке!")
				isRepeatAttempt = false
				continue
			}

			isRepeatAttempt = true
		case 5:
			if len(shop) > 0 {
				for _, product := range numsID {
					fmt.Printf("ID товара: %d, Название: %s, Цена: %.2f, Количество: %d\n", product, shop[product].Name, shop[product].Price, shop[product].Quantity)
				}
			} else {
				fmt.Println("Список пуст!")
			}

		case 6:
			fmt.Println("Введите цену!")
			fmt.Scanln(&price)

			fmt.Println("Список товаров ниже указанной цены!")
			for _, priceProduct := range shop {
				if priceProduct.Price < price {
					fmt.Printf("Название %s, цена %.2f\n", priceProduct.Name, priceProduct.Price)
					isFound = true
				}
			}

			if !isFound {
				fmt.Println("Нет товаров дешевле указанной цены. Введите другую цену!")
				isRepeatAttempt = false
				continue
			}

			isRepeatAttempt = true
		case 7:
			fmt.Println("Введите количество одного товара!")
			fmt.Scanln(&quantity)

			fmt.Println("Список товаров: ")
			isFound = false
			for _, product := range shop {
				if product.Quantity < quantity {
					fmt.Printf("Название: %s, Цена: %.2f, Количество: %d\n", product.Name, product.Price, product.Quantity)
					isFound = true
				}
			}

			if !isFound {
				fmt.Println("Пуст")
				isRepeatAttempt = false
				continue
			}

			isRepeatAttempt = true
		default:
			fmt.Println("Такого значения нет в списке операций. Повторите попытку")
		}

	}
}

//2 - версия мини-проекта.

/*
   Мини-проект №5 — CLI-магазин (инвентарь товаров)

   1. Добавить товар
   2. Найти товар по названию
   3. Обновить данные товара (название, цену или количество)
   4. Удалить товар
   5. Показать все товары
   6. Показать товары дешевле заданной цены
   7. Показать товары, которых меньше определённого количества
   0. Выход
*/

/*
package main

import (
	"bufio"
	"fmt"
	"os"
)

type Product struct {
	ID       int
	Name     string
	Price    float64
	Quantity int
}

func addProduct(product map[int]*Product, id, quantity int, name string, price float64) {
	product[id] = &Product{
		ID:       id,
		Name:     name,
		Price:    price,
		Quantity: quantity,
	}
}

func findProductByName(product map[int]*Product, name string) bool {
	isFound := false
	for _, nameProduct := range product {
		if nameProduct.Name == name {
			fmt.Printf("Найден: %s\n", nameProduct.Name)
			isFound = true
		}
	}

	return isFound
}

func updateProduct(product map[int]*Product, id, choice int, scanner *bufio.Scanner) bool {
	var (
		showMenu = true
		price    float64
		quantity int
	)

	switch choice {
	case 1:
		fmt.Println("Введите новое название товара!")
		scanner.Scan()
		name := scanner.Text()

		product[id].Name = name
	case 2:
		fmt.Println("Введите новую цену товара!")
		_, err := fmt.Scanln(&price)
		if err != nil {
			showMenu = false
			fmt.Println("Неверный ввод. Повторите попытку")
		}

		product[id].Price = price
	case 3:
		fmt.Println("Введите новое количество одного товара!")

		_, err := fmt.Scanln(&quantity)
		if err != nil {
			showMenu = false
			fmt.Println("Неверный ввод. Повторите попытку")
		}

		product[id].Quantity = quantity
	}

	return showMenu
}

func deleteProduct(product map[int]*Product, productIDs []int, id int) []int {
	delete(product, id)
	productIDs = removeID(productIDs, id)

	return productIDs
}

func printProducts(products map[int]*Product, productIDs []int) {
	for _, id := range productIDs {
		fmt.Printf("ID товара: %d, Название: %s, Цена: %.2f, Количество: %d\n", id, products[id].Name, products[id].Price, products[id].Quantity)
	}
}

func printProductsBelowPrice(products map[int]*Product, price float64) bool {
	isFound := false
	for _, priceProduct := range products {
		if priceProduct.Price < price {
			fmt.Printf("Название %s, цена %.2f\n", priceProduct.Name, priceProduct.Price)
			isFound = true
		}
	}

	return isFound
}

func printProductsBelowQuantity(products map[int]*Product, quantity int) bool {
	isFound := false
	for _, product := range products {
		if product.Quantity < quantity {
			fmt.Printf("Название: %s, Цена: %.2f, Количество: %d\n", product.Name, product.Price, product.Quantity)
			isFound = true
		}
	}

	return isFound
}

func main() {
	var (
		id       int
		count    int
		action   int
		price    float64
		quantity int
		choice   int
	)

	productIDs := []int{}
	products := make(map[int]*Product)
	scanner := bufio.NewScanner(os.Stdin)

	showMenu := true
	for {
		if showMenu {
			fmt.Println("Введите цифру из списка операции.")
			fmt.Println("1. Добавить товар")
			fmt.Println("2. Найти товар по названию")
			fmt.Println("3. Обновить данные товара (название, цену или количество)")
			fmt.Println("4. Удалить товар")
			fmt.Println("5. Показать все товары")
			fmt.Println("6. Показать товары дешевле заданной цены")
			fmt.Println("7. Показать товары, которых меньше определённого количества")
			fmt.Println("0. Выход")

			_, err := fmt.Scanln(&action)
			if err != nil {
				fmt.Println("Неверный ввод. Повторите попытку!")
				continue
			}
		}

		showMenu = true
		isFound := false

		switch action {
		case 0:
			return
		case 1:
			count++
			productIDs = append(productIDs, count)

			fmt.Println("Введите название товара!")
			scanner.Scan()
			name := scanner.Text()

			fmt.Println("Введите цену товара (не отрицательную):")
			fmt.Scanln(&price)

			if price < 0 {
				fmt.Println("Ошибка: цена не может быть отрицательной!")
				continue
			}

			fmt.Println("Введите количество товара (не отрицательное):")
			fmt.Scanln(&quantity)

			if quantity < 0 {
				fmt.Println("Ошибка: количество не может быть отрицательным!")
				continue
			}

			addProduct(products, count, quantity, name, price)
			fmt.Println("Товар добавлен!")
		case 2:
			fmt.Println("Введите название товара!")
			scanner.Scan()
			name := scanner.Text()

			isFound = findProductByName(products, name)
			if !isFound {
				fmt.Println("Под таким название товар не найден. Повторите попытку!")
				showMenu = false
				continue
			}

		case 3:
			fmt.Println("Введите номер ID товара!")
			fmt.Scanln(&id)

			if _, ok := products[id]; !ok {
				fmt.Println("Товар с таким ID не найден!")
				showMenu = false
				continue
			}

			fmt.Println("Введите цифру операции из списка!")
			fmt.Println("1. Изменить название товара")
			fmt.Println("2. Изменить цену товара")
			fmt.Println("3. Изменить количество одного товара")

			_, err := fmt.Scanln(&choice)
			if err != nil {
				fmt.Println("Неверный ввод. Повторите попытку!")
			}

			if choice >= 1 && choice <= 3 {
				showMenu = updateProduct(products, id, choice, scanner)
			} else {
				fmt.Println("Такого значения нет в списке. Повторите попытку!")
				showMenu = false
				continue
			}

			if showMenu {
				fmt.Println("Выполнено!")
			} else {
				fmt.Println("Операция не выполнена. Повторите попытку!")
				continue
			}
		case 4:
			fmt.Println("Введите номер ID товара!")
			fmt.Scanln(&id)

			if _, ok := products[id]; ok {
				productIDs = deleteProduct(products, productIDs, id)
				fmt.Println("Удаление завершено!")
			} else {
				fmt.Println("Под таким ID нет товара в списке!")
				showMenu = false
				continue
			}
		case 5:
			if len(products) > 0 {
				printProducts(products, productIDs)
			} else {
				fmt.Println("Список пуст!")
			}
		case 6:
			fmt.Println("Введите цену!")
			fmt.Scanln(&price)

			fmt.Println("Список товаров ниже указанной цены!")

			isFound = printProductsBelowPrice(products, price)
			if !isFound {
				fmt.Println("Нет товаров дешевле указанной цены. Введите другую цену!")
				showMenu = false
				continue
			}
		case 7:
			fmt.Println("Введите количество одного товара!")
			fmt.Scanln(&quantity)

			fmt.Println("Список товаров: ")
			isFound = printProductsBelowQuantity(products, quantity)

			if !isFound {
				fmt.Println("Пуст")
				showMenu = false
				continue
			}
		default:
			fmt.Println("Такого значения нет в списке операций. Повторите попытку")
		}

	}
}

func removeID(nums []int, id int) []int {
	for i, val := range nums {
		if val == id {
			return append(nums[:i], nums[i+1:]...)
		}
	}

	return nums
}
*/
