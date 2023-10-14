package main

import "fmt"

type Shape interface {
	Area() float64
}
type Rectangle struct {
	Length float64
	Width  float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func ShapeArea(s Shape) {
	fmt.Println("the area is", s.Area())
}

func main() {

	circle := Circle{8.34}
	rectange := Rectangle{2.5, 6.7}
	ShapeArea(circle)
	ShapeArea(rectange)
}

// Description: Create a program that demonstrates the use of interfaces.
// Instructions:
// Define an interface called Shape with a method Area() that returns a float64.
// Create two struct types, Circle and Rectangle, both implementing the Shape interface with their respective Area() methods.
// Create instances of both Circle and Rectangle and calculate and print their areas using the Area() method.
