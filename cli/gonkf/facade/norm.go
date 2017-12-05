package facade

import (
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/text/cli/gonkf/norm"
	"github.com/spiegel-im-spiegel/text/normalize"
)

// normCmd represents the conv command
var normCmd = &cobra.Command{
	Use:   "norm [flags] [text file]",
	Short: "Unicode normalization",
	Long:  "Unicode normalization",
	RunE: func(cmd *cobra.Command, args []string) error {
		str, err := cmd.Flags().GetString("form")
		if err != nil {
			return err
		}
		form := normalize.FormofNormalize(str)
		outPath, err := cmd.Flags().GetString("output")
		if err != nil {
			return err
		}

		reader := cui.Reader()
		if len(args) > 0 {
			file, err2 := os.Open(args[0]) //args[0] is maybe file path
			if err != nil {
				return err2
			}
			defer file.Close()
			reader = file
		}
		dst, err := norm.Run(reader, form)
		if err != nil {
			return err
		}

		if len(outPath) > 0 {
			file, err := os.OpenFile(outPath, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				return err
			}
			defer file.Close()
			io.Copy(file, dst)
		} else {
			cui.WriteFrom(dst)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(normCmd)

	normCmd.Flags().StringP("form", "f", "nfc", "normalization form")
	normCmd.Flags().StringP("output", "o", "", "output file path")
}
