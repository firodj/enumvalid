package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	var green Color = "green"

	assert.Equal(t, "green", green.String())

	var gray Color = "gray"

	assert.Equal(t, "", gray.String())
}
