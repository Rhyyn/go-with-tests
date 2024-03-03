package structsmethodsinterfaces

import "testing"

func TestCalcPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := CalcPerimeter(rectangle)
	expect := 40.0

	if got != expect {
		t.Errorf("got %.2f but expected %.2f", got, expect)
	}
}

// Table driven tests
func TestArea(t *testing.T) {
	areaTests := []struct {
		name   string
		shape  Shape
		expect float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12.0, Height: 6.0}, expect: 72.0},
		{name: "Rectangle", shape: Rectangle{Width: 44, Height: 13}, expect: 572.0},
		{name: "Circle", shape: Circle{Radius: 10}, expect: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, expect: 36.0},
		{name: "Triangle", shape: Triangle{Base: 12.45, Height: 6.12}, expect: 38.097},
	}

	for _, test := range areaTests {
		t.Run(test.name, func(t *testing.T) {
			hasArea := test.shape.Area()
			if hasArea != test.expect {
				t.Errorf("%#v hasArea %g but expected %g", test.shape, hasArea, test.expect)
			}
		})
	}
}

// Single tests
// func TestArea(t *testing.T) {
// 	checkArea := func(t testing.TB, shape Shape, expect float64) {
// 		t.Helper()
// 		got := shape.Area()

// 		if got != expect {
// 			t.Errorf("got %g but expected %g", got, expect)
// 		}
// 	}

// 	t.Run("testing calc area of a rectangle", func(t *testing.T) {
// 		rectangle := Rectangle{12.0, 6.0}
// 		checkArea(t, rectangle, 72.0)
// 	})
// 	t.Run("testing calc area of a circle", func(t *testing.T) {
// 		circle := Circle{10}
// 		checkArea(t, circle, 314.1592653589793)
// 	})
// }
