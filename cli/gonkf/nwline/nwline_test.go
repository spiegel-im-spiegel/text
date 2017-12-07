package nwline

import (
	"bytes"
	"io"
	"testing"

	"github.com/spiegel-im-spiegel/text/newline"
)

func TestRun(t *testing.T) {
	testCase := []struct {
		nl  newline.Option
		txt string
		res string
	}{
		{nl: newline.Unknown, txt: "こんにちは\nこんにちは\rこんにちは\r\nこんにちは", res: "こんにちは\nこんにちは\rこんにちは\r\nこんにちは"},
		{nl: newline.LF, txt: "こんにちは\nこんにちは\rこんにちは\r\nこんにちは", res: "こんにちは\nこんにちは\nこんにちは\nこんにちは"},
		{nl: newline.CR, txt: "こんにちは\nこんにちは\rこんにちは\r\nこんにちは", res: "こんにちは\rこんにちは\rこんにちは\rこんにちは"},
		{nl: newline.CRLF, txt: "こんにちは\nこんにちは\rこんにちは\r\nこんにちは", res: "こんにちは\r\nこんにちは\r\nこんにちは\r\nこんにちは"},
	}

	for _, tst := range testCase {
		res := Run(bytes.NewBufferString(tst.txt), tst.nl)
		buf := new(bytes.Buffer)
		io.Copy(buf, res)
		if buf.String() != tst.res {
			t.Errorf("Newline(%s)  = \"%v\", want \"%v\".", tst.txt, buf, tst.res)
		}
	}
}
