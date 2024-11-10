package commands

import (
	"fmt"
	"tauth/constants"
	"tauth/database"
	"tauth/database/seeder"

	"github.com/spf13/cobra"
)

func Seed() *cobra.Command {
	var err error
	return &cobra.Command{
		Use: constants.Command.SEED,
		RunE: func(cmd *cobra.Command, args []string) error {
			dbConnection, sqlConnection := database.Connection()
			defer sqlConnection.Close()
			begin := dbConnection.Begin()

			for i, seed := range seeder.Seed(begin) {
				if err = seed.Run(begin); err != nil {
					begin.Rollback()
					fmt.Println("[Seeder] Running seed failed")
					panic(err)
				}
				fmt.Println("[", i, "]: ", "Seed table: ", seed.TableName)
			}
			begin.Commit()
			fmt.Println("Seeding Completed")
			return nil
		},
	}
}
