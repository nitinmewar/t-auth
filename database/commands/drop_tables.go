package commands

import (
	"errors"
	"fmt"
	"tauth/config"
	"tauth/constants"
	"tauth/database"

	"github.com/spf13/cobra"
)

func DropTables() *cobra.Command {
	return &cobra.Command{
		Use: constants.Command.DROP_TABLES,
		RunE: func(cmd *cobra.Command, args []string) error {
			if config.App.Env != constants.Environment.LOCAL {
				fmt.Println("Warning: Environment is not local. Tables wont be dropped")
				return nil
			}
			fmt.Println("App env is local")

			dbConnection, sqlConnection := database.Connection()
			defer sqlConnection.Close()

			var tableNames []string
			if err := dbConnection.Table("information_schema.tables").
				Where("table_schema = ?", "public").Pluck("table_name", &tableNames).Error; err != nil {
				panic(err)
			}

			if len(tableNames) > 0 {
				for i, tableName := range tableNames {
					if err := dbConnection.Migrator().DropTable(tableName); err != nil {
						return errors.New("Error: While dropping tables:" + tableName)
					}
					fmt.Println("[", i, "]: ", "dropped table: ", tableName)
				}
			}
			fmt.Println("Dropped all tables sucessfully")
			return nil
		},
	}
}
