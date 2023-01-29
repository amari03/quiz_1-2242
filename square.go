package main

import (
	"math"
)

type Shape interface {
	Area() float64
}

func square(num float64) (float64, float64) {
	//area := math.Pow(num,2)
	//perimeter := 4 * num
	return (math.Pow(num, 2)), (4 * num)
}

/*func main() {
	var num = 4

	log.Println(square(float64(num)))
}*/
