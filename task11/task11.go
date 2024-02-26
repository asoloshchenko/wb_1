package main

import "fmt"

// Реализовать пересечение двух неупорядоченных множеств.

//
func sliceToSet[T int | string](slice []T) map[T]struct{} {
	set := make(map[T]struct{}, len(slice)) // создание пустого множества
	for _, val := range slice {
		set[val] = struct{}{} // добавление в словарь
	}
	return set
}

func Intersect[T int | string](a, b map[T]struct{}) map[T]struct{} {
	m := make(map[T]struct{})

	for k := range b { // перебор второго множества
		if _, ok := a[k]; ok { // если элемент есть в первом множестве
			m[k] = struct{}{} // добавляем в новое множество
		}
	}
	return m
}
func main() {
	a := []string{"cat", "cat", "dog", "cat", "tree", "meow"}
	b := []string{"cat", "cat", "meow", "cat", "sun"}

	A := sliceToSet(a)
	B := sliceToSet(b)

	fmt.Println("Set A:", A)
	fmt.Println("Set B:", B)

	C := Intersect(B, A)
	fmt.Println("Intersection of A and B:", C)
}
