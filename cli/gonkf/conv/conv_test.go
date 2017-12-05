package conv

import (
	"bytes"
	"io"
	"testing"

	"github.com/spiegel-im-spiegel/text/detect"
	"github.com/spiegel-im-spiegel/text/newline"
)

var testCase = []struct {
	from detect.CharEncoding
	to   detect.CharEncoding
	txt  []byte
	res  []byte
}{
	{
		from: detect.ISO2022JP,
		to:   detect.ShiftJIS,
		txt:  []byte{0x1b, 0x24, 0x42, 0x24, 0x33, 0x24, 0x73, 0x24, 0x4b, 0x24, 0x41, 0x24, 0x4f, 0x40, 0x24, 0x33, 0x26, 0x1b, 0x28, 0x42},
		res:  []byte{0x82, 0xb1, 0x82, 0xf1, 0x82, 0xc9, 0x82, 0xbf, 0x82, 0xcd, 0x90, 0xa2, 0x8a, 0x45},
	},
	{
		from: detect.ISO2022JP,
		to:   detect.UTF8,
		txt:  []byte{0x1b, 0x24, 0x42, 0x24, 0x33, 0x24, 0x73, 0x24, 0x4b, 0x24, 0x41, 0x24, 0x4f, 0x40, 0x24, 0x33, 0x26, 0x1b, 0x28, 0x42},
		res:  []byte("こんにちは世界"),
	},
	{
		from: detect.UTF8,
		to:   detect.ISO2022JP,
		txt:  []byte("こんにちは世界\n"),
		res:  []byte{0x1b, 0x24, 0x42, 0x24, 0x33, 0x24, 0x73, 0x24, 0x4b, 0x24, 0x41, 0x24, 0x4f, 0x40, 0x24, 0x33, 0x26, 0x1b, 0x28, 0x42, 0x0a},
	},
}

func TestRunUnknown(t *testing.T) {

	for _, tst := range testCase {
		opt := &Options{}
		res, err := Run(bytes.NewReader(tst.txt), opt)
		if err != nil {
			t.Errorf("Run(%v)  = \"%v\", want nil.", tst.txt, err)
		}
		buf := new(bytes.Buffer)
		io.Copy(buf, res)
		if bytes.Compare(buf.Bytes(), tst.txt) != 0 {
			t.Errorf("Run(%v)  = \"%v\", want \"%v\".", tst.txt, buf, tst.txt)
		}
	}
}

func TestRunErr(t *testing.T) {

	for _, tst := range testCase {
		opt := &Options{}
		opt.SetSrcEncoding(tst.from)
		_, err := Run(bytes.NewReader(tst.txt), opt)
		if err == nil {
			t.Errorf("Run(%v) = nil, want not nil.", tst.txt)
		}
	}
}

func TestRun(t *testing.T) {

	for _, tst := range testCase {
		opt := &Options{}
		opt.SetSrcEncoding(tst.from)
		opt.SetDstEncoding(tst.to)
		opt.SetNewline(newline.LF)
		res, err := Run(bytes.NewReader(tst.txt), opt)
		if err != nil {
			t.Errorf("Run(%v)  = \"%v\", want nil.", tst.txt, err)
		}
		buf := new(bytes.Buffer)
		io.Copy(buf, res)
		if bytes.Compare(buf.Bytes(), tst.res) != 0 {
			t.Errorf("Run(%v)  = \"%v\", want \"%v\".", tst.txt, buf, tst.res)
		}
	}
}

func TestRunGuess(t *testing.T) {

	for _, tst := range testCase {
		opt := &Options{}
		opt.SetDstEncoding(tst.to)
		res, err := Run(bytes.NewReader(tst.txt), opt)
		if err != nil {
			t.Errorf("Run(%v)  = \"%v\", want nil.", tst.txt, err)
		}
		buf := new(bytes.Buffer)
		io.Copy(buf, res)
		if bytes.Compare(buf.Bytes(), tst.res) != 0 {
			t.Errorf("Run(%v)  = \"%v\", want \"%v\".", tst.txt, buf, tst.res)
		}
	}
}
