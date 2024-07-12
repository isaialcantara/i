package cmd

import (
	"os"

	"github.com/isaialcantara/i/cmd/gen"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "i",
	Short: "Isai's Swiss Army Knife.",
	Long:  `A collection of tools useful for software development. Well, at least useful to me.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(gen.Cmd)
}
