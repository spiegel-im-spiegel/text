package facade

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/spiegel-im-spiegel/gocli/exitcode"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

func TestNormNormal(t *testing.T) {
	testCases := []struct {
		args   []string
		in     string
		out    string
		outErr string
	}{
		{args: []string{"norm", "-f", "nfkc"}, in: "ﾍﾟﾝｷﾞﾝ", out: "ペンギン", outErr: ""},
		{args: []string{"norm", "-f", "nfkd"}, in: "ﾍﾟﾝｷﾞﾝ", out: "ペンギン", outErr: ""},
		{args: []string{"norm", "-f", "nfc"}, in: "ペンギン", out: "ペンギン", outErr: ""},
		{args: []string{"norm", "-f", "nfd"}, in: "ペンギン", out: "ペンギン", outErr: ""},
		{args: []string{"norm"}, in: "ペンギン", out: "ペンギン", outErr: ""},
		{args: []string{"norm", "-o", "testdata/out.test", "testdata/UTF8.txt"}, in: "", out: "", outErr: ""},
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

func TestNormOptErr(t *testing.T) {
	testCases := []struct {
		args   []string
		in     string
		out    string
		outErr string
	}{
		{args: []string{"norm", "-f", "xxx"}, out: ""},
		{args: []string{"norm", "-f"}, out: ""},
		{args: []string{"norm", "testdata/noexist.txt"}, out: ""},
		{args: []string{"norm", "-o", "testdata", "testdata/UTF8.txt"}, out: ""},
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
