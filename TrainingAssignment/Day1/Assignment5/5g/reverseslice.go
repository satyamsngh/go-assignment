package main

import "fmt"

func main() {
	number := []int{1, 2, 3, 4, 5, 6, 7}
	var b int
	for i := 0; i < len(number)/2+1; i++ {

		b = number[i]
		number[i] = number[(len(number)-1)-i]
		number[len(number)-1-i] = b
	}
	fmt.Println(number)

}

// Description: Write a program that reverses the order of elements in a slice.
