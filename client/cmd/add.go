package cmd

import (
	"os"

	"chainsafe/client/store"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(RegisterAddCommand())
}

// RegisterAddCommand register add subcommand and flags 
func RegisterAddCommand() *cobra.Command {
	c := &cobra.Command{
		Use:  "add",
		Run: func(cmd *cobra.Command, args []string) {
			c := store.NewClient()
			if err := c.AddAndSave(args); err != nil {
				os.Exit(1)
			}
		},
	}
	return c
}
