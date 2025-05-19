package gen

import (
	"github.com/spf13/cobra"
)

func GenCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "gen",
		Aliases: []string{"g", "generate"},
		Short:   "Generates useful data",
		Long:    "This command can be used to generate many kinds of data, like UUIDs.",
	}

	cmd.AddCommand(UUIDCmd())

	return cmd
}
