package facade

import (
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/text/cli/gonkf/list"
	"github.com/spiegel-im-spiegel/text/cli/gonkf/norm"
	"github.com/spiegel-im-spiegel/text/ecode"
	"github.com/spiegel-im-spiegel/text/normalize"
)

//newNormCmd returns cobra.Command instance for norm sub-command
func newNormCmd() *cobra.Command {
	normCmd := &cobra.Command{
		Use:   "norm [flags] [text file]",
		Short: "Unicode normalization",
		Long:  "Unicode normalization (UTF-8 text only)",
		RunE: func(cmd *cobra.Command, args []string) error {
			str, _ := cmd.Flags().GetString("form")
			form := normalize.FormofNormalize(str)
			if form == normalize.Unknown {
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
			dst := norm.Run(reader, form)

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

	normCmd.Flags().StringP("form", "f", "nfc", "normalization form ["+list.NormOptionsList("|")+"]")
	normCmd.Flags().StringP("output", "o", "", "output file path")

	return normCmd
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
