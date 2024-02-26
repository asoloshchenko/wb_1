package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree)
// создать для нее собственное множество.

func sliceToSet[T int | string](slice []T) map[T]struct{} {
	set := make(map[T]struct{}, len(slice))
	for _, val := range slice {
		set[val] = struct{}{}
	}
	return set
}
func main() {
	a := []string{"cat", "cat", "dog", "cat", "tree"}
	A := sliceToSet(a)
	fmt.Print("A: {")
	for k := range A {
		fmt.Print(k, ", ")
	}
	fmt.Println("}")
}
