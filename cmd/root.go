package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-curriculum",
	Short: "Go curriculum is a CLI tool for generating curriculum vitae",
	Long: `
Go curriculum is a CLI tool for generating curriculum vitae in pdf or html format,
based from a yaml config file.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
