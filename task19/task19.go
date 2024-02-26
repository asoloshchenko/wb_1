package main

import "fmt"

// Разработать программу, которая
// переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
// Символы могут быть unicode.
func main() {
	test := "Аыфавыф, 世界! 👩🏾‍🦰👱🏾🧑🏾‍⚖️"

	fmt.Println(Reverse(test))
}

func Reverse(s string) string {
	temp := []rune(s)
	lenght := len(temp)
	for i := 0; i < lenght/2; i++ {
		temp[i], temp[lenght-1-i] = temp[lenght-1-i], temp[i]
	}
	return string(temp)
}
