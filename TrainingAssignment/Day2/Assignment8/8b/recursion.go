package main

import "fmt"

func main() {

	number := 5
	a := Factorial(number)
	fmt.Println("the factorial of the number is", a)

}

func Factorial(number int) int {
	if number == 0 {
		return 1

	}
	return number * Factorial(number-1)
}

// Write a function that calculates the factorial of a given integer. Use recursion for this task.
