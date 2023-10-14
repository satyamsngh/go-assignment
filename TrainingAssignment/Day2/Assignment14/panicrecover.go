package main

import "fmt"

func main() {
	safeDivide()
	fmt.Println("recover from the panic")

}
func Recover() {
	err := recover()
	if err != nil {
		fmt.Println(err)
		return
	}
}
func safeDivide() {
	var n, d int
	fmt.Scanln(&n)
	fmt.Scanln(&d)
	Divide(n, d)
}

func Divide(n int, d int) {
	defer Recover()
	if d == 0 {
		panic("it can't be divide")
	}
	fmt.Println(n / d)
	
}

// Description: In this assignment, you will create a simple program that simulates a situation where a panic occurs, and you'll use recover and defer to gracefully handle the panic.
// Instructions:

// Write a function called divide that takes two integer arguments, numerator and denominator. Inside the function, check if the denominator is zero. If it is, intentionally raise a panic using the panic function with an appropriate error message.
// Create another function called safeDivide that calls the divide function and uses the recover function to handle any panics that may occur.
// In the main function, call safeDivide with different values for the numerator and denominator. Ensure that you test cases where the denominator is both zero and non-zero.
// Use a defer statement in the safeDivide function to recover from any panics and print an error message instead of crashing the program.
// Test your program by calling safeDivide with different inputs and observe how it gracefully handles panics and continues execution
