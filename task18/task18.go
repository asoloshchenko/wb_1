package main

import (
	"fmt"
	"sync"
)

// Реализовать структуру-счетчик, которая будет инкрементироваться
// в конкурентной среде. По завершению программа должна выводить
// итоговое значение счетчика.
type Counter struct {
	counter int
	sync.Mutex
}

func (c *Counter) Increment() {
	c.Lock()
	defer c.Unlock()
	c.counter++
}

func (c *Counter) Value() int {
	return c.counter
}

func main() {
	var wg sync.WaitGroup
	counter := new(Counter)
	testCount := 1000 // Number of concurrent increments

	// Concurrently increment the counter
	for i := 0; i < testCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait() // Wait for all goroutines to finish

	// After all increments, the counter value should be equal to testCount
	finalValue := counter.Value()
	if finalValue != testCount {
		fmt.Printf("Test Failed: Expected counter value %d, got %d\n", testCount, finalValue)
	} else {
		fmt.Println("Test Passed")
	}

	fmt.Println("Final counter value:", finalValue)
}
