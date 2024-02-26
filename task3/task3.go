package main

import (
	"fmt"
)

// Дана последовательность чисел: 2,4,6,8,10.
// Найти сумму их квадратов с использованием
// конкурентных вычислений.

func main() {
	arr := []int{2, 4, 6, 8, 10}

	resultChan := make(chan int) // Канал для передачи данных

	for _, num := range arr {
		go func(n int) { // Запуск горутины для каждого числа
			resultChan <- n * n // Передача квадрата в канал
		}(num)
	}

	total := 0
	for range arr {
		total += <-resultChan // Sum up the squares from the channel
	}

	fmt.Println("Sum of squares:", total)
}
