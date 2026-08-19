package main

import "testing"

func CheckEquals(t testing.TB, got, want float64) {
	t.Helper()

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{12.5, 3.5}, want: 43.75},
		{name: "Circle", shape: Circle{10.0}, want: 314.1592653589793},
		{name: "Square", shape: Square{10.0}, want: 20.0},
		{name: "Triangle", shape: Triangle{12, 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()

			if got != tt.want {
				t.Errorf("%#v got %.2f, want %.2f", tt.shape, got, tt.want)
			}
		})
	}
}
