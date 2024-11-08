package seeds

import (
	"database/sql"
	"tauth/entities"
	"time"

	"gorm.io/gorm"
)

func Users(db *gorm.DB) error {
	// Seed 1
	err := db.Create(&entities.Users{
		PID:          sql.NullString{String: "user_1", Valid: true},
		FirstName:    "nitin",
		LastName:     "mewar",
		Password:     "$2a$12$hX2GqiSl/HLMq4efnJlfl.aH4ndpBQ8sTBQMVntwq7Uyxzq/Akd2.", //27111995
		PrimaryEmail: "nitin@gmail.com",
		IsDeleted:    false,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}).Error

	if err != nil {
		return err
	}

	return err
}
