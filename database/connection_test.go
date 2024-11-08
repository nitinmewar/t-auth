package database

import (
	"tauth/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	config.LoadConfigs()

	gormDB, sqlDB := Connection()
	assert.NotEmpty(t, gormDB)
	assert.NotEmpty(t, sqlDB)
}
