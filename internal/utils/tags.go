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

func GetTagMap(s interface{}, tagName string, reverse bool) (interface{}, error) {
	valueOf, err := checkStruct(s)
	if err != nil {
		return nil, err
	}

	if reverse {
		return makeTagMapReverse(valueOf, tagName), nil
	}

	return makeTagMap(valueOf, tagName), nil
}

func makeTagMap(valueOf reflect.Value, tagName string) map[interface{}]string {
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

	return m
}

func makeTagMapReverse(valueOf reflect.Value, tagName string) map[string]interface{} {
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

	return m
}

func checkStruct(s interface{}) (reflect.Value, error) {
	valueOf := reflect.Indirect(reflect.ValueOf(s))
	isStruct := valueOf.Kind() == reflect.Struct
	if !isStruct {
		return valueOf, errors.New("Expecting pointer to struct")
	}
	return valueOf, nil
}
