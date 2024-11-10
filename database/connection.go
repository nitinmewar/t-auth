package database

import (
	"database/sql"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	faqserviceDB *gorm.DB
	sqlDB        *sql.DB
)

func Connection() (*gorm.DB, *sql.DB) {
	var err error
	if faqserviceDB != nil && sqlDB != nil {
		return faqserviceDB, sqlDB
	}
	dsn := os.Getenv("DB_URL")

	faqserviceDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("[Connection], Error in opening db")
	}

	sqlDB, err = faqserviceDB.DB()
	if err != nil {
		log.Fatal("[Connection], Error in setting sqldb")
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return faqserviceDB, sqlDB
}
