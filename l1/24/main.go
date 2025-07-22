package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func (p *Point) String() string {
	return fmt.Sprintf("(%f, %f)", p.x, p.y)
}
func (p *Point) Distance(other *Point) float64 {
	return math.Sqrt(
		math.Pow(p.x-other.x, 2) +
			math.Pow(p.y-other.y, 2))
}
func main() {
	a := NewPoint(1.0, 2.0)
	b := NewPoint(4.0, 6.0)
	fmt.Println(a.Distance(b))

}
