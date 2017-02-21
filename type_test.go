package slicelement

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_checkInputValid(t *testing.T) {
	type args struct {
		Data    interface{}
		Element interface{}
	}
	tests := []struct {
		Name    string
		Args    args
		WantErr bool
	}{
		// TODO: Add test cases.
		{
			Name: "Joe",
			Args: args{
				Data:    "data1",
				Element: 1,
			},
			WantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			if err := checkInputValid(tt.Args.Data, tt.Args.Element); (err != nil) != tt.WantErr {
				t.Errorf("checkInputValid() error = %v, wantErr %v", err, tt.WantErr)
			}
		})
	}
}

func Test_checkData(t *testing.T) {
	type args struct {
		Data interface{}
	}
	type Complex struct {
		Name    string
		Args    args
		WantErr bool
	}
	var (
		// TODO: Add test cases.
		tests []interface{} = []interface{}{
			[]*Complex{},
			&Complex{},
			[...]*Complex{
				&Complex{
					Name: "Joe",
					Args: args{
						Data: "hello,world",
					},
				},
			},
			[]*Complex{
				&Complex{
					Name: "David",
					Args: args{
						Data: "thank you",
					},
				},
			},
		}
	)
	for _, tt := range tests {
		if _, err := Contains(tt, "Joe", "Name"); err != nil {
			fmt.Println("checkElement() error = ", err)
		}
	}
}

func Test_checkElement(t *testing.T) {
	var (
		// TODO: Add test cases.
		tests []interface{} = []interface{}{}
		elem1 *int          = new(int)  // illegal
		elem2 *int                      // illegal
		elem3 **int         = new(*int) // illegal
	)
	*elem1 = 3
	tests = append(tests, elem1, elem2, elem3)
	for _, tt := range tests {
		if err := checkElement(tt); err != nil {
			fmt.Println("checkElement() error = ", err)
		}
	}
}

func Test_getKind(t *testing.T) {
	type args struct {
		Val reflect.Value
	}
	tests := []struct {
		Name     string
		Args     args
		WantKind reflect.Kind
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			if gotKind := getKind(tt.Args.Val); gotKind != tt.WantKind {
				t.Errorf("getKind() = %v, want %v", gotKind, tt.WantKind)
			}
		})
	}
}

func Test_getSliceUnderlyKind(t *testing.T) {
	type args struct {
		Data interface{}
	}
	tests := []struct {
		Name     string
		Args     args
		WantKind reflect.Kind
		WantErr  bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			gotKind, err := getSliceUnderlyKind(tt.Args.Data)
			if (err != nil) != tt.WantErr {
				t.Errorf("getSliceUnderlyKind() error = %v, wantErr %v", err, tt.WantErr)
				return
			}
			if gotKind != tt.WantKind {
				t.Errorf("getSliceUnderlyKind() = %v, want %v", gotKind, tt.WantKind)
			}
		})
	}
}

func TestContains(t *testing.T) {
	type args struct {
		Data    interface{}
		Element interface{}
		Tag     string
	}
	tests := []struct {
		Name        string
		Args        args
		WantIsExist bool
		WantErr     bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			gotIsExist, err := Contains(tt.Args.Data, tt.Args.Element, tt.Args.Tag)
			if (err != nil) != tt.WantErr {
				t.Errorf("Contains() error = %v, wantErr %v", err, tt.WantErr)
				return
			}
			if gotIsExist != tt.WantIsExist {
				t.Errorf("Contains() = %v, want %v", gotIsExist, tt.WantIsExist)
			}
		})
	}
}
