package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

func main() {
	randomNumbers := make([]int, 10_000)
	for i := 0; i < cap(randomNumbers)-1; i++ {
		randomNumbers[i] = rand.Intn(10_000)
	}
	start := time.Now()
	QuickSort1(randomNumbers)
	elapsed1 := time.Since(start)
	start = time.Now()
	QuickSort2(randomNumbers)
	elapsed2 := time.Since(start)
	fmt.Printf("1)%v\n2)%v\n", elapsed1, elapsed2)
	fmt.Println("diff:" + fmt.Sprint(elapsed1-elapsed2))

}

func QuickSort1(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	pivot := a[len(a)/2]
	var left, right []int
	for i := 1; i < len(a); i++ {
		if a[i] < pivot {
			left = append(left, a[i])
		} else {
			right = append(right, a[i])
		}
	}
	left = QuickSort1(left)
	right = QuickSort1(right)
	return append(append(left, pivot), right...)
}

func QuickSort2(a []int) []int {
	return quickSortBoundary(a, 0, len(a)-1)
}

// рекурсивно вызывает сортировку
func quickSortBoundary(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSortBoundary(arr, low, p-1)
		arr = quickSortBoundary(arr, p+1, high)
	}
	return arr
}

// меняет массив и
// возвращает индекс
// опорного элемента
func partition(arr []int, low, high int) ([]int, int) { // меняет массив и
	// возвращает индекс
	pivot := arr[high] // опорного элемента
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}
