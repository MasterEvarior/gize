package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvVar(t *testing.T) {
	t.Setenv("TEST_VAR", "my-value")

	result := GetEnvVar("TEST_VAR")

	assert.Equal(t, "my-value", result)
}
