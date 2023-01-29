package main

import "testing"

func TestSquare_Area(t *testing.T) {
	result, _ := square(4.0)
	got := result
	want := 16.0

	if got != want {
		t.Errorf("got %g want %g,", want, got)
	}
}

func TestSquare_Perimeter(t *testing.T) {
	_, result := square(4.0)
	got := result
	want := 16.0

	if got != want {
		t.Errorf("got %g want %g,", want, got)
	}
}
