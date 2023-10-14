package main

import "fmt"

func main() {
	a := [5]int{1, 3, 6, 8, 0}

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}

// Create an array of integers with 5 elements and print each element. (use simple for loop)
