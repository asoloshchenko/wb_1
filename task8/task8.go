package main

import (
	"fmt"
)

// Дана переменная int64.
// Разработать программу которая устанавливает i-й бит в 1 или 0.

func raiseBit(n int64, i int) (int64, error) {
	if i < 0 || i > 64 {
		return 0, fmt.Errorf("i must be in range 0..63")
	}

	return int64(n | (1 << i)), nil // побитовое или
}

func lowerBit(n int64, i int) (int64, error) {
	if i < 0 || i > 64 {
		return 0, fmt.Errorf("i must be in range 0..63")
	}

	return int64(n &^ (1 << i)), nil // побитовое and not
}

func main() {
	var test2 int64 = 1
	fmt.Printf("%064b\n", test2)
	test2, _ = raiseBit(test2, 63)
	fmt.Printf("%064b\n", test2)
	test2, _ = lowerBit(test2, 63)
	fmt.Println(test2)
}
