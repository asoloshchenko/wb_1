package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// К каким негативным последствиям может привести данный фрагмент кода,
// и как это исправить? Приведите корректный пример реализации.

// Проблема 1:
// Слайс ссылается на большой базовый массив
// Проблема 2:
// justString = v[:100] - взятие первых 100 байт строки,
// но это не равносильно первым 100 символам
var justString string

func createHugeString(size int) string {
	return strings.Repeat("abc Привет 🎈", 10000)
}
func someFunc() {
	v := createHugeString(1 << 10)
	start := 0
	runeSlice := make([]rune, 0, 100)
	for i := 0; i < 100; i++ {
		if start >= len(v) {
			break
		}
		r, size := utf8.DecodeRuneInString(v[start:])

		if r == utf8.RuneError {
			break
		}

		runeSlice = append(runeSlice, r)
		start += size
	}
	justString = string(runeSlice)
}

func main() {
	someFunc()
	fmt.Println(justString)
	fmt.Println(len(justString))
	fmt.Println(utf8.RuneCountInString(justString))
}
