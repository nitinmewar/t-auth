package commands

import (
	"tauth/constants"
	"tauth/utils"

	"github.com/spf13/cobra"
)

func Migrate() *cobra.Command {
	return &cobra.Command{
		Use: constants.Command.MIGRATE,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := utils.Migrate()
			return err
		},
	}
}
