package slicelement

import (
	"reflect"
	"testing"
)

func TestIndex_getIndexInt(t *testing.T) {
	type args struct {
		Data    interface{}
		Element interface{}
	}
	tests := []struct {
		Name      string
		T         *Index
		Args      args
		WantIndex int
		WantErr   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			index := &Index{}
			gotIndex, err := index.getIndexInt(tt.Args.Data, tt.Args.Element)
			if (err != nil) != tt.WantErr {
				t.Errorf("Index.getIndexInt() error = %v, WantErr %v", err, tt.WantErr)
				return
			}
			if gotIndex != tt.WantIndex {
				t.Errorf("Index.getIndexInt() = %v, Want %v", gotIndex, tt.WantIndex)
			}
		})
	}
}

func TestIndex_getIndexString(t *testing.T) {
	type args struct {
		Data    interface{}
		Element interface{}
	}
	tests := []struct {
		Name      string
		T         *Index
		Args      args
		WantIndex int
		WantErr   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			index := &Index{}
			gotIndex, err := index.getIndexString(tt.Args.Data, tt.Args.Element)
			if (err != nil) != tt.WantErr {
				t.Errorf("Index.getIndexString() error = %v, WantErr %v", err, tt.WantErr)
				return
			}
			if gotIndex != tt.WantIndex {
				t.Errorf("Index.getIndexString() = %v, Want %v", gotIndex, tt.WantIndex)
			}
		})
	}
}

func TestIndex_getIndexFloat32(t *testing.T) {
	type args struct {
		Data    interface{}
		Element interface{}
	}
	tests := []struct {
		Name      string
		T         *Index
		Args      args
		WantIndex int
		WantErr   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			index := &Index{}
			gotIndex, err := index.getIndexFloat32(tt.Args.Data, tt.Args.Element)
			if (err != nil) != tt.WantErr {
				t.Errorf("Index.getIndexFloat32() error = %v, WantErr %v", err, tt.WantErr)
				return
			}
			if gotIndex != tt.WantIndex {
				t.Errorf("Index.getIndexFloat32() = %v, Want %v", gotIndex, tt.WantIndex)
			}
		})
	}
}

func TestIndex_getIndexUint(t *testing.T) {
	type Args struct {
		Data    interface{}
		Element interface{}
	}
	tests := []struct {
		Name      string
		T         *Index
		Args      Args
		WantIndex int
		WantErr   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			index := &Index{}
			gotIndex, err := index.getIndexUint(tt.Args.Data, tt.Args.Element)
			if (err != nil) != tt.WantErr {
				t.Errorf("Index.getIndexUint() error = %v, WantErr %v", err, tt.WantErr)
				return
			}
			if gotIndex != tt.WantIndex {
				t.Errorf("Index.getIndexUint() = %v, Want %v", gotIndex, tt.WantIndex)
			}
		})
	}
}

func TestIndex_getIndexStruct(t *testing.T) {
	type Args struct {
		Data    interface{}
		Element interface{}
		Tag     string
	}
	tests := []struct {
		Name      string
		T         *Index
		Args      Args
		WantIndex int
		WantErr   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			index := &Index{}
			gotIndex, err := index.getIndexStruct(tt.Args.Data, tt.Args.Element, tt.Args.Tag)
			if (err != nil) != tt.WantErr {
				t.Errorf("Index.getIndexStruct() error = %v, WantErr %v", err, tt.WantErr)
				return
			}
			if gotIndex != tt.WantIndex {
				t.Errorf("Index.getIndexStruct() = %v, Want %v", gotIndex, tt.WantIndex)
			}
		})
	}
}

func TestIndex_decodeStruct(t *testing.T) {
	type args struct {
		DataVal reflect.Value
		Element interface{}
		Tag     string
	}
	tests := []struct {
		Name        string
		T           *Index
		Args        args
		WantIsExist bool
		WantErr     bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			index := &Index{}
			gotIsExist, err := index.decodeStruct(tt.Args.DataVal, tt.Args.Element, tt.Args.Tag)
			if (err != nil) != tt.WantErr {
				t.Errorf("Index.decodeStruct() error = %v, WantErr %v", err, tt.WantErr)
				return
			}
			if gotIsExist != tt.WantIsExist {
				t.Errorf("Index.decodeStruct() = %v, Want %v", gotIsExist, tt.WantIsExist)
			}
		})
	}
}
