package main

import (
	"testing"
)

func TestTri_Area(t *testing.T) {
	//function signature
	t.Run("triangle", func(t *testing.T) {
		triangle := Triangle{3.0, 4.0}
		got := triangle.Area()
		want := 6.0

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}

func TestTri_Per(t *testing.T) {
	t.Run("triangle", func(t *testing.T) {
		triangle := Triangle{3.0, 4.0}
		got := triangle.Perimeter()
		want := 12.0

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}
