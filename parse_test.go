package slicelement

import (
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
	tests := []struct {
		Name    string
		Args    args
		WantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			if err := checkData(tt.Args.Data); (err != nil) != tt.WantErr {
				t.Errorf("checkData() error = %v, wantErr %v", err, tt.WantErr)
			}
		})
	}
}

func Test_checkElement(t *testing.T) {
	type args struct {
		Element interface{}
	}
	tests := []struct {
		Name    string
		Args    args
		WantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			if err := checkElement(tt.Args.Element); (err != nil) != tt.WantErr {
				t.Errorf("checkElement() error = %v, wantErr %v", err, tt.WantErr)
			}
		})
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
		wantErr     bool
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
