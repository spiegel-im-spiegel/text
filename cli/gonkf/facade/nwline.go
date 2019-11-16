package facade

import (
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/text/cli/gonkf/list"
	"github.com/spiegel-im-spiegel/text/cli/gonkf/nwline"
	"github.com/spiegel-im-spiegel/text/ecode"
	"github.com/spiegel-im-spiegel/text/newline"
)

//newNwlineCmd returns cobra.Command instance for nwline sub-command
func newNwlineCmd() *cobra.Command {
	nwlineCmd := &cobra.Command{
		Use:   "nwline [flags] [text file]",
		Short: "Convert newline of text",
		Long:  "Convert newline of text",
		RunE: func(cmd *cobra.Command, args []string) error {
			str, _ := cmd.Flags().GetString("form")
			form := newline.TypeofNewline(str)
			if form == newline.Unknown {
				return errs.Wrap(ecode.ErrNoImplement, "", errs.WithContext("form", str))
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
			dst := nwline.Run(reader, form)

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

	nwlineCmd.Flags().StringP("form", "f", "lf", "newline form ["+list.AvailableNewlineOptionsList("|")+"]")
	nwlineCmd.Flags().StringP("output", "o", "", "output file path")

	return nwlineCmd
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
