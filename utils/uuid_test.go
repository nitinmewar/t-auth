package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UUID(t *testing.T) {
	uuid := UUID()
	assert.NotEmpty(t, uuid)
}
