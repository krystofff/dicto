package cmd

import (
	"fmt"

	"github.com/bregydoc/gtranslate"
	"github.com/spf13/cobra"
)

// transCmd represents the trans command
var transCmd = &cobra.Command{
	Use:   "trans",
	Short: "Translates word from -f  to target -t language codes",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		input := args[0]
		from, _ := cmd.Flags().GetString("from")
		to, _ := cmd.Flags().GetString("to")

		translated, err := gtranslate.TranslateWithParams(
			input,
			gtranslate.TranslationParams{
				From: from,
				To:   to,
			},
		)
		if err != nil {
			panic(err)
		}

		fmt.Printf("from:  %s | to: %s \n", input, translated)

	},
}

func init() {
	rootCmd.AddCommand(transCmd)

	transCmd.Flags().StringP("from", "f", "", "From Language code")
	transCmd.Flags().StringP("to", "t", "", "Language code target")
}
