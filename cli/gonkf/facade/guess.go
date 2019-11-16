package facade

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/text/cli/gonkf/guess"
)

//newGuessCmd returns cobra.Command instance for guess sub-command
func newGuessCmd() *cobra.Command {
	guessCmd := &cobra.Command{
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
			return cui.Outputln(guess.Run(reader))
		},
	}

	return guessCmd
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
