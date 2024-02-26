package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».
func main() {
	test := "snow dog sun"
	fmt.Println(ReverseWords(test))
}

func ReverseWords(s string) string {
	words := strings.Split(s, " ") // разделяем строку на слова
	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-1-i] = words[len(words)-1-i], words[i] // меняем местами
	}

	return strings.Join(words, " ")
}
