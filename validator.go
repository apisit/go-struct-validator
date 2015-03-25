package validator

import (
	"errors"
	"fmt"
	_ "log"
	"reflect"
	"strings"
)

func Validate(v interface{}) error {
	vo := reflect.ValueOf(v)
	to := reflect.TypeOf(v)

	if vo.Kind() != reflect.Struct {
		return errors.New("Parameter is not struct")
	}

	for i := 0; i < vo.NumField(); i++ {
		fieldType := to.Field(i)
		tag := fieldType.Tag.Get("validator")
		isRequired := strings.Contains(tag, "require")
		if isRequired {
			fieldValue := vo.Field(i)
			if fieldValue.Kind() == reflect.String {
				value := strings.TrimSpace(fieldValue.String())
				if len(value) == 0 {
					return errors.New(fmt.Sprintf("%s is an empty string", fieldType.Name))
				}
			} else if fieldValue.Kind() == reflect.Slice {
				if fieldValue.Len() == 0 {
					return errors.New(fmt.Sprintf("%s is an empty array", fieldType.Name))
				}
			} else if fieldValue.Kind() == reflect.Ptr && fieldValue.IsNil() {
				return errors.New(fmt.Sprintf("%s is nil", fieldType.Name))
			}
		}
	}
	return nil
}
