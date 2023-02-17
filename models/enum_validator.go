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
	// register function to get tag name from json tags.
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	v.RegisterValidation("enum", GetValidateEnum(v))
	v.RegisterValidation("required_if_element", GetValidateRequiredIfElement(v))
}

func GetValidateEnum(v *validator.Validate) validator.Func {
	return func(fl validator.FieldLevel) bool {
		if enum, ok := fl.Field().Interface().(EnumValid); ok {
			return enum.Valid()
		}
		return false
	}
}

func GetValidateRequiredIfElement(v *validator.Validate) validator.Func {
	return func(fl validator.FieldLevel) bool {
		params := strings.SplitN(fl.Param(), " ", 2)

		if !requireCheckFieldElem(fl.Parent(), params[0], params[1]) {
			return true
		}

		return validateRequired(v, fl.Field().Interface())
	}
}

func requireCheckFieldElem(val reflect.Value, name string, value string) bool {
	field := val.FieldByName(name)

	required := false
	switch field.Kind() {
	case reflect.Slice:
		for i := 0; i < field.Len(); i++ {
			if field.Index(i).String() == value {
				required = true
				break
			}
		}
		return required
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}

func validateRequired(v *validator.Validate, val interface{}) bool {
	return v.Var(val, "required") == nil
}
