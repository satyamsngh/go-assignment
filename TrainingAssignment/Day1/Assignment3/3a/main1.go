package main

import (
	"fmt"
	"trainingassignment/Day1/Assignment3/3a/calculator"
)

var a int
var b int

func main() {

	fmt.Scanf("Enter the number:", &a)
	fmt.Scanf("Enter the number:", &b)
	fmt.Println(calculator.Difference(a, b))
	fmt.Println(calculator.Product(a, b))
	fmt.Println(calculator.QuoRem(a, b))
	fmt.Println(calculator.Square(a))
	fmt.Println(calculator.Sum(a, b))
}
