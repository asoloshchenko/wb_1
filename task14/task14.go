package main

import (
	"fmt"
	"reflect"
)

// Разработать программу, которая в рантайме способна
// определить тип переменной: int, string, bool, channel из переменной типа interface{}.
func main() {
	a := make(chan int)
	b := false
	s := "string"
	n := 1

	arr := []interface{}{a, s, n, b}
	for _, val := range arr {
		switch v := reflect.ValueOf(val); v.Kind() {
		case reflect.Int:
			fmt.Println("int")
		case reflect.String:
			fmt.Println("string")
		case reflect.Bool:
			fmt.Println("bool")
		case reflect.Chan:
			fmt.Printf("channel %v\n", v.Type().Elem())
		default:
			fmt.Printf("unexpected type %T\n", v)
		}
	}

}
