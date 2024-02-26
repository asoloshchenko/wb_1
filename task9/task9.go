package main

import (
	"fmt"
)

// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x)
// из массива, во второй — результат операции x*2, после чего данные из
// второго канала должны выводиться в stdout.

func main() {
	mslice := []int{1, 2, 3, 4, 5}
	ch1, ch2 := make(chan int), make(chan int)

	go func(ch_out chan<- int) {
		for _, it := range mslice {
			ch_out <- it
		}
		close(ch_out)
	}(ch1)

	go func(ch_in <-chan int, ch_out chan<- int) {
		for {
			if x, ok := <-ch_in; ok {
				ch_out <- x * x
			} else {
				close(ch_out)
				return
			}
		}
	}(ch1, ch2)

	for {
		if x, ok := <-ch2; ok {
			fmt.Println(x)
		} else {
			return
		}
	}

}

// def func(ch1, ch2 chan int) {
// 	ch1 <- 1
// }
