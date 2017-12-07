package norm

import (
	"bytes"
	"io"
	"testing"

	"github.com/spiegel-im-spiegel/text/normalize"
)

func TestRun(t *testing.T) {
	testCase := []struct {
		opt normalize.Option
		txt string
		res string
	}{
		{opt: normalize.Unknown, txt: "ペンギン", res: "ペンギン"},
		{opt: normalize.NFC, txt: "ペンギン", res: "ペンギン"},
		{opt: normalize.NFD, txt: "ペンギン", res: "ペンギン"},
		{opt: normalize.NFKC, txt: "ﾍﾟﾝｷﾞﾝ", res: "ペンギン"},
		{opt: normalize.NFKD, txt: "ﾍﾟﾝｷﾞﾝ", res: "ペンギン"},
	}

	for _, tst := range testCase {
		res := Run(bytes.NewBufferString(tst.txt), tst.opt)
		buf := new(bytes.Buffer)
		io.Copy(buf, res)
		if buf.String() != tst.res {
			t.Errorf("Do(%v)  = \"%v\", want \"%v\".", tst.txt, buf.Bytes(), []byte(tst.res))
		}
	}
}
