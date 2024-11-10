package main

import (
	"tauth/database/commands"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{}

	cmd.AddCommand(commands.DropTables())
	cmd.AddCommand(commands.Migrate())
	cmd.AddCommand(commands.Seed())
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}

}
