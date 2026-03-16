package main

import(
	"fmt"
)

type Shape interface {
	Area() float64
	Length() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c Circle) Length() float64 {
	return 2 * 3.14 * c.radius
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Length() float64 {
	return 2 * (r.width + r.height)
}

func main() {
	circle := Circle{radius: 5}
	rectangle := Rectangle{width: 4, height: 6}
	fmt.Println(circle.Area())    // Output: 78.5
	fmt.Println(circle.Length())  // Output: 31.4
	fmt.Println(rectangle.Area()) // Output: 24
	fmt.Println(rectangle.Length()) // Output: 20
}