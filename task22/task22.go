package main

import (
	"fmt"
	"math/big"
)

// Разработать программу, которая перемножает, делит, складывает,
// вычитает две числовых переменных a,b, значение которых > 2^20.
func main() {
	a := big.NewInt(1<<63 - 1)
	b := big.NewInt(1<<63 - 1)
	fmt.Printf("a:%d\nb:%d\n", a, b)

	fmt.Println("Div:" + fmt.Sprint(new(big.Int).Div(a, b)))
	fmt.Println("Mul:" + fmt.Sprint(new(big.Int).Mul(a, b)))
	fmt.Println("Sum:" + fmt.Sprint(new(big.Int).Add(a, b)))
	fmt.Println("Sub:" + fmt.Sprint(new(big.Int).Sub(a, b)))

}
