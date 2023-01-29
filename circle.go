package main

import (
	"math"
)

// num 6
type Circle struct {
	radius float64
}

// 7.method to calulate area of circle
func (c Circle) Area() float64 {
	answer3 := math.Pi * math.Pow(c.radius, 2)
	return answer3
}

// 8. method to calculate perimeter of circle
func (c Circle) Perimeter() float64 {
	answer4 := 2 * math.Pi * c.radius
	return answer4
}

/*func main() {
	tinyCircle := Circle{
		radius: 6,
	}
	log.Println(tinyCircle.Area())
	log.Println(tinyCircle.Perimeter())
}*/
