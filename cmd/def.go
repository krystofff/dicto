package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/bregydoc/gtranslate"
)

// defCmd represents the def command
var defCmd = &cobra.Command{
	Use:   "def",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		text := "Hed"egydoc/gtrnslate
		translated,spf13gcobanTranslateWithParams(
			text,
			gtranslate.TranslationParams{
				From: "en",
				To:   "es",
			},
		)
		if err != nil {
			panic(err)
		}

		fmt.Printf("en: %s | es: %s \n", text, translated)
	},
}

func init() {
	rootCmd.AddCommand(defCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// defCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// defCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
