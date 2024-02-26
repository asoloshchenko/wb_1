package Point

type Point struct {
	x, y int
}

func NewPoint(x, y int) *Point {
	return &Point{x, y}
}

func (p *Point) X() int {
	return p.x
}

func (p *Point) Y() int {
	return p.y
}
