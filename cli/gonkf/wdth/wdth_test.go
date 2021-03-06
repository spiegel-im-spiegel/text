package wdth

import (
	"bytes"
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
		buf := &bytes.Buffer{}
		if _, err := buf.ReadFrom(Run(bytes.NewBufferString(tst.txt), tst.w)); err != nil {
			t.Errorf("Run(%v) is \"%v\", want nil.", tst.txt, err)
		}
		if buf.String() != tst.res {
			t.Errorf("wdth.Run(%s)  = \"%v\", want \"%v\".", tst.txt, buf, tst.res)
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
