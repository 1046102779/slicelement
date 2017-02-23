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
	union := &Union{}
	kindA, err := getSliceUnderlyKind(dataA)
	if err != nil {
		err = errors.Wrap(err, "GetUnion")
		return
	}
	kindA = getKindByKind(kindA)
	switch kindA {
	case reflect.Int:
		result, err = union.getInt(dataA, dataB)
	case reflect.Float32:
		result, err = union.getFloat32(dataA, dataB)
	case reflect.String:
		result, err = union.getString(dataA, dataB)
	case reflect.Struct:
		result, err = union.getStruct(dataA, dataB, tagName)
	}
	return
}

// interaction operation
func GetInteraction(dataA interface{}, dataB interface{}, tagName string) (result interface{}, err error) {
	var needAdd bool = false
	if needAdd, err = checkSetInputData(dataA, dataB); err != nil {
		err = errors.Wrap(err, "Union")
		return
	} else if !needAdd {
		return dataA, nil
	}
	interaction := &Interaction{}
	kindA, err := getSliceUnderlyKind(dataA)
	if err != nil {
		err = errors.Wrap(err, "GetInteraction")
		return
	}
	kindA = getKindByKind(kindA)
	switch kindA {
	case reflect.Int:
		result, err = interaction.getInt(dataA, dataB)
	case reflect.Float32:
		result, err = interaction.getFloat32(dataA, dataB)
	case reflect.String:
		result, err = interaction.getString(dataA, dataB)
	case reflect.Struct:
		result, err = interaction.getStruct(dataA, dataB, tagName)
	}
	return
}

// difference operation
func GetDifference(dataA interface{}, dataB interface{}, tagName string) (result interface{}, err error) {
	var needAdd bool = false
	if needAdd, err = checkSetInputData(dataA, dataB); err != nil {
		err = errors.Wrap(err, "Union")
		return
	} else if !needAdd {
		return dataA, nil
	}
	diff := &Difference{}
	kindA, err := getSliceUnderlyKind(dataA)
	if err != nil {
		err = errors.Wrap(err, "GetDifference")
		return
	}
	kindA = getKindByKind(kindA)
	switch kindA {
	case reflect.Int:
		result, err = diff.getInt(dataA, dataB)
	case reflect.Float32:
		result, err = diff.getFloat32(dataA, dataB)
	case reflect.String:
		result, err = diff.getString(dataA, dataB)
	case reflect.Struct:
		result, err = diff.getStruct(dataA, dataB, tagName)
	}
	return
}

// union
type Union struct{}

func (t *Union) getString(dataA interface{}, dataB interface{}) (result interface{}, err error) {
	valueB := reflect.ValueOf(dataB)
	result = dataA
	resValue := reflect.ValueOf(result)
	kind := reflect.ValueOf(dataA).Type().Elem().Kind()
	for index := 0; index < valueB.Len(); index++ {
		var subIndex int = 0
		for subIndex = 0; subIndex < resValue.Len(); subIndex++ {
			resElem := reflect.Indirect(resValue.Index(subIndex))
			bElem := reflect.Indirect(valueB.Index(index))
			if resElem.Interface() == bElem.Interface() {
				break
			}
		}
		if subIndex == resValue.Len() {
			if kind != reflect.Ptr {
				result = append(result.([]string), reflect.Indirect(valueB.Index(index)).Interface().(string))
			} else {
				result = append(result.([]*string), valueB.Index(index).Interface().(*string))
			}
		}
	}
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
