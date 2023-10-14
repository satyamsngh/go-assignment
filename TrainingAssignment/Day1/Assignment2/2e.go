package main

import "fmt"

var (
	const a int    = 4
	const b string = "hello"
)

func main() {

	a = 5
	b = "world"
	fmt.Print(a, b)

	//syntax error: unexpected const, expected name

}
