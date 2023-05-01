package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	expected := 40.0

	if got != expected {
		t.Errorf("got %.2f expected %.2f", got, expected)
	}
}

func TestArea(t *testing.T) {

	common := func(t *testing.T, shape Shape, expected float64) {
		t.Helper()
		got := shape.Area()

		if got != expected {
			t.Errorf("got %.2f expected %.2f", got, expected)
		}
	}

	t.Run("should return area of rectangle", func(t *testing.T) {
		rectangle := Rectangle{2.0, 10.0}
		common(t, rectangle, 20.0)
	})

	t.Run("should return area of circle", func(t *testing.T) {
		circle := Circle{10}
		common(t, circle, 314.1592653589793)
	})

}

func TestAreaByTable(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{2.0, 10.0}, want: 20.0},
		{name: "Circle", shape: Circle{10}, want: 314.1592653589793},
	}

	for _, test := range areaTests {
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Area()
			if got != test.want {
				t.Errorf("got %.2f expected %.2f", got, test.want)
			}
		})

	}
}
