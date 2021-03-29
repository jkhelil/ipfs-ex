package cmd

import (
	"os"
	"github.com/spf13/cobra"
)

func init() {
	addFlag(rootCmd.Flags(), &flag{name: "server-url", short: "s", defaultValue: "http://127.0.0.1:5001/api/v0/add", desc: "Filestore server url"})
	addFlag(rootCmd.Flags(), &flag{name: "log-level", short: "l", desc: "logging verbosity", defaultValue: "info"})
}
// rootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:   "store",
	Short: "store is a tool to operate ipfs storing",
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Usage(); err != nil {
			os.Exit(1)
		}
	},
}

// Execute adds all child commands to the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}