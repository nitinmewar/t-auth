package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {

	gormDB, sqlDB := Connection()
	assert.NotEmpty(t, gormDB)
	assert.NotEmpty(t, sqlDB)
}
