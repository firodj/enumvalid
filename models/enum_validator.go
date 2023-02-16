package models

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type EnumValid interface {
	Valid() bool
}

func Register(v *validator.Validate) {
	v.RegisterValidation("enum", ValidateEnum)
	v.RegisterValidation("required_if_element", ValidateRequiredIfElement)
}

func ValidateEnum(fl validator.FieldLevel) bool {
	if enum, ok := fl.Field().Interface().(EnumValid); ok {
		return enum.Valid()
	}
	return false
}

func ValidateRequiredIfElement(fl validator.FieldLevel) bool {
	params := strings.SplitN(fl.Param(), " ", 2)

	if !requireCheckFieldElem(fl.Parent(), params[0], params[1]) {
		return true
	}

	return hasValue(fl)
}

func requireCheckFieldElem(val reflect.Value, name string, value string) bool {
	field := val.FieldByName(name)

	required := false
	if field.Kind() == reflect.Slice {
		for i := 0; i < field.Len(); i++ {
			if field.Index(i).String() == value {
				required = true
				break
			}
		}
	} else {
		panic(fmt.Sprintf("Bad field type %T", field.Interface()))
	}
	return required
}

func hasValue(fl validator.FieldLevel) bool {
	field := fl.Field()
	switch field.Kind() {
	case reflect.Slice, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Chan, reflect.Func:
		return !field.IsNil()
	default:
		return field.IsValid() && field.Interface() != reflect.Zero(field.Type()).Interface()
	}
}
