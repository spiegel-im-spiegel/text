package facade

import (
	"runtime"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gocli/exitcode"
	"github.com/spiegel-im-spiegel/gocli/rwi"
	"github.com/spiegel-im-spiegel/text/ecode"
)

var (
	//Name is applicatin name
	Name = "gonkf"
	//Version is version for applicatin
	Version string
)

var (
	cui = rwi.New() //CUI instance
)

//newRootCmd returns cobra.Command instance for root command
func newRootCmd(ui *rwi.RWI, args []string) *cobra.Command {
	cui = ui
	rootCmd := &cobra.Command{
		Use:   Name,
		Short: "Network Kanji Filter by Golang",
		Long:  "Network Kanji Filter by Golang",
		RunE: func(cmd *cobra.Command, args []string) error {
			return errs.Wrap(ecode.ErrNoImplement, "no command")
		},
	}

	rootCmd.AddCommand(
		newConvCmd(),
		newGuessCmd(),
		newNwlineCmd(),
		newNormCmd(),
		newVersionCmd(),
		newWidthCmd(),
	)
	rootCmd.SetArgs(args)
	rootCmd.SetOutput(ui.ErrorWriter())

	return rootCmd
}

//Execute is called from main function
func Execute(ui *rwi.RWI, args []string) (exit exitcode.ExitCode) {
	defer func() {
		//panic hundling
		if r := recover(); r != nil {
			_ = cui.OutputErrln("Panic:", r)
			for depth := 0; ; depth++ {
				pc, src, line, ok := runtime.Caller(depth)
				if !ok {
					break
				}
				_ = cui.OutputErrln(" ->", depth, ":", runtime.FuncForPC(pc).Name(), ":", src, ":", line)
			}
			exit = exitcode.Abnormal
		}
	}()

	//execution
	exit = exitcode.Normal
	if err := newRootCmd(ui, args).Execute(); err != nil {
		exit = exitcode.Abnormal
	}
	return
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
