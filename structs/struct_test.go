package structs

import (
	"go-by-example/utils"
	"testing"
)

func TestPerimeter(t *testing.T) {
	var rect Rectangle = Rectangle{Width: 10, Height: 10}
	got := rect.Perimeter()
	var want float32 = 40.0
	utils.Assert(t, got, want)
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float32) {
		t.Helper()
		got := shape.Area()
		utils.Assert(t, got, want)
	}
	t.Run("rectangles", func(t *testing.T) {
		var rect Rectangle = Rectangle{Width: 10, Height: 10}
		var want float32 = 100
		checkArea(t, rect, want)
	})
	t.Run("circle", func(t *testing.T) {
		var circle Circle = Circle{Radius: 10}
		var want float32 = 314.15927
		checkArea(t, circle, want)
	})
}
