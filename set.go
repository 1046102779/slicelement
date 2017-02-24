// add list operation, include union, interaction and difference
// 1. Union(data []interface{},  dest []interface{}, tagName string)
// 2. Interaction(data []interface{}, dest interface{}, tagName string)
// 3. Difference(data []interface{}, dest interface{}, tagName string)
package slicelement

import (
	"reflect"
	"strings"

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
	case reflect.Int, reflect.Float32, reflect.String:
		result, err = union.getNonStruct(dataA, dataB)
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
	case reflect.Int, reflect.Float32, reflect.String:
		result, err = interaction.getNonStruct(dataA, dataB)
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
	case reflect.Int, reflect.Float32, reflect.String:
		result, err = diff.getNonStruct(dataA, dataB)
	case reflect.Struct:
		result, err = diff.getStruct(dataA, dataB, tagName)
	}
	return
}

// union
type Union struct{}

func (t *Union) getNonStruct(dataA interface{}, dataB interface{}) (result interface{}, err error) {
	valueB := reflect.ValueOf(dataB)
	resultVal := reflect.ValueOf(dataA)
	for index := 0; index < valueB.Len(); index++ {
		var subIndex int = 0
		for subIndex = 0; subIndex < resultVal.Len(); subIndex++ {
			resultElem := reflect.Indirect(resultVal.Index(subIndex))
			bElem := reflect.Indirect(valueB.Index(index))
			if resultElem.Interface() == bElem.Interface() {
				break
			}
		}
		if subIndex == resultVal.Len() {
			resultVal = reflect.Append(resultVal, valueB.Index(index))
		}
	}
	return resultVal.Interface(), nil
}

// tagName: unique key
func (t *Union) getStruct(dataA interface{}, dataB interface{}, tagName string) (result interface{}, err error) {
	if strings.TrimSpace(tagName) == "" {
		err = errors.New("slice struct's FieldName can't be empty")
		return
	}
	resultVal := reflect.ValueOf(dataA)
	// use GetIndex
	valueB := reflect.ValueOf(dataB)
	dataBFieldIndex := getStructTagIndex(valueB.Type().Elem(), tagName)
	if dataBFieldIndex < 0 {
		err = errors.New("field `" + tagName + "` not exist in struct")
		return
	}
	var isExist bool = false
	for index := 0; index < valueB.Len(); index++ {
		underlyFieldValueB := reflect.Indirect(valueB.Index(index).Field(dataBFieldIndex))
		if isExist, err = Contains(dataA, underlyFieldValueB.Interface(), tagName); err != nil {
			err = errors.Wrap(err, "getStruct")
			return
		} else if !isExist {
			resultVal = reflect.Append(resultVal, valueB.Index(index))
		}
	}
	return resultVal.Interface(), nil
}

// get index of the field name in  struct data
func getStructTagIndex(typ reflect.Type, tagName string) int {
	for index := 0; index < typ.NumField(); index++ {
		if typ.Field(index).Name == tagName {
			return index
		}
	}
	return -1
}

// interaction
type Interaction struct{}

func (t *Interaction) getNonStruct(dataA interface{}, dataB interface{}) (result interface{}, err error) {
	valueB := reflect.ValueOf(dataB)
	// new zero value
	resultVal := reflect.MakeSlice(reflect.ValueOf(dataA).Type(), 0, 0)
	valueA := reflect.ValueOf(dataA)
	for index := 0; index < valueB.Len(); index++ {
		var subIndex int = 0
		for subIndex = 0; subIndex < valueA.Len(); subIndex++ {
			aElem := reflect.Indirect(valueA.Index(subIndex))
			bElem := reflect.Indirect(valueB.Index(index))
			if aElem.Interface() == bElem.Interface() {
				resultVal = reflect.Append(resultVal, valueB.Index(index))
				break
			}
		}
	}
	return resultVal.Interface(), nil
}

func (t *Interaction) getStruct(dataA interface{}, dataB interface{}, tagName string) (result interface{}, err error) {
	if strings.TrimSpace(tagName) == "" {
		err = errors.New("slice struct's FieldName can't be empty")
		return
	}
	// new zero value
	resultVal := reflect.MakeSlice(reflect.ValueOf(dataA).Type(), 0, 0)
	// use GetIndex
	valueB := reflect.ValueOf(dataB)
	dataBFieldIndex := getStructTagIndex(valueB.Type().Elem(), tagName)
	if dataBFieldIndex < 0 {
		err = errors.New("field `" + tagName + "` not exist in struct")
		return
	}
	var isExist bool = false
	for index := 0; index < valueB.Len(); index++ {
		underlyFieldValueB := reflect.Indirect(valueB.Index(index).Field(dataBFieldIndex))
		if isExist, err = Contains(dataA, underlyFieldValueB.Interface(), tagName); err != nil {
			err = errors.Wrap(err, "getStruct")
			return
		} else if isExist {
			resultVal = reflect.Append(resultVal, valueB.Index(index))
		}
	}
	return resultVal.Interface(), nil
}

type Difference struct{}

func (t *Difference) getNonStruct(dataA interface{}, dataB interface{}) (result interface{}, err error) {
	valueB := reflect.ValueOf(dataB)
	// new zero value
	resultVal := reflect.MakeSlice(reflect.ValueOf(dataA).Type(), 0, 0)
	valueA := reflect.ValueOf(dataA)
	for index := 0; index < valueA.Len(); index++ {
		var subIndex int = 0
		for subIndex = 0; subIndex < valueB.Len(); subIndex++ {
			aElem := reflect.Indirect(valueA.Index(index))
			bElem := reflect.Indirect(valueB.Index(subIndex))
			if aElem.Interface() == bElem.Interface() {
				break
			}
		}
		if subIndex == valueB.Len() {
			resultVal = reflect.Append(resultVal, valueA.Index(index))
		}
	}
	return resultVal.Interface(), nil
}

func (t *Difference) getStruct(dataA interface{}, dataB interface{}, tagName string) (result interface{}, err error) {
	if strings.TrimSpace(tagName) == "" {
		err = errors.New("slice struct's FieldName can't be empty")
		return
	}
	// new zero value
	resultVal := reflect.MakeSlice(reflect.ValueOf(dataA).Type(), 0, 0)
	// use GetIndex
	valueA := reflect.ValueOf(dataA)
	dataAFieldIndex := getStructTagIndex(valueA.Type().Elem(), tagName)
	if dataAFieldIndex < 0 {
		err = errors.New("field `" + tagName + "` not exist in struct")
		return
	}
	var isExist bool = false
	for index := 0; index < valueA.Len(); index++ {
		underlyFieldValueA := reflect.Indirect(valueA.Index(index).Field(dataAFieldIndex))
		if isExist, err = Contains(dataB, underlyFieldValueA.Interface(), tagName); err != nil {
			err = errors.Wrap(err, "getStruct")
			return
		} else if !isExist {
			resultVal = reflect.Append(resultVal, valueA.Index(index))
		}
	}
	return resultVal.Interface(), nil
	return
}
