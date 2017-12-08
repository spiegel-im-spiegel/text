package newline

import (
	"bytes"
	"io"
	"testing"
)

func TestConvert(t *testing.T) {
	testCase := []struct {
		nl  Option
		txt string
		res string
	}{
		{nl: Unknown, txt: "こんにちは\nこんにちは\rこんにちは\r\nこんにちは", res: "こんにちは\nこんにちは\rこんにちは\r\nこんにちは"},
		{nl: LF, txt: "こんにちは\nこんにちは\rこんにちは\r\nこんにちは", res: "こんにちは\nこんにちは\nこんにちは\nこんにちは"},
		{nl: CR, txt: "こんにちは\nこんにちは\rこんにちは\r\nこんにちは", res: "こんにちは\rこんにちは\rこんにちは\rこんにちは"},
		{nl: CRLF, txt: "こんにちは\nこんにちは\rこんにちは\r\nこんにちは", res: "こんにちは\r\nこんにちは\r\nこんにちは\r\nこんにちは"},
	}

	for _, tst := range testCase {
		res := Reader(bytes.NewBufferString(tst.txt), tst.nl)
		buf := new(bytes.Buffer)
		io.Copy(buf, res)
		if buf.String() != tst.res {
			t.Errorf("Newline(%s)  = \"%v\", want \"%v\".", tst.txt, buf, tst.res)
		}
	}
}
