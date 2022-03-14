package main

import "fmt"

type Shaper interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.width * rectangle.height
}

type Square struct {
	width float64
}

func (square Square) Area() float64 {
	return square.width * square.width
}

type Circle struct {
	redius float64
}

func (circle Circle) Area() float64 {
	return circle.redius * circle.redius * 3.1415
}

func ComputeArea(shape Shaper) {
	fmt.Printf("Area of %+v=%.3f\n", shape, shape.Area())
}

func main() {
	rectangle := Rectangle{5, 10}
	square := Square{5}
	circle := Circle{5}

	shapes := [...]Shaper{rectangle, square, circle}

	for _, shape := range shapes {
		ComputeArea(shape)
	}

}
