package facade

import (
	"io"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/text"
	"github.com/spiegel-im-spiegel/text/cli/gonkf/list"
	"github.com/spiegel-im-spiegel/text/cli/gonkf/nwline"
	"github.com/spiegel-im-spiegel/text/newline"
)

//newNwlineCmd returns cobra.Command instance for nwline sub-command
func newNwlineCmd() *cobra.Command {
	nwlineCmd := &cobra.Command{
		Use:   "nwline [flags] [text file]",
		Short: "Convert newline of text",
		Long:  "Convert newline of text",
		RunE: func(cmd *cobra.Command, args []string) error {
			form := newline.LF
			str, err := cmd.Flags().GetString("form")
			if err != nil {
				return err
			}
			if len(str) > 0 {
				nl := newline.TypeofNewline(str)
				if nl == newline.Unknown {
					return errors.Wrapf(text.ErrNoImplement, "error form %s", str)
				}
				form = nl
			}
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
			dst, err := nwline.Run(reader, form)
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

	nwlineCmd.Flags().StringP("form", "f", "lf", "newline form ["+list.AvailableNewlineOptionsList("|")+"]")
	nwlineCmd.Flags().StringP("output", "o", "", "output file path")

	return nwlineCmd
}
