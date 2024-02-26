package main

import (
	"fmt"
	"time"
)

func sleep(milliseconds int) {
	start := time.Now()
	for time.Since(start).Milliseconds() < int64(milliseconds) {
		// ждем, пока не пройдет достаточно времени
	}
}

func main() {
	fmt.Println("Wait..")
	sleep(10000)
	fmt.Println("Hello, world!")
}
