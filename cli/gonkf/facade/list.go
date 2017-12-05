package facade

import (
	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/text/cli/gonkf/list"
)

// listCmd represents the version command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List of available character encoding",
	Long:  "List of available character encoding",
	RunE: func(cmd *cobra.Command, args []string) error {
		cui.Outputln("available encoding:", list.AvailableEncodingList())
		cui.Outputln("   type of newline:", list.AvailableNewlineOptionsList())
		cui.Outputln("normalization form:", list.NormOptionsList())
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
