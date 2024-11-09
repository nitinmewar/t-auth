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
	var keyStrokes entities.KeystrokeProfile

	usersM := Migrate{TableName: "users",
		Run: func(d *gorm.DB) error { return db.AutoMigrate(&users) }}
	keyStrokeM := Migrate{TableName: "keystroke_profile",
		Run: func(d *gorm.DB) error { return db.AutoMigrate(&keyStrokes) }}
	return []Migrate{
		usersM,
		keyStrokeM,
	}
}
