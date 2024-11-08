package utils

import (
	"fmt"
	"tauth/database"
	"tauth/database/migrator"
)

func Migrate() error {
	dbConnection, _ := database.Connection()
	//defer sqlConnection.Close()

	begin := dbConnection.Begin()

	for i, migrate := range migrator.AutoMigrate(begin) {
		if err := migrate.Run(begin); err != nil {
			begin.Rollback()
			fmt.Println("[Migrate] Running raw sql schema creation failed")
			panic(err)
		}
		fmt.Println("[", i, "]: ", "Migrate table: ", migrate.TableName)
	}
	begin.Commit()
	fmt.Println("Migration Completed")
	return nil
}
