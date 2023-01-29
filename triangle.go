package main

import (
	"log"
	"math"
)

// num1
type Triangle struct {
	base   float64
	height float64
}

// 2.method to calculate area
func (t Triangle) Area() float64 {
	answer := t.base * t.height * 0.5
	return answer
}

// 3.method to calculate perimeter
func (t Triangle) Perimeter() float64 {
	side_c := math.Sqrt(math.Pow(t.base, 2) + math.Pow(t.height, 2))
	answer2 := t.base + t.height + side_c
	return answer2
}

// num 4
func main() {
	rightAngle := Triangle{
		base:   3,
		height: 4,
	}
	log.Println(rightAngle.Area())
	log.Println(rightAngle.Perimeter())
}
