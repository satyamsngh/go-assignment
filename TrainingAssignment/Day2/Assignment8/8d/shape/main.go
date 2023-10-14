package main

import (
	"fmt"
	"math"
	"trainingassignment/Day2/Assignment8/8d/model"
)

func main() {

	radius := model.Circle{23}
	Circle(radius)

}

func Circle(c model.Circle) {
	pi := math.Pi
	fmt.Println("The area of the circle is", 2*pi*float64(c.Radius)*float64(c.Radius))

}
