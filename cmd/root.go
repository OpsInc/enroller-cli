package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

//nolint:gochecknoglobals
var rootCmd = &cobra.Command{
	Use:   "enroller-client",
	Short: "enroller-client is a client to communicate with the enroller ecosystem",
	Long: `A simple cli to send instructions to the enroller ecosystem.
Complete documentation is available at https://github.com/OpsInc/enroller-client`,
	Run: func(cmd *cobra.Command, args []string) { //nolint:revive
		_ = cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

//nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(validateCmd)
	rootCmd.AddCommand(dispatchCmd)
}
