package main

import (
	"fmt"
	"trainingassignment/Day1/Assignment3/3b/temperature"
)

//var b int

func main() {
	var f int

	fmt.Scanln(&f)
	d := temperature.Temperature(f)

	fmt.Print("The temperature in celsius is :", d)
}
