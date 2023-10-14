package main

import "fmt"

type Animal interface {
	MakeSound() string
}
type Dog struct {
	Sound string
}

type Cat struct {
	Sound string
}

func (d Dog) MakeSound() string {
	return d.Sound

}
func (c Cat) MakeSound() string {
	return c.Sound
}

func Sounds(a Animal) string {
	return a.MakeSound()
}

func main() {
	cat := Cat{"meeow"}
	dog := Dog{"bhooo"}
	fmt.Println(Sounds(cat))
	fmt.Println(Sounds(dog))

}

// Description: Create a program that demonstrates interface implementation with multiple structs.
// Instructions:
// Define an interface Animal with a method MakeSound() that returns a string.
// Create two structs Dog and Cat, both implementing the Animal interface with their respective MakeSound() methods.
// Create instances of Dog and Cat and call their MakeSound() methods to print the sounds they make.
