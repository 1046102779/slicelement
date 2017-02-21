package slicelement

import (
	"reflect"
	"testing"
)

func TestContain_isContainString(t *testing.T) {
	type Args struct {
		Data    interface{}
		Element interface{}
	}
	tests := []struct {
		Name        string
		t           *Contain
		Args        Args
		WantIsExist bool
		WantErr     bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			contain := &Contain{}
			gotIsExist, err := contain.isContainString(tt.Args.Data, tt.Args.Element)
			if (err != nil) != tt.WantErr {
				t.Errorf("Contain.isContainString() error = %v, WantErr %v", err, tt.WantErr)
				return
			}
			if gotIsExist != tt.WantIsExist {
				t.Errorf("Contain.isContainString() = %v, want %v", gotIsExist, tt.WantIsExist)
			}
		})
	}
}

func TestContain_isContainInt(t *testing.T) {
	type Args struct {
		Data    interface{}
		Element interface{}
	}
	tests := []struct {
		Name        string
		t           *Contain
		Args        Args
		WantIsExist bool
		WantErr     bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			contain := &Contain{}
			gotIsExist, err := contain.isContainInt(tt.Args.Data, tt.Args.Element)
			if (err != nil) != tt.WantErr {
				t.Errorf("Contain.isContainInt() error = %v, WantErr %v", err, tt.WantErr)
				return
			}
			if gotIsExist != tt.WantIsExist {
				t.Errorf("Contain.isContainInt() = %v, want %v", gotIsExist, tt.WantIsExist)
			}
		})
	}
}

func TestContain_isContainFloat32(t *testing.T) {
	type Args struct {
		Data    interface{}
		Element interface{}
	}
	tests := []struct {
		Name        string
		t           *Contain
		Args        Args
		WantIsExist bool
		WantErr     bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			contain := &Contain{}
			gotIsExist, err := contain.isContainFloat32(tt.Args.Data, tt.Args.Element)
			if (err != nil) != tt.WantErr {
				t.Errorf("Contain.isContainFloat32() error = %v, WantErr %v", err, tt.WantErr)
				return
			}
			if gotIsExist != tt.WantIsExist {
				t.Errorf("Contain.isContainFloat32() = %v, want %v", gotIsExist, tt.WantIsExist)
			}
		})
	}
}

func TestContain_isContainUint(t *testing.T) {
	type Args struct {
		Data    interface{}
		Element interface{}
	}
	tests := []struct {
		Name        string
		t           *Contain
		Args        Args
		WantIsExist bool
		WantErr     bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			contain := &Contain{}
			gotIsExist, err := contain.isContainUint(tt.Args.Data, tt.Args.Element)
			if (err != nil) != tt.WantErr {
				t.Errorf("Contain.isContainUint() error = %v, WantErr %v", err, tt.WantErr)
				return
			}
			if gotIsExist != tt.WantIsExist {
				t.Errorf("Contain.isContainUint() = %v, want %v", gotIsExist, tt.WantIsExist)
			}
		})
	}
}

func TestContain_isContainStructs(t *testing.T) {
	type Args struct {
		Data    interface{}
		Element interface{}
		tag     string
	}
	tests := []struct {
		Name        string
		t           *Contain
		Args        Args
		WantIsExist bool
		WantErr     bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			contain := &Contain{}
			gotIsExist, err := contain.isContainStructs(tt.Args.Data, tt.Args.Element, tt.Args.tag)
			if (err != nil) != tt.WantErr {
				t.Errorf("Contain.isContainStructs() error = %v, WantErr %v", err, tt.WantErr)
				return
			}
			if gotIsExist != tt.WantIsExist {
				t.Errorf("Contain.isContainStructs() = %v, want %v", gotIsExist, tt.WantIsExist)
			}
		})
	}
}

func TestContain_decodeStruct(t *testing.T) {
	type Args struct {
		DataVal reflect.Value
		Element interface{}
		tag     string
	}
	tests := []struct {
		Name        string
		t           *Contain
		Args        Args
		WantIsExist bool
		WantErr     bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			contain := &Contain{}
			gotIsExist, err := contain.decodeStruct(tt.Args.DataVal, tt.Args.Element, tt.Args.tag)
			if (err != nil) != tt.WantErr {
				t.Errorf("Contain.decodeStruct() error = %v, WantErr %v", err, tt.WantErr)
				return
			}
			if gotIsExist != tt.WantIsExist {
				t.Errorf("Contain.decodeStruct() = %v, want %v", gotIsExist, tt.WantIsExist)
			}
		})
	}
}
