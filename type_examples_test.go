package slicelement

import (
	"fmt"

	"github.com/1046102779/slicelement"
)

func ExampleContain_Int() {
	var (
		data []int = []int{1, 2, 3, 4, 5}
		elem int   = 2
	)
	isExist, err := slicelement.Contains(data, elem, "")
	if err != nil {
		fmt.Println(err.Error())
	}
	if isExist {
		fmt.Println("elem in data is exist")
	} else {
		fmt.Println("elem in data is not exist")
	}
	return
}

func ExampleContain_Uint() {
	var (
		data []uint = []uint{1, 2, 3, 4, 5}
		elem uint   = 2
	)
	isExist, err := slicelement.Contains(data, elem, "")
	if err != nil {
		fmt.Println(err.Error())
	}
	if isExist {
		fmt.Println("elem in data is exist")
	} else {
		fmt.Println("elem in data is not exist")
	}
}

func ExampleContain_String() {
	var (
		data []string = []string{"abc", "def", "hig"}
		elem string   = "def"
	)
	isExist, err := slicelement.Contains(data, elem, "")
	if err != nil {
		fmt.Println(err.Error())
	}
	if isExist {
		fmt.Println("elem in data is exist")
	} else {
		fmt.Println("elem in data is not exist")
	}
}

func ExampleContain_Float32() {
	var (
		data []float32 = []float32{1, 2, 3, 4, 5}
		elem float32   = 2
	)
	isExist, err := slicelement.Contains(data, elem, "")
	if err != nil {
		fmt.Println(err.Error())
	}
	if isExist {
		fmt.Println("elem in data is exist")
	} else {
		fmt.Println("elem in data is not exist")
	}
}

func ExampleIndex_Int() {
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
	isExist, err := slicelement.Contains(data, elem, "Age")
	if err != nil {
		//fmt.Println(errors.Cause(err).Error())
		fmt.Println(err.Error())
	}
	if isExist {
		fmt.Println("elem in data is exist")
	} else {
		fmt.Println("elem in data is not exist")
	}
}
