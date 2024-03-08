package cmd

import (
	"github.com/paulcalimache/go-cv/pkg/curriculum"
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
	generateCmd.Flags().StringP("template", "t", "classic", "CV Template to use")
}

func generate(cmd *cobra.Command, args []string) error {
	file, _ := cmd.Flags().GetString("file")
	output, _ := cmd.Flags().GetString("output")
	templ, _ := cmd.Flags().GetString("template")

	cv, err := curriculum.ParseCV(file)
	if err != nil {
		return err;
	}
	return cv.Render(output, templ);
}
