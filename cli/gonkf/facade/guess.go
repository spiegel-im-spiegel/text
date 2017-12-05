package facade

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/text/cli/gonkf/guess"
)

// guessCmd represents the guess command
var guessCmd = &cobra.Command{
	Use:   "guess [flags] [text file]",
	Short: "Guess character encoding of text",
	Long:  "Guess character encoding of text",
	RunE: func(cmd *cobra.Command, args []string) error {
		reader := cui.Reader()
		if len(args) > 0 {
			file, err := os.Open(args[0]) //args[0] is maybe file path
			if err != nil {
				return err
			}
			defer file.Close()
			reader = file
		}
		encoding := guess.Run(reader)
		cui.Outputln(encoding)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(guessCmd)
}
