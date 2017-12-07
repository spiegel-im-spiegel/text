package wdth

import (
	"bytes"
	"io"
	"testing"

	"github.com/spiegel-im-spiegel/text/width"
)

func TestRun(t *testing.T) {
	testCase := []struct {
		w   width.Option
		txt string
		res string
	}{
		{w: width.Unknown, txt: "１２３４５６７８９０ｱｲｳｴｵｶｷｸｹｺＡＢＣＤＥＦＧＨＩＪＫ", res: "１２３４５６７８９０ｱｲｳｴｵｶｷｸｹｺＡＢＣＤＥＦＧＨＩＪＫ"},
		{w: width.Fold, txt: "１２３４５６７８９０ｱｲｳｴｵｶｷｸｹｺＡＢＣＤＥＦＧＨＩＪＫ", res: "1234567890アイウエオカキクケコABCDEFGHIJK"},
		{w: width.Narrow, txt: "１２３４５６７８９０ｱｲｳｴｵｶｷｸｹｺＡＢＣＤＥＦＧＨＩＪＫ", res: "1234567890ｱｲｳｴｵｶｷｸｹｺABCDEFGHIJK"},
		{w: width.Widen, txt: "１２３４５６７８９０ｱｲｳｴｵｶｷｸｹｺＡＢＣＤＥＦＧＨＩＪＫ", res: "１２３４５６７８９０アイウエオカキクケコＡＢＣＤＥＦＧＨＩＪＫ"},
	}

	for _, tst := range testCase {
		res := Run(bytes.NewBufferString(tst.txt), tst.w)
		buf := new(bytes.Buffer)
		io.Copy(buf, res)
		if buf.String() != tst.res {
			t.Errorf("wdth.Run(%s)  = \"%v\", want \"%v\".", tst.txt, buf, tst.res)
		}
	}
}
