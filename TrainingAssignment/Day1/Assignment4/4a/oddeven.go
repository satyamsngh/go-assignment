package main

import "fmt"

func main() {

	var a int
	fmt.Scanln(&a)

	OddEven(a)

}

func OddEven(a int) {

	if a%2 == 0 {
		fmt.Println("its is an even number")
		return

	}
	fmt.Println("its an odd number")
}
