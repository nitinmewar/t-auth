package seeder

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestSeeder(t *testing.T) {
	db := &gorm.DB{}

	seeds := Seed(db)
	assert.NotEmpty(t, seeds)
}
