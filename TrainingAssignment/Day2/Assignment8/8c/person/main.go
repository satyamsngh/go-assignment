package main

import (
	"fmt"
	"trainingassignment/Day2/Assignment8/8c/model"
)

func main() {
	Person1 := model.Person{"satyam", 22}
	person2 := model.Person{"sandeep", 45}
	PrintPerson(Person1)
	PrintPerson(person2)

}

func PrintPerson(p model.Person) {
	fmt.Println("the name of the person is", p.Name, "and the age", p.Age)
}
