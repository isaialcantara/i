package gen

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "gen",
	Aliases: []string{"g", "generate"},
	Short:   "Generates useful data",
	Long:    "This command can be used to generate many kinds of data, like UUIDs.",
}

func init() {
	Cmd.AddCommand(uuidCmd)
}
