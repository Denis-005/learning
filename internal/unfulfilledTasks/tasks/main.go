package tasks

import "fmt"

// Замечательное число.
// Дано натуральное число - N. Определите является ли данное число замечательным.
// Число называется замечательным, если оно делится на сумму своих цифр без остатка.

func wonderfulNum() {
	var n int // Объявили переменную n.
	fmt.Scan(&n)
	d := 0
	sum := 0      // Объявили переменную и инициализировали переменную sum (Сумматор).
	isNum := true // Объявили и инициализировали булевую переменную isNum.

	i := 0      // Объявили и инициализировали переменную цикла i.
	for i < n { // В оператор цикла for записываем условное выражение, что i < n и если условное выражение принимает значение true то переходим в тело.
		//d := n % 10     // Объявили и инициализировали переменную d для временного хранения последней цифры числа в переменной n.
		sum += n        // Складываем в сумматор.
		n /= 10         // Избавляемся от последней цифры.
		if n%sum != 0 { // В условии if делим n на sum и если значение не равное 0 то преходим в тело блока if.
			isNum = false // Присваиваем в переменную isNum значение false.
		}
	}
	d = n % 10
	fmt.Println(d)

	if isNum { // Если в переменной isNum значение true то выводим аргумент блока if, если же false то аргумент блока else.
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
