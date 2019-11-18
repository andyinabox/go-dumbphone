package utils

import (
	"errors"
	"reflect"
)

func GetTag(s interface{}, fieldName string, tagName string) (string, error) {

	valueOf, err := checkStruct(s)
	if err != nil {
		return "", err
	}

	field, ok := valueOf.Type().FieldByName(fieldName)
	if !ok {
		return "", nil
	}

	return field.Tag.Get(tagName), nil
}

// func GetFieldValueByName(s interface{}, fieldName string) (interface{}, error) {
// 	valueOf, err := checkStruct(s)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return valueOf.FieldByName(fieldName).Interface(), nil
// }

func GetTagMap(s interface{}, tagName string) (map[interface{}]string, error) {
	valueOf, err := checkStruct(s)
	if err != nil {
		return nil, err
	}

	m := make(map[interface{}]string)
	fieldIndex := []int{0}
	var tagValue string
	var fieldValue interface{}

	for i := 0; i < valueOf.NumField(); i++ {
		fieldIndex[0] = i
		tagValue = valueOf.Type().FieldByIndex(fieldIndex).Tag.Get(tagName)
		fieldValue = valueOf.Field(i).Interface()

		m[fieldValue] = tagValue
	}

	return m, nil
}

func GetTagMapReverse(s interface{}, tagName string) (map[string]interface{}, error) {
	valueOf, err := checkStruct(s)
	if err != nil {
		return nil, err
	}

	m := make(map[string]interface{})
	fieldIndex := []int{0}
	var tagValue string
	var fieldValue interface{}

	for i := 0; i < valueOf.NumField(); i++ {
		fieldIndex[0] = i
		tagValue = valueOf.Type().FieldByIndex(fieldIndex).Tag.Get(tagName)
		fieldValue = valueOf.Field(i).Interface()

		m[tagValue] = fieldValue
	}

	return m, nil
}

func checkStruct(s interface{}) (reflect.Value, error) {
	valueOf := reflect.Indirect(reflect.ValueOf(s))
	isStruct := valueOf.Kind() == reflect.Struct
	if !isStruct {
		return valueOf, errors.New("Expecting pointer to struct")
	}
	return valueOf, nil
}
