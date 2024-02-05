package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func InitPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func main() {

	point1 := InitPoint(12.0, 16.0)
	point2 := InitPoint(12.0, 23.0)
	fmt.Printf("%.2f \n", Distance(point1, point2))
}

func Distance(p1, p2 Point) float64 {
	dx := p2.x - p1.x
	dy := p2.y - p1.y
	return math.Sqrt(dx*dx + dy*dy)
}
