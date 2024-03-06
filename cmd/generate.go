package cmd

import (
	"github.com/paulcalimache/go-cv/pkg/cv"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a curriculum vitae",
	Long:  `Generate a curriculum vitae based on the config file pass in flag`,
	RunE:  generate,
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringP("file", "f", "", "Yaml data file")
	generateCmd.MarkFlagRequired("file")

	generateCmd.Flags().StringP("output", "o", "output.html", "Output directory")
}

func generate(cmd *cobra.Command, args []string) error {
	file, err := cmd.Flags().GetString("file")
	if err != nil {
		return err
	}
	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return err
	}
	return cv.Generate(file, output)
}
