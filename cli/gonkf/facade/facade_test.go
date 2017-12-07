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
			rwi.Writer(out),
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
