// Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point с инкапсулированными
// параметрами x,y и конструктором.

package main

import (
	"fmt"
	"math"
)

// A struct for a 2D-Point
type Point struct {
	x float64
	y float64
}

// A Point Constructor
func NewPoint(x float64, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

// Calculates euclidean distance
func Distance(p1 *Point, p2 *Point) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
}

func main() {
	p1 := NewPoint(0, 0)
	p2 := NewPoint(3, 4)
	fmt.Println(Distance(p1, p2))
}
