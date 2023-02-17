package models

import (
	"encoding/json"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func CreatePayload[T any](jsonStr string) *T {
	var payload T

	if err := json.Unmarshal([]byte(jsonStr), &payload); err != nil {
		panic(err)
	}

	return &payload
}

func CreateValidator() *validator.Validate {
	v := validator.New()
	Register(v)

	return v
}

func TestEnumValid(t *testing.T) {
	jsonStr := `{"color":"green"}`

	payload := CreatePayload[Payload](jsonStr)

	validate := CreateValidator()

	err := validate.Struct(payload)
	assert.NoError(t, err)
}

func TestEnumInvalid(t *testing.T) {
	jsonStr := `{"color":"yellow"}`

	payload := CreatePayload[Payload](jsonStr)

	validate := CreateValidator()

	err := validate.Struct(payload)
	if assert.Error(t, err) {
		errs := err.(validator.ValidationErrors)

		assert.Len(t, errs, 1)
		assert.Equal(t, "color", errs[0].Field())
		assert.Equal(t, "enum", errs[0].Tag())
	}
}

func TestEnumOtherMissing(t *testing.T) {
	jsonStr := `{"color":"other","other":""}`

	payload := CreatePayload[Payload](jsonStr)

	validate := CreateValidator()

	err := validate.Struct(payload)
	if assert.Error(t, err) {
		errs := err.(validator.ValidationErrors)

		assert.Len(t, errs, 1)
		assert.Equal(t, "other", errs[0].Field())
		assert.Equal(t, "required_if", errs[0].Tag())
	}
}

func TestEnumOtherExist(t *testing.T) {
	jsonStr := `{"color":"other","other":"yellow"}`

	payload := CreatePayload[Payload](jsonStr)

	validate := CreateValidator()

	err := validate.Struct(payload)
	assert.NoError(t, err)
}

func TestMulEnumValidMissing(t *testing.T) {
	jsonStr := `{"colors":["green","red","other"],"other":""}`

	payload := CreatePayload[MulPayload](jsonStr)

	validate := CreateValidator()

	err := validate.Struct(payload)

	if assert.Error(t, err) {
		errs := err.(validator.ValidationErrors)

		assert.Len(t, errs, 1)
		assert.Equal(t, "other", errs[0].Field())
		assert.Equal(t, "required_if_element", errs[0].Tag())
	}
}

func TestMulEnumValidExist(t *testing.T) {
	jsonStr := `{"colors":["green","red","other"],"other":"yellow"}`

	payload := CreatePayload[MulPayload](jsonStr)

	validate := CreateValidator()

	err := validate.Struct(payload)

	assert.NoError(t, err)
}
