package cmd

import (
	"os"

	"github.com/isaialcantara/i/cmd/gen"
	"github.com/spf13/cobra"
)

type App struct {
	rootCmd *cobra.Command
}

func NewApp() *App {
	app := &App{
		rootCmd: &cobra.Command{
			Use:   "i",
			Short: "Isai's Swiss Army Knife.",
			Long:  `A collection of tools useful for software development. Well, at least useful to me.`,
		},
	}

	app.rootCmd.AddCommand(gen.GenCmd())
	return app
}

func (a *App) Execute() {
	err := a.rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
