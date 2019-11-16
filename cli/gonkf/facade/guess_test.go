package facade

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/spiegel-im-spiegel/gocli/exitcode"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

func TestGuessNormal(t *testing.T) {
	testCases := []struct {
		args   []string
		in     string
		out    string
		outErr string
	}{
		{args: []string{"guess", "testdata/UTF8.txt"}, in: "", out: "UTF-8\n", outErr: ""},
		{args: []string{"guess", "testdata/SHIFT_JIS.txt"}, in: "", out: "Shift_JIS\n", outErr: ""},
		{args: []string{"guess", "testdata/EUCJP.txt"}, in: "", out: "EUC-JP\n", outErr: ""},
		{args: []string{"guess", "testdata/JIS.txt"}, in: "", out: "ISO-2022-JP\n", outErr: ""},
		{args: []string{"guess"}, in: "ペンギン", out: "UTF-8\n", outErr: ""},
	}

	for _, tc := range testCases {
		in := strings.NewReader(tc.in)
		out := new(bytes.Buffer)
		errOut := new(bytes.Buffer)
		ui := rwi.New(
			rwi.WithReader(in),
			rwi.WithWriter(out),
			rwi.WithErrorWriter(errOut),
		)
		exit := Execute(ui, tc.args)
		if exit != exitcode.Normal {
			t.Errorf("Execute() err = \"%v\", want \"%v\".", exit, exitcode.Normal)
		}
		if out.String() != tc.out {
			t.Errorf("Execute() Stdout = \"%v\", want \"%v\".", out.String(), tc.out)
		}
		if errOut.String() != tc.outErr {
			t.Errorf("Execute() Stderr = \"%v\", want \"%v\".", errOut.String(), tc.outErr)
		}
	}
}

func TestGuessOptErr(t *testing.T) {
	testCases := []struct {
		args   []string
		in     string
		out    string
		outErr string
	}{
		{args: []string{"guess", "-x", "testdata/UTF8.txt"}, out: ""},
		{args: []string{"guess", "testdata/noexist.txt"}, out: ""},
	}

	for _, tc := range testCases {
		out := new(bytes.Buffer)
		ui := rwi.New(
			rwi.WithWriter(out),
		)
		err := newRootCmd(ui, tc.args).Execute()
		if err == nil {
			t.Error("Execute() err = nil, not want nil.")
		} else {
			fmt.Println("info:", err)
		}
		if out.String() != tc.out {
			t.Errorf("Execute() Stdout = \"%v\", want \"%v\".", out.String(), tc.out)
		}
	}
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
