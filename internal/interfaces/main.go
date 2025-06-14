package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return (r.Height + r.Width) * 2
}

func main() {
	rectangle := Rectangle{
		Width:  10,
		Height: 5,
	}

	fmt.Println(rectangle.Area(), rectangle.Perimeter())
}
