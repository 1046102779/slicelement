package slicelement

import (
	"fmt"

	"github.com/1046102779/slicelement"
)

func ExampleIndex_int() {
	var (
		data []int = []int{1, 2, 3, 4, 5}
		elem int   = 2
	)
	index, err := slicelement.GetIndex(data, elem, "")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("index=", index)
	return
}

func ExampleIndex_uint() {
	var (
		data []uint = []uint{1, 2, 3, 4, 5}
		elem uint   = 2
	)
	index, err := slicelement.GetIndex(data, elem, "")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("index=", index)
	return
}

func ExampleIndex_string() {
	var (
		data []string = []string{"abc", "def", "hig"}
		elem string   = "def"
	)
	index, err := slicelement.GetIndex(data, elem, "")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("index=", index)
	return
}

func ExampleIndex_float32() {
	var (
		data []float32 = []float32{1, 2, 3, 4, 5}
		elem float32   = 2
	)
	index, err := slicelement.GetIndex(data, elem, "")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("index=", index)
	return
}

func ExampleIndex_struct() {
	type Person struct {
		Name     string
		Age      int
		Children []string
	}
	data := []*Person{
		&Person{
			Name:     "John",
			Age:      29,
			Children: []string{"David", "Lily", "Bruce Lee"},
		},
		&Person{
			Name:     "Joe",
			Age:      18,
			Children: []string{},
		},
	}
	elem := 18
	index, err := slicelement.GetIndex(data, elem, "Age")
	if err != nil {
		//fmt.Println(errors.Cause(err).Error())
		fmt.Println(err.Error())
	}
	fmt.Println("index=", index)
	// output: index=1
}
