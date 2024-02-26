package main

import (
	"fmt"
	p "level1/task24/Point"
	"math"
)

func Dist(p1, p2 *p.Point) float64 {
	return math.Sqrt(float64((p1.X()-p2.X())*(p1.X()-p2.X()) + (p1.Y()-p2.Y())*(p1.Y()-p2.Y())))
}

func main() {
	p1 := p.NewPoint(-3, -4)
	p2 := p.NewPoint(3, 4)
	fmt.Println(Dist(p1, p2))
}
