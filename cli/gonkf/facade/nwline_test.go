package facade

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/spiegel-im-spiegel/gocli/exitcode"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

func TestNwlineNormal(t *testing.T) {
	testCases := []struct {
		args   []string
		in     string
		out    string
		outErr string
	}{
		{args: []string{"nwline", "-f", "lf"}, in: "こんにちは\nこんにちは\rこんにちは\r\nこんにちは", out: "こんにちは\nこんにちは\nこんにちは\nこんにちは", outErr: ""},
		{args: []string{"nwline", "-f", "cr"}, in: "こんにちは\nこんにちは\rこんにちは\r\nこんにちは", out: "こんにちは\rこんにちは\rこんにちは\rこんにちは", outErr: ""},
		{args: []string{"nwline", "-f", "crlf"}, in: "こんにちは\nこんにちは\rこんにちは\r\nこんにちは", out: "こんにちは\r\nこんにちは\r\nこんにちは\r\nこんにちは", outErr: ""},
		{args: []string{"nwline"}, in: "こんにちは\nこんにちは\rこんにちは\r\nこんにちは", out: "こんにちは\nこんにちは\nこんにちは\nこんにちは", outErr: ""},
		{args: []string{"nwline", "-o", "testdata/out.test", "testdata/UTF8.txt"}, in: "", out: "", outErr: ""},
	}

	for _, tc := range testCases {
		in := bytes.NewBufferString(tc.in)
		out := new(bytes.Buffer)
		errOut := new(bytes.Buffer)
		ui := rwi.New(
			rwi.Reader(in),
			rwi.Writer(out),
			rwi.ErrorWriter(errOut),
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

func TestNwlineOptErr(t *testing.T) {
	testCases := []struct {
		args   []string
		in     string
		out    string
		outErr string
	}{
		{args: []string{"nwline", "-f", "xxx"}, out: ""},
		{args: []string{"nwline", "-f"}, out: ""},
		{args: []string{"nwline", "testdata/noexist.txt"}, out: ""},
		{args: []string{"nwline", "-o", "testdata", "testdata/UTF8.txt"}, out: ""},
	}

	for _, tc := range testCases {
		out := new(bytes.Buffer)
		ui := rwi.New(
			rwi.Writer(out),
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
