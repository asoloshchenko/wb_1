package main

import (
	"fmt"
	"sync"
)

// Реализовать конкурентную запись данных в map.

type SafeMap struct {
	data map[int]int
	mx   sync.RWMutex
}

func (m *SafeMap) Set(key int, value int) {
	m.mx.Lock() // Блокировка мьютекса для чтения и записи
	// defer m.mx.Unlock()
	m.data[key] = value
	m.mx.Unlock()
}

func (m *SafeMap) Get(key int) (int, bool) {
	m.mx.RLock() // Блокировка мьютекса для записи
	defer m.mx.RUnlock()
	if val, ok := m.data[key]; ok {
		return val, true
	} else {
		return 0, false
	}
}

func main() {
	var wg sync.WaitGroup
	sm := SafeMap{data: make(map[int]int)} // Создаем мапу

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			sm.Set(val, val*val) // Сохраняем данные
		}(i)
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(key int) {
			defer wg.Done()
			_, ok := sm.Get(key) // Читаем данные
			if !ok {
				fmt.Printf("Key %d not found\n", key)
			}
		}(i)
	}

	wg.Wait() // Wait for all goroutines to finish.

}
