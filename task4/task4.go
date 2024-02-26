package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

// Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные данные
// из канала и выводят в stdout. Необходима возможность выбора
// количества воркеров при старте
func main() {
	var nWorkers int
	if len(os.Args) > 1 { // Чтение количества воркеров из командной строки
		n, err := strconv.Atoi(os.Args[1])

		if err != nil {
			fmt.Println(err)
			return
		}
		if n <= 0 {
			fmt.Println("n must be > 0")
			return
		}
		nWorkers = n
	} else {
		nWorkers = 5 // Значение по умолчанию
	}

	c := make(chan os.Signal, 1)                   // Создание канала для сигнала Ctrl+C
	signal.Notify(c, os.Interrupt, syscall.SIGINT) // Подписка на Ctrl+C

	mainCh := make(chan int) // Создание канала для передачи значений

	wg := &sync.WaitGroup{} // Ожидание завершения всех воркеров

	wg.Add(nWorkers)
	for i := 0; i < nWorkers; i++ {
		go func(n int) {
			defer wg.Done()
			for {
				if val, ok := <-mainCh; ok {
					// Если канал не закрыт и есть данные
					fmt.Println("Worker ", n, "received", val)
				} else {
					// Если канал закрыт
					fmt.Println("Worker ", n, "finished")
					return
				}

			}

		}(i + 1) // Номер воркера
	}

	data := 0
	for {
		select {
		case <-c: // При получении сигнала Ctrl+C
			fmt.Println("\nCtrl+C pressed. Exiting...")
			close(mainCh) // Закрываем канал
			wg.Wait()     // Ожидаем завершения всех воркеров
			return
		default: // По умолчанию посылаем данные в канал
			mainCh <- data
			data++
		}
	}
}
