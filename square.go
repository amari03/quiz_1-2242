package main

import (
	"math"
	"log"
)


// this function only has one parameter 
//but it prints to 2 return values
func square(num float64) (float64, float64) {
	//area := math.Pow(num,2)
	//perimeter := 4 * num
	return (math.Pow(num, 2)), (4 * num)
}

func main() {
	var num = 4

	log.Println(square(float64(num)))
}
