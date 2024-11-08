package migrator

import (
	"tauth/entities"

	"gorm.io/gorm"
)

type Migrate struct {
	TableName string
	Run       func(*gorm.DB) error
}

func AutoMigrate(db *gorm.DB) []Migrate {
	var users entities.Users

	usersM := Migrate{TableName: "users",
		Run: func(d *gorm.DB) error { return db.AutoMigrate(&users) }}
	return []Migrate{
		usersM,
	}
}
