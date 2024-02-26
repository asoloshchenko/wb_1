package main

import (
	"fmt"
	"sync"
)

// Написать программу, которая конкурентно рассчитает значение
// квадратов чисел взятых из массива (2,4,6,8,10) и выведет их
// квадраты в stdout.

func main() {
	var wg sync.WaitGroup
	array := []int{2, 4, 6, 8, 10}

	wg.Add(len(array)) // increment counter
	for i := 0; i < len(array); i++ {
		go func(el int) {
			defer wg.Done() // same as wg.Done()
			fmt.Println(el * el)
		}(array[i]) // передаем копию, иначе возможно race condition
	}

	wg.Wait() // wait for all goroutines to finish

}
