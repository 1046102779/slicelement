package slicelement

import (
	"fmt"
)

// union
func ExampleGetUnion_int_int() {
	dataA := []int{1, 2, 3, 4, 5}
	dataB := []int{2, 4, 6, 7}
	temp, err := GetUnion(dataA, dataB, "")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("temp: ", temp)
	return
}

func ExampleGetUnion_uint_uint() {
	dataA := []uint{1, 2, 3, 4, 5}
	dataB := []uint{2, 4, 6, 7}
	temp, err := GetUnion(dataA, dataB, "")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("temp: ", temp)
	return
}

func ExampleGetUnion_float32_float32() {
	dataA := []float32{1, 2, 3, 4, 5}
	dataB := []float32{2, 4, 6, 7}
	temp, err := GetUnion(dataA, dataB, "")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("temp: ", temp)
	return
}

func ExampleGetUnion_string_string() {
	str1, str2, str3, str4, str5 := "1", "2", "3", "4", "5"
	dataA := []*string{&str1, &str2, &str3}
	dataB := []*string{&str2, &str3, &str4, &str5}
	temp, err := GetUnion(dataA, dataB, "")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("temp: ", temp)
}

func ExampleGetUnion_struct_struct() {
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
	if value, err := GetUnion(studentA, studentB, "Age"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("studentA U studentB, result: ", value)
	}
}

// interaction
func ExampleGetInteraction_int_int() {
	dataA := []int{1, 2, 3, 4, 5}
	dataB := []int{2, 4, 6, 7}
	temp, err := GetInteraction(dataA, dataB, "")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("temp: ", temp)
	return
}

func ExampleGetInteraction_uint_uint() {
	dataA := []uint{1, 2, 3, 4, 5}
	dataB := []uint{2, 4, 6, 7}
	temp, err := GetInteraction(dataA, dataB, "")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("temp: ", temp)
	return
}

func ExampleGetInteraction_float32_float32() {
	dataA := []float32{1, 2, 3, 4, 5}
	dataB := []float32{2, 4, 6, 7}
	temp, err := GetInteraction(dataA, dataB, "")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("temp: ", temp)
	return
}

func ExampleGetInteraction_string_string() {
	str1, str2, str3, str4, str5 := "1", "2", "3", "4", "5"
	dataA := []*string{&str1, &str2, &str3}
	dataB := []*string{&str2, &str3, &str4, &str5}
	temp, err := GetInteraction(dataA, dataB, "")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("temp: ", temp)
}

func ExampleGetInteraction_struct_struct() {
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
	if value, err := GetInteraction(studentA, studentB, "Age"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("studentA U studentB, result: ", value)
	}
}

// difference
func ExampleGetDifference_int_int() {
	dataA := []int{1, 2, 3, 4, 5}
	dataB := []int{2, 4, 6, 7}
	temp, err := GetDifference(dataA, dataB, "")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("temp: ", temp)
	return
}

func ExampleGetDifference_uint_uint() {
	dataA := []uint{1, 2, 3, 4, 5}
	dataB := []uint{2, 4, 6, 7}
	temp, err := GetDifference(dataA, dataB, "")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("temp: ", temp)
	return
}

func ExampleGetDifference_float32_float32() {
	dataA := []float32{1, 2, 3, 4, 5}
	dataB := []float32{2, 4, 6, 7}
	temp, err := GetDifference(dataA, dataB, "")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("temp: ", temp)
	return
}

func ExampleGetDifference_string_string() {
	str1, str2, str3, str4, str5 := "1", "2", "3", "4", "5"
	dataA := []*string{&str1, &str2, &str3}
	dataB := []*string{&str2, &str3, &str4, &str5}
	temp, err := GetDifference(dataA, dataB, "")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("temp: ", temp)
}

func ExampleGetDifference_struct_struct() {
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
	if value, err := GetDifference(studentA, studentB, "Age"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("studentA U studentB, result: ", value)
	}
}
