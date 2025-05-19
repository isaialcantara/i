package gen

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func UUIDCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "uuid",
		Short: "Generates a UUID.",
		Long:  `This command generates a new UUID. By default, the command will generate a UUIDv4.`,
		Run: func(cmd *cobra.Command, args []string) {
			value := uuid.New().String()
			fmt.Println(value)
		},
	}
}
