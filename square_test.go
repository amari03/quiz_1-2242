package main

import "testing"

func TestSquare_Area(t *testing.T) {
	//the ,_ is called Blank Identifier
	//it is used to ignore one of the return types because the function 
	//being tested here has 2 return types
	result, _ := square(4.0)
	got := result
	want := 16.0

	if got != want {
		t.Errorf("got %g want %g,", got, want)
	}
}

func TestSquare_Perimeter(t *testing.T) {
	_, result := square(4.0)
	got := result
	want := 16.0

	if got != want {
		t.Errorf("got %g want %g,", got, want)
	}
}

//Use of g will print a more precise decimal number 
//in the error message