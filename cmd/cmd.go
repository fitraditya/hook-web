package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use: "omnichan",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
		SilenceUsage: true,
	}
)

func Execute() {
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(webCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
