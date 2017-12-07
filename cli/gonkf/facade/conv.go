package facade

import (
	"io"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/text"
	"github.com/spiegel-im-spiegel/text/cli/gonkf/conv"
	"github.com/spiegel-im-spiegel/text/cli/gonkf/list"
	"github.com/spiegel-im-spiegel/text/detect"
)

//newConvCmd returns cobra.Command instance for conv sub-command
func newConvCmd() *cobra.Command {
	convCmd := &cobra.Command{
		Use:   "conv [flags] [text file]",
		Short: "Convert character encoding of text",
		Long:  "Convert character encoding of text",
		RunE: func(cmd *cobra.Command, args []string) error {
			opt := &conv.Options{}
			str, err := cmd.Flags().GetString("src-encoding")
			if err != nil {
				return err
			}
			if len(str) > 0 {
				e := list.TypeofEncoding(str)
				if e == detect.Unknown {
					return errors.Wrapf(text.ErrNoImplement, "error character encoding %s", str)
				}
				opt.SetSrcEncoding(e)
			}
			str, err = cmd.Flags().GetString("dst-encoding")
			if err != nil {
				return err
			}
			if len(str) > 0 {
				e := list.TypeofEncoding(str)
				if e == detect.Unknown {
					return errors.Wrapf(text.ErrNoImplement, "error character encoding %s", str)
				}
				opt.SetDstEncoding(e)
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
			dst, err := conv.Run(reader, opt)
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

	convCmd.Flags().StringP("src-encoding", "s", "", "encoding of src ["+list.AvailableEncodingList("|")+"]")
	convCmd.Flags().StringP("dst-encoding", "d", "utf8", "encoding of dest ["+list.AvailableEncodingList("|")+"]")
	convCmd.Flags().StringP("output", "o", "", "output file path")

	return convCmd
}
