package migrator

import (
	"tauth/config"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestMigrator(t *testing.T) {
	config.LoadConfigs()
	db := &gorm.DB{}

	migrate := AutoMigrate(db)
	assert.NotEmpty(t, migrate)
}
