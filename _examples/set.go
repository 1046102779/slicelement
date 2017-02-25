package main

import (
	"fmt"

	"github.com/1046102779/slicelement"
)

func main() {
	type Student struct {
		Name string
		Age  int
	}
	studentA := []Student{
		Student{
			Name: "donghai",
			Age:  29,
		},
		Student{
			Name: "jixaing",
			Age:  19,
		},
	}

	studentB := []Student{
		Student{
			Name: "Joe",
			Age:  18,
		},
		Student{
			Name: "David",
			Age:  19,
		},
	}
	if value, err := slicelement.GetUnion(studentA, studentB, "Age"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("studentA U studentB, result: ", value)
	}
	if value, err := slicelement.GetInteraction(studentA, studentB, "Age"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("studentA ∩ studentB, result: ", value)
	}
	if value, err := slicelement.GetDifference(studentA, studentB, "Age"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("studentA - studentB, result: ", value)
	}
}
