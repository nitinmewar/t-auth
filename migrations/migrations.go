package main

import (
	"tauth/config"
	"tauth/database/commands"

	"github.com/spf13/cobra"
)

func main() {
	config.LoadConfigs()
	cmd := &cobra.Command{}

	cmd.AddCommand(commands.DropTables())
	cmd.AddCommand(commands.Migrate())
	cmd.AddCommand(commands.Seed())
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}

}
