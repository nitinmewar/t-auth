package seeder

import (
	"tauth/database/seeds"

	"gorm.io/gorm"
)

type seed struct {
	TableName string
	Run       func(*gorm.DB) error
}

func Seed(db *gorm.DB) []seed {
	users := seed{TableName: "users", Run: func(d *gorm.DB) error { return seeds.Users(db) }}
	return []seed{
		users,
	}
}
