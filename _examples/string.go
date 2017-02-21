package main

import (
	"fmt"

	"github.com/1046102779/slicelement"
)

func main() {
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
