package facade

import (
	"io"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/text"
	"github.com/spiegel-im-spiegel/text/cli/gonkf/list"
	"github.com/spiegel-im-spiegel/text/cli/gonkf/wdth"
	"github.com/spiegel-im-spiegel/text/width"
)

//newWidthCmd returns cobra.Command instance for width sub-command
func newWidthCmd() *cobra.Command {
	widthCmd := &cobra.Command{
		Use:   "width [flags] [text file]",
		Short: "Convert character width of text",
		Long:  "Convert character width of text (UTF-8 text only)",
		RunE: func(cmd *cobra.Command, args []string) error {
			str, _ := cmd.Flags().GetString("form")
			form := width.FormofWidth(str)
			if form == width.Unknown {
				return errors.Wrapf(text.ErrNoImplement, "error form %s", str)
			}
			outPath, _ := cmd.Flags().GetString("output")

			reader := cui.Reader()
			if len(args) > 0 {
				file, err := os.OpenFile(args[0], os.O_RDONLY, 0400) //args[0] is maybe file path
				if err != nil {
					return err
				}
				defer file.Close()
				reader = file
			}
			dst := wdth.Run(reader, form)

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

	widthCmd.Flags().StringP("form", "f", "fold", "form of width ["+list.WidthOptionsList("|")+"]")
	widthCmd.Flags().StringP("output", "o", "", "output file path")

	return widthCmd
}
