package main

import (
	"fmt"
	"slices"
)

// Удалить i-ый элемент из слайса.
func main() {
	//a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	a := []int{1, 2}
	b := a[:]
	fmt.Println(b)

	a = DelOrig(a, 0)

	fmt.Println(a)
	fmt.Println(b)

}

func DelElemOrder(a []int, index int) []int {
	return append(a[:index], a[index+1:]...)
}

func DelElemOnurdered(a []int, index int) []int {
	a[index], a[len(a)-1] = a[len(a)-1], a[index]
	return a[:len(a)-1]
}

func DelOrig(a []int, index int) []int {
	return slices.Delete(a, index, index+1)
}
