package facade

import (
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/text/cli/gonkf/conv"
	"github.com/spiegel-im-spiegel/text/cli/gonkf/list"
	"github.com/spiegel-im-spiegel/text/detect"
	"github.com/spiegel-im-spiegel/text/ecode"
)

//newConvCmd returns cobra.Command instance for conv sub-command
func newConvCmd() *cobra.Command {
	convCmd := &cobra.Command{
		Use:   "conv [flags] [text file]",
		Short: "Convert character encoding of text",
		Long:  "Convert character encoding of text",
		RunE: func(cmd *cobra.Command, args []string) error {
			opt := &conv.Options{}
			str, _ := cmd.Flags().GetString("src-encoding")
			if len(str) > 0 {
				e := list.TypeofEncoding(str)
				if e == detect.Unknown {
					return errs.Wrap(ecode.ErrNoImplement, "", errs.WithContext("src-encoding", str))
				}
				opt.SetSrcEncoding(e)
			}
			str, _ = cmd.Flags().GetString("dst-encoding")
			if len(str) > 0 {
				e := list.TypeofEncoding(str)
				if e == detect.Unknown {
					return errs.Wrap(ecode.ErrNoImplement, "", errs.WithContext("dst-encoding", str))
				}
				opt.SetDstEncoding(e)
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
				_, err = io.Copy(file, dst)
				return err
			}
			return cui.WriteFrom(dst)
		},
	}
	convCmd.Flags().StringP("src-encoding", "s", "", "encoding of src ["+list.AvailableEncodingList("|")+"]")
	convCmd.Flags().StringP("dst-encoding", "d", "utf8", "encoding of dest ["+list.AvailableEncodingList("|")+"]")
	convCmd.Flags().StringP("output", "o", "", "output file path")

	return convCmd
}

/* MIT License
 *
 * Copyright 2017-2019 Spiegel
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */
