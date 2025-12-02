//Готовое решение задачи а моё решение идёт следующим.

//Дана строка s, найдите длину самой длинной подстроки без повторяющихся символов.

//package main
//
//import "fmt"
//
//func main() {
//	s := "abcabcbb"
//	runes := []rune(s)
//	seen := make(map[rune]int)
//	maxLen := 0
//	start := 0
//
//	for i, r := range runes {
//		if pos, ok := seen[r]; ok && pos >= start {
//			start = pos + 1 // сдвигаем начало окна
//		}
//		seen[r] = i
//		if i-start+1 > maxLen {
//			maxLen = i - start + 1
//		}
//	}
//
//	fmt.Println(maxLen)
//}

// Моя версия.

package main

import (
	"fmt"
)

func main() {
	word := "abcbdef"
	symbols := []rune(word)

	maxLen, count := 0, 1
	mapSymbol := make(map[rune]struct{})
	mapSymbol[symbols[0]] = struct{}{}

	for i := 1; i < len(symbols); i++ {
		if symbols[i-1] != symbols[i] {
			if _, ok := mapSymbol[symbols[i]]; !ok {
				count++
			}

			mapSymbol[symbols[i]] = struct{}{}
		} else {
			mapSymbol = make(map[rune]struct{})
			mapSymbol[symbols[i]] = struct{}{}
			count = 1
		}

		if count > maxLen {
			maxLen = count
		}
	}

	fmt.Println(maxLen)
}
