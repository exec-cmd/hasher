package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "hasher",
	Short: "A simple hashing utility",
	Long:  "A simple hashing utility that supports multiple hash algorithms",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
