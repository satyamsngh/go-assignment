package main

import (
	"fmt"
	"math/rand"
)

var i int

func main() {
	var a int
	var i int
	a = rand.Intn(10)
	fmt.Scanln(&i)
	Game(a, i)

}

func Game(a, i int) {
	if a < i {
		fmt.Println("its to high")
		fmt.Scanln(&i)
		Game(a, i)

	} else if a > i {
		fmt.Println("it's too low")
		fmt.Scanln(&i)
		Game(a, i)
	} else {
		fmt.Println("you guessed it right")
	}
}

// Create a simple guessing game.
// . Generate a random number between 1 and 10, and let the user guess the number.
// (use rand.Intn() method to generate random numbers)

// Example on how to use rand package to generate random number: https://gobyexample.com/random-numbers

// Ask the user for the inputs and Provide hints to the user like "too high" or "too low" until they guess the correct number. If the user guesses the number correctly say “you guessed it right” & quit.
