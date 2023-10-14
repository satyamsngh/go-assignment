package main

import (
	"fmt"
)

func main() {
	a := []string{"apple", "banana", "orange", "grapes"}

	for _, v := range a {
		fmt.Println(v)
	}
}
