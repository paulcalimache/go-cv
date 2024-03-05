/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/paulcalimache/go-cv/generate"
	"github.com/paulcalimache/go-cv/types"
	"github.com/spf13/cobra"
)

var format types.Format

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a curriculum vitae",
	Long:  `Generate a curriculum vitae based on the config file pass in flag`,
	Run: func(cmd *cobra.Command, args []string) {
		data, _ := cmd.Flags().GetString("data")
		output, _ := cmd.Flags().GetString("output")
		err := generate.Generate(data, output, string(format))
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringP("data", "d", "", "data file")
	generateCmd.Flags().StringP("output", "o", "", "Output directory")
	generateCmd.Flags().VarP(&format, "format", "f", "Format of the generated cv (html or pdf)")

	generateCmd.MarkFlagsRequiredTogether("data", "output")
}
