package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/kristofff/dicto/utils"
	"github.com/spf13/cobra"
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
		input := args[0]
		url := "https://api.dictionaryapi.dev/api/v1/entries/en/" + input
		fmt.Println("Getting data | url: ", url)
		res := utils.Request(url, "application/json")
		dst := &bytes.Buffer{}
		if err := json.Indent(dst, res, "", "  "); err != nil {
			panic(err)
		}

		fmt.Println(dst.String())
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
