package commands

import (
	"fmt"
	"tauth/config"
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
			if config.App.Env != constants.Environment.LOCAL {
				fmt.Println("Warning: Environment is not local. Tables wont be seeded")
				return nil
			}
			fmt.Println("App env is local")
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
