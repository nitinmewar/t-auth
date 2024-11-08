package commands

import (
	"tauth/config"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

/* -------------------------------------------------------------------------- */
/*                             positive test cases                            */
/* -------------------------------------------------------------------------- */
func TestDropTables(t *testing.T) {
	t.Skip()
	config.LoadConfigs()

	cmd := &cobra.Command{}

	cmd.AddCommand(DropTables())
	err := DropTables().RunE(cmd, nil)
	assert.Empty(t, err)
}

func TestMigrateTables(t *testing.T) {
	config.LoadConfigs()

	cmd := &cobra.Command{}

	cmd.AddCommand(Migrate())
	err := Migrate().RunE(cmd, nil)
	assert.Empty(t, err)
}

func TestSeedTables(t *testing.T) {
	t.Skip()
	config.LoadConfigs()

	cmd := &cobra.Command{}

	cmd.AddCommand(Seed())
	err := Migrate().RunE(cmd, nil)
	assert.Empty(t, err)
}
