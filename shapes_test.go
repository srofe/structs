package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10., 10.0}
	got := Perimeter(rectangle)
	want := 40.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		text  string
		shape Shape
		want  float64
	}{
		{"A Rectangle shall calculate its area", Rectangle{Width: 12, Height: 6}, 72},
		{"A Circle shall calculate its area", Circle{Radius: 10}, 314.1592653589793},
		{"A Triangle shall calculate its area", Triangle{Base: 12, Height: 6}, 36},
	}

	for _, tt := range areaTests {
		t.Run(tt.text, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("got %.2f want %.2f", got, tt.want)
			}
		})
	}
}
