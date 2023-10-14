package main

import "fmt"

type Rectangle struct {
	Width  int
	Height int
}

func main() {
	a:=Rectangle{2,6}
	a.Calculate()

}

func (r Rectangle) Calculate() {
	fmt.Println("Perimeter of rectange is:", 2*(r.Height+r.Width))
	fmt.Println("Area of rectangle is:", r.Height*r.Width)
}

// Description: Create a struct with methods for common operations.
// Instructions:
// Define a struct called Rectangle with fields for Width and Height.
// Create methods for calculating the area and perimeter of a rectangle on this struct.
// Create an instance of the Rectangle struct and use the methods to calculate its area and perimeter.
// Print the results.
