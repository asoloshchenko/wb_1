package main

import (
	"fmt"
	"time"
)

// Разработать программу, которая будет последовательно отправлять
// значения в канал, а с другой стороны канала — читать.
// По истечению N секунд программа должна завершаться.

func main() {
	ver2()
}

// func ver1() {
// 	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

// 	mainCh := make(chan int)
// 	wg := &sync.WaitGroup{}
// 	wg.Add(2)

// 	go func(ctx context.Context) {
// 		defer wg.Done()
// 		for {
// 			select {
// 			case <-ctx.Done():
// 				close(mainCh)
// 				return
// 			default:
// 				time.Sleep(500 * time.Millisecond)
// 				mainCh <- rand.Int()
// 			}
// 		}
// 	}(ctx)

// 	go func(ctx context.Context) {
// 		defer wg.Done()
// 		for {
// 			if val, ok := <-mainCh; ok {
// 				fmt.Printf("Recived: %d\n", val)
// 			} else {
// 				fmt.Println("goroutine finished")
// 				return
// 			}

// 		}

// 	}(context.Background())
// 	wg.Wait()

// }

func ver2() {
	quit := time.After(5 * time.Second) // Канал, в который через 5 секунд принимаем сигнал
	mainCh := make(chan int)            // Канал для передачи значений

	go func() { // Горутина для отправки значений в канал
		i := 1
		for {
			mainCh <- i
			time.Sleep(500 * time.Millisecond)
			i++
			//fmt.Println("sent")
		}
	}()

	go func() { // Горутина для чтения значений из канала
		for {
			val := <-mainCh
			fmt.Printf("Recived: %d\n", val)
		}

	}()
	<-quit // Ждем сигнала и завершаем программу main
}
