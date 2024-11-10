package migrator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestMigrator(t *testing.T) {
	db := &gorm.DB{}

	migrate := AutoMigrate(db)
	assert.NotEmpty(t, migrate)
}
