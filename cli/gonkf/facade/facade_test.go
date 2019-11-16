package facade

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/spiegel-im-spiegel/gocli/exitcode"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

func TestFacadeNoCmd(t *testing.T) {
	testCases := []struct {
		args   []string
		in     string
		out    string
		outErr string
	}{
		{args: []string{}, out: ""},
	}

	for _, tc := range testCases {
		out := new(bytes.Buffer)
		ui := rwi.New(
			rwi.WithWriter(out),
		)
		exit := Execute(ui, tc.args)
		if exit != exitcode.Abnormal {
			t.Errorf("Execute() err = \"%v\", want \"%v\".", exit, exitcode.Abnormal)
		}
		if out.String() != tc.out {
			t.Errorf("Execute() Stdout = \"%v\", want \"%v\".", out.String(), tc.out)
		}
	}
}

func TestFacadeCmdErr(t *testing.T) {
	testCases := []struct {
		args   []string
		in     string
		out    string
		outErr string
	}{
		{args: []string{"xxx"}, out: ""},
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
