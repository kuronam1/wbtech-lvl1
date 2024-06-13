package main

import (
	"fmt"
	"math"
)

var (
	x1 = 1.0
	x2 = 4.0
	y1 = 2.0
	y2 = 6.0
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func Distance(point1, point2 Point) float64 {
	return math.Sqrt(SquareSub(point2.x, point1.x) + SquareSub(point2.y, point1.y))
}

func SquareSub(num1, num2 float64) float64 {
	result := num1 - num2
	return math.Pow(result, 2)
}

func main() {
	firstPoint := NewPoint(x1, y1)
	secondPoint := NewPoint(x2, y2)

	fmt.Printf("%f\n", Distance(firstPoint, secondPoint))
}
