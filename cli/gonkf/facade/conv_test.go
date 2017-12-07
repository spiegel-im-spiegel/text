package facade

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/spiegel-im-spiegel/gocli/exitcode"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

func TestConvNormal(t *testing.T) {
	testCases := []struct {
		args   []string
		in     string
		out    string
		outErr string
	}{
		{args: []string{"conv", "testdata/UTF8.txt"}, in: "", out: "こんにちは世界", outErr: ""},
		{args: []string{"conv", "-s", "utf8", "testdata/UTF8.txt"}, in: "", out: "こんにちは世界", outErr: ""},
		{args: []string{"conv", "-s", "utf8", "-d", "utf8", "testdata/UTF8.txt"}, in: "", out: "こんにちは世界", outErr: ""},
		{args: []string{"conv", "testdata/SHIFT_JIS.txt"}, in: "", out: "こんにちは。世界の国から。", outErr: ""},
		{args: []string{"conv", "-s", "sjis", "testdata/SHIFT_JIS.txt"}, in: "", out: "こんにちは。世界の国から。", outErr: ""},
		{args: []string{"conv", "-s", "sjis", "-d", "utf8", "testdata/SHIFT_JIS.txt"}, in: "", out: "こんにちは。世界の国から。", outErr: ""},
		{args: []string{"conv", "testdata/EUCJP.txt"}, in: "", out: "こんにちは。世界の国から。", outErr: ""},
		{args: []string{"conv", "-s", "euc", "testdata/EUCJP.txt"}, in: "", out: "こんにちは。世界の国から。", outErr: ""},
		{args: []string{"conv", "-s", "euc", "-d", "utf8", "testdata/EUCJP.txt"}, in: "", out: "こんにちは。世界の国から。", outErr: ""},
		{args: []string{"conv", "testdata/JIS.txt"}, in: "", out: "こんにちは世界", outErr: ""},
		{args: []string{"conv", "-s", "jis", "testdata/JIS.txt"}, in: "", out: "こんにちは世界", outErr: ""},
		{args: []string{"conv", "-s", "jis", "-d", "utf8", "testdata/JIS.txt"}, in: "", out: "こんにちは世界", outErr: ""},
		{args: []string{"conv"}, in: "こんにちは世界", out: "こんにちは世界", outErr: ""},
		{args: []string{"conv", "-o", "testdata/out.test"}, in: "こんにちは世界", out: "", outErr: ""},
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

func TestConvNormal2(t *testing.T) {
	testCases := []struct {
		args   []string
		in     string
		out    []byte
		outErr string
	}{
		{args: []string{"conv", "-d", "sjis", "testdata/JIS.txt"}, in: "", out: []byte{0x82, 0xb1, 0x82, 0xf1, 0x82, 0xc9, 0x82, 0xbf, 0x82, 0xcd, 0x90, 0xa2, 0x8a, 0x45}, outErr: ""},
		{args: []string{"conv", "-d", "jis"}, in: "こんにちは世界\n", out: []byte{0x1b, 0x24, 0x42, 0x24, 0x33, 0x24, 0x73, 0x24, 0x4b, 0x24, 0x41, 0x24, 0x4f, 0x40, 0x24, 0x33, 0x26, 0x1b, 0x28, 0x42, 0x0a}, outErr: ""},
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
		if bytes.Compare(out.Bytes(), tc.out) != 0 {
			t.Errorf("Execute() Stdout = \"%v\", want \"%v\".", out.Bytes(), tc.out)
		}
		if errOut.String() != tc.outErr {
			t.Errorf("Execute() Stderr = \"%v\", want \"%v\".", errOut.String(), tc.outErr)
		}
	}
}

func TestConvOptErr(t *testing.T) {
	testCases := []struct {
		args   []string
		in     string
		out    string
		outErr string
	}{
		{args: []string{"conv", "-x", "testdata/UTF8.txt"}, out: ""},
		{args: []string{"conv", "-s", "xxx", "testdata/UTF8.txt"}, out: ""},
		{args: []string{"conv", "-d", "xxx", "testdata/UTF8.txt"}, out: ""},
		{args: []string{"conv", "-o", "testdata", "testdata/UTF8.txt"}, out: ""},
		{args: []string{"conv", "testdata/noexist.txt"}, out: ""},
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
