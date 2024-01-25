package helper

import (
	"fmt"
	"reflect"
)

var cutter map[string]string = map[string]string{
	"d": "decode",
	"e": "encode",
}

func GetFiledFromStruct(structure interface{}, fieldName string) (interface{}, error) {
	valueOf := reflect.ValueOf(structure)
	fieldValue := reflect.Indirect(valueOf).FieldByName(fieldName)
	if !fieldValue.IsValid() {
		curField := cutter[fieldName]
		fieldValue = reflect.Indirect(valueOf).FieldByName(FirstLettertoUpper(curField))

		if !fieldValue.IsValid() {
			return nil, fmt.Errorf("Поле %s не найдено", fieldName)
		}
	}

	return fieldValue.Interface(), nil
}

func GetStringFieldFromStruct(structure interface{}, fieldName string) (string, error) {
	valueOf := reflect.ValueOf(structure)
	fieldValue := reflect.Indirect(valueOf).FieldByName(fieldName)
	if !fieldValue.IsValid() {
		return "", fmt.Errorf("Поле %s не найдено", fieldName)
	}

	return fieldValue.String(), nil
}
