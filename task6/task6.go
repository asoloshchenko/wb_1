package main

import (
	"context"
	"fmt"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.

func main() {
	stopByVar()
}

func stopByChan() {
	stop := make(chan struct{}) // Канал, в который будем передавать сигнал для остановки

	go func() {
		for {
			select {
			case <-stop: // При получении сигнала
				fmt.Println("goutine stopped")
				return
			default:
				continue
			}
		}
	}()

	time.Sleep(500 * time.Millisecond)
	stop <- struct{}{} // Отправляем сигнал остановки
	time.Sleep(500 * time.Millisecond)
	fmt.Println("main finished")
}

func stopByCtx() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done(): // При получении сигнала отмены контекста
				fmt.Println("goutine stopped")
				return
			default:
				continue
			}
		}
	}()

	time.Sleep(500 * time.Millisecond)
	cancel() // Вызываем функцию отмены контекста
	time.Sleep(500 * time.Millisecond)
	fmt.Println("main finished")
}

func stopByVar() {
	stop := false

	go func() {
		for {
			if stop { // Захватываем переменную stop и если она true завершаем горутину
				fmt.Println("goutine stopped")
				return
			}
		}
	}()

	time.Sleep(500 * time.Millisecond)
	stop = true // Записываем в переменную stop значение true
	time.Sleep(500 * time.Millisecond)
	fmt.Println("main finished")
}
