package models

import (
	"encoding/json"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func CreatePayload(t *testing.T, jsonStr string) *Payload {
	var payload Payload
	err := json.Unmarshal([]byte(jsonStr), &payload)

	assert.NoError(t, err)
	return &payload
}

func CreateMulPayload(t *testing.T, jsonStr string) *MulPayload {
	var payload MulPayload
	err := json.Unmarshal([]byte(jsonStr), &payload)

	assert.NoError(t, err)
	return &payload
}

func CreateValidator() *validator.Validate {
	v := validator.New()
	Register(v)

	return v
}

func TestEnumValid(t *testing.T) {
	jsonStr := `{"color":"green"}`

	payload := CreatePayload(t, jsonStr)

	validate := CreateValidator()

	err := validate.Struct(payload)
	assert.NoError(t, err)
}

func TestEnumInvalid(t *testing.T) {
	jsonStr := `{"color":"yellow"}`

	payload := CreatePayload(t, jsonStr)

	validate := CreateValidator()

	err := validate.Struct(payload)
	if assert.Error(t, err) {
		errs := err.(validator.ValidationErrors)

		assert.Len(t, errs, 1)
		assert.Equal(t, "Color", errs[0].Field())
		assert.Equal(t, "enum", errs[0].Tag())
	}
}

func TestEnumOtherMissing(t *testing.T) {
	jsonStr := `{"color":"other","other":""}`

	payload := CreatePayload(t, jsonStr)

	validate := CreateValidator()

	err := validate.Struct(payload)
	if assert.Error(t, err) {
		errs := err.(validator.ValidationErrors)

		assert.Len(t, errs, 1)
		assert.Equal(t, "Other", errs[0].Field())
		assert.Equal(t, "required_if", errs[0].Tag())
	}
}

func TestEnumOtherExist(t *testing.T) {
	jsonStr := `{"color":"other","other":"yellow"}`

	payload := CreatePayload(t, jsonStr)

	validate := CreateValidator()

	err := validate.Struct(payload)
	assert.NoError(t, err)
}

func TestMulEnumValidMissing(t *testing.T) {
	jsonStr := `{"colors":["green","red","other"],"other":""}`

	payload := CreateMulPayload(t, jsonStr)

	validate := CreateValidator()

	err := validate.Struct(payload)

	if assert.Error(t, err) {
		errs := err.(validator.ValidationErrors)

		assert.Len(t, errs, 1)
		assert.Equal(t, "Other", errs[0].Field())
		assert.Equal(t, "required_if_element", errs[0].Tag())
	}
}

func TestMulEnumValidExist(t *testing.T) {
	jsonStr := `{"colors":["green","red","other"],"other":"yellow"}`

	payload := CreateMulPayload(t, jsonStr)

	validate := CreateValidator()

	err := validate.Struct(payload)

	assert.NoError(t, err)
}
