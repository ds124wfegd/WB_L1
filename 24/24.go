package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{x, y}
}

func (p Point) Distance(other Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Hypot(dx, dy)
}

func main() {
	a := NewPoint(10, 10)
	b := NewPoint(0, 0)
	fmt.Println(a.Distance(b)) // ~14,1421
}
