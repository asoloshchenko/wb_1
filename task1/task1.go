package main

import "fmt"

// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской
// структуры Human (аналог наследования)

type Human struct { // структура Human
	Name    string
	Surname string
	Age     int
}

func (h *Human) FullName() string { // метод FullName
	return h.Name + " " + h.Surname
}

func (h *Human) Introduce() { // метод Introduce
	fmt.Println("Hello, my name is", h.FullName(), "and I'm", h.Age)
}

type Action struct { // структура Action
	Human // встраивание (has-a relationship)
	Field int
}

func main() {
	a := &Action{ // создание экземпляра структуры Action
		Human: Human{
			Name:    "John",
			Surname: "Doe",
			Age:     30,
		},
		Field: 42,
	}

	a.Introduce() // вызов метода Introduce, принадлежащего структуре Human
}
