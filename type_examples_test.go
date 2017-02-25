package slicelement

import (
	"fmt"

	"github.com/1046102779/slicelement"
)

func ExampleContains_int() {
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

func ExampleContains_uint() {
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

func ExampleContains_string() {
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

func ExampleContains_float32() {
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
