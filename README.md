#  package slicelement
slicelement is a Go library for finding the input element whether exists in data of composite slice type, while providing helpful error handling. now it supports:

1. the complex type contains `[]struct/[]*struct`, `[]int/[]*int`, `[]float/[]*float`, `[]string/[]*string`, `[]byte/[]*byte`
2. the element type contains `struct`, `int`, `string`, `float`, etc

## Installation

Standard  `go get`:

```
    $  go get -v -u github.com/1046102779/slicelement
```

## Index

```go
//  find the element whether exists in data, if exist, return true, nil 
func Contains(data interface{}, elem interface{}, tag string) (bool, error)

// get the element index in data, if not exist, return -1, nil. 
func GetIndex(data interface{}, elem interface{}, tag string) (int, error)

desc: if the data's type is not `[]*struct/[]struct`, the `tag` value is empty
```


## Usage & Example

three quick code examples are shown below:

### example 1: []struct  
```go 
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

###  example 3:  []struct  
```go
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
```
