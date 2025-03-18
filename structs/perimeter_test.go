package structs

import "testing"

//	func TestPerimeter(t *testing.T) {
//		got := Perimeter(10.0, 10.0)
//		want := 40.0
//
//		if got != want {
//			t.Errorf("got %.2f want %.2f", got, want)
//		}
//	}
//
//	func TestArea(t *testing.T) {
//		got := Area(12.0, 6.0)
//		want := 72.0
//
//		if got != want {
//			t.Errorf("got %.2f want %.2f", got, want)
//		}
//	}
func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

//func TestArea(t *testing.T) {
//	rectangle := Rectangle{12.0, 6.0}
//	got := rectangle.Area()
//	want := 72.0
//
//	if got != want {
//		t.Errorf("got %.2f want %.2f", got, want)
//	}
//}

func TestArea(t *testing.T) {

	//checkArea := func(t *testing.T, shape Shape, want float64) {
	//	t.Helper()
	//	got := shape.Area()
	//	if got != want {
	//		t.Errorf("got %.2f want %.2f", got, want)
	//	}
	//}
	//
	//t.Run("rectangles", func(t *testing.T) {
	//	rectangle := Rectangle{12, 6}
	//	//got := rectangle.Area()
	//	want := 72.0
	//
	//	//if got != want {
	//	//	t.Errorf("got %g want %g", got, want)
	//	//}
	//	checkArea(t, rectangle, want)
	//})
	//
	//t.Run("circles", func(t *testing.T) {
	//	circle := Circle{10}
	//	//got := circle.Area()
	//	want := 314.1592653589793
	//
	//	//if got != want {
	//	//	t.Errorf("got %g want %g", got, want)
	//	//}
	//	checkArea(t, circle, want)
	//})
	tcs := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "rectangle", shape: Rectangle{10.0, 10.0}, want: 100.0},
		{name: "circle", shape: Circle{10.0}, want: 314.1592653589793},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.shape.Area()
			if got != tc.want {
				t.Errorf("%#v got %.2f want %.2f", tc.shape, got, tc.want)
			}
		})
	}
}
