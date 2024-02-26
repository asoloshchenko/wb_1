package main

import (
	"fmt"
)

// Реализовать паттерн «адаптер» на любом примере.
func main() {
	cat := &Cat{}
	dog := &Dog{}

	animals := []Adapter{
		NewCatAdapter(cat),
		NewDogAdapter(dog),
	}
	for _, a := range animals {
		a.Call()
	}
}

type Adapter interface { // Адаптер
	Call()
}

// Адаптируемые структуры
type Cat struct{}

func (c *Cat) Meow() {
	fmt.Println("meow")
}

type Dog struct{}

func (d *Dog) Bark() {
	fmt.Println("bark")
}

// ---------------------
type CatAdapter struct { // Адаптер кота
	*Cat
}

func (c *CatAdapter) Call() { // Метод, удовлетворяющий интерфейсу
	c.Meow()
}

type DogAdapter struct { // Адаптер собаки
	*Dog
}

func (d *DogAdapter) Call() { // Метод, удовлетворяющий интерфейсу
	d.Bark()
}

// Конструкторы адаптеров (Возвращает интерфейс)
func NewCatAdapter(c *Cat) Adapter { // Конструктор адаптера кота
	return &CatAdapter{c}
}
func NewDogAdapter(d *Dog) Adapter { // Конструктор адаптера собаки
	return &DogAdapter{d}
}
