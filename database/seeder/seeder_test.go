package seeder

import (
	"tauth/config"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestSeeder(t *testing.T) {
	config.LoadConfigs()
	db := &gorm.DB{}

	seeds := Seed(db)
	assert.NotEmpty(t, seeds)
}
