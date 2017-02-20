#  package slicelement
slicelement is a Go library for finding the input element whether exists in data of composite slice type, while providing helpful error handling. now it supports:

1. the complex type contains `[]struct/[]*struct`, `[]int/[]*int`, `[]float/[]*float`, `[]string/[]*string`, `[]byte/[]*byte`
2. the element type contains `struct`, `int`, `string`, `float`, etc

## Installation

Standard  `go get`:

```
    $  go get github.com/1046102779/slicelement
```

## Usage & Example

A quick code example is shown below:

### example 1: []struct  
```go 
data := []*Person{
    &Student{
        Name:     "John",
        Age:      29,
        Children: []string{"David", "Lily", "Bruce Lee"},
    },
    &Student{
        Name:     "Joe",
        Age:      18,
        Children: []string{},
    },
}
elem := 18
isExist, err := slicelement.Contains(data, elem, "Age")
if err != nil {
    fmt.Println(err.Error())
}
if isExist {
    fmt.Println("elem in data is exist")
} else {
    fmt.Println("elem in data is not exist")
}
// output: elem in data is exist
```

### example 2: []string

```go
var (
    data []int = []int{1, 2, 3, 4, 5}
    elem int = 4
)
isExist, err := slicelement.Contains(data, elem, "")
if err !=nil{
    fmt.Println(err.Error())
}
if isExist {
    fmt.Println("elem in data is exist")
} else {
    fmt.Println("elem in data is not exist")
}
// output: elem in data is exist
```
