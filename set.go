// add list operation, include union, interaction and difference
// 1. Union(data []interface{},  dest []interface{}, tagName string)
// 2. Interaction(data []interface{}, dest interface{}, tagName string)
// 3. Difference(data []interface{}, dest interface{}, tagName string)
package slicelement

import (
	"reflect"

	"github.com/pkg/errors"
)

// check input dataA, if data is nil, it will new object
func CheckSetDataA(data interface{}) (err error) {
	value := reflect.ValueOf(data)
	if data == nil {
		// it need new object
		underlyType := value.Type().Elem()
		newValue := reflect.MakeSlice(underlyType, 0, 0)
		if value.CanSet() {
			value.Set(newValue)
		}
		return
	}
	if value.Kind() != reflect.Slice {
		err = errors.New("the first input data must be slice type")
		return
	}
	return
}

// check input dataB, if dataB is nil and check whether the data  need to be added
func CheckSetDataB(data interface{}) (needAdd bool, err error) {
	value := reflect.ValueOf(data)
	if value.Kind() != reflect.Slice && value.Kind() != reflect.Array {
		err = errors.New("the second data must be slice or array")
		return
	}
	if value.IsNil() {
		return false, nil
	}
	return true, nil
}

// check two input datas
func checkSetInputData(dataA interface{}, dataB interface{}) (needAdd bool, err error) {
	// check dataA
	if err = CheckSetDataA(dataA); err != nil {
		err = errors.Wrap(err, "checkSetInputData")
		return
	}
	// check dataB
	if needAdd, err = CheckSetDataB(dataB); err != nil {
		err = errors.Wrap(err, "checkSetInputData")
		return false, err
	} else if !needAdd {
		return false, nil
	}
	// if dataA and dataB is not the same, return err
	underlyTypeA, err := getSliceUnderlyKind(dataA)
	if err != nil {
		err = errors.Wrap(err, "checkSetInputData")
		return
	}
	underlyTypeB, err := getSliceUnderlyKind(dataB)
	if err != nil {
		err = errors.Wrap(err, "checkSetInputData")
		return
	}
	if underlyTypeA != underlyTypeB {
		err = errors.New("input datas type are not the same")
		return
	}
	return
}

// union operation
func GetUnion(dataA interface{}, dataB interface{}, tagName string) (result interface{}, err error) {
	var needAdd bool = false
	if needAdd, err = checkSetInputData(dataA, dataB); err != nil {
		err = errors.Wrap(err, "Union")
		return
	} else if !needAdd {
		return dataA, nil
	}
	// struct type
	// not struct type, include string, int, float32 etc.
	return
}

// interaction operation
func GetInteraction(dataA interface{}, dataB interface{}, tagName string) (result interface{}, err error) {
	return
}

// difference operation
func GetDifference(dataA interface{}, dataB interface{}, tagName string) (result interface{}, err error) {
	return
}

// union
type Union struct{}

func (t *Union) getString(dataA interface{}, dataB interface{}) (result interface{}, err error) {
	return
}

func (t *Union) getInt(dataA interface{}, dataB interface{}) (result interface{}, err error) {
	return
}

func (t *Union) getFloat32(dataA interface{}, dataB interface{}) (result interface{}, err error) {
	return
}

func (t *Union) getStruct(dataA interface{}, dataB interface{}, tagName string) (result interface{}, err error) {
	return
}

// interaction
type Interaction struct{}

func (t *Interaction) getString(dataA interface{}, dataB interface{}) (result interface{}, err error) {
	return
}

func (t *Interaction) getInt(dataA interface{}, dataB interface{}) (result interface{}, err error) {
	return
}

func (t *Interaction) getFloat32(dataA interface{}, dataB interface{}) (result interface{}, err error) {
	return
}

func (t *Interaction) getStruct(dataA interface{}, dataB interface{}, tagName string) (result interface{}, err error) {
	return
}

type Difference struct{}

func (t *Difference) getString(dataA interface{}, dataB interface{}) (result interface{}, err error) {
	return
}

func (t *Difference) getInt(dataA interface{}, dataB interface{}) (result interface{}, err error) {
	return
}

func (t *Difference) getFloat32(dataA interface{}, dataB interface{}) (result interface{}, err error) {
	return
}

func (t *Difference) getStruct(dataA interface{}, dataB interface{}, tagName string) (result interface{}, err error) {
	return
}
