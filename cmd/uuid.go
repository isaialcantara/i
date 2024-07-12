package cmd

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// uuidCmd represents the uuid command
var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "Generates or validates a UUID.",
	Long: `This command generates a new UUID or validates an existing one.
	By default, the command will generate a UUIDv4.`,
	Run: func(cmd *cobra.Command, args []string) {
		value := uuid.New().String()

		fmt.Println(value)
	},
}

func init() {
	rootCmd.AddCommand(uuidCmd)
}
