package main

import (
	"testing"
)

func TestCircle_Area(t *testing.T) {
	t.Run("circle", func(t *testing.T) {
		circle := Circle{6.0}
		got := circle.Area()
		want := 113.09733552923255

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}

func TestCircle_Per(t *testing.T) {
	t.Run("circle", func(t *testing.T) {
		circle := Circle{6.0}
		got := circle.Perimeter()
		want := 37.69911184307752

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}
