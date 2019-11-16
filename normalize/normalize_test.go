package normalize

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestReader(t *testing.T) {
	testCase := []struct {
		opt Option
		txt string
		res string
	}{
		{opt: Unknown, txt: "ペンギン", res: "ペンギン"},
		{opt: NFC, txt: "ペンギン", res: "ペンギン"},
		{opt: NFD, txt: "ペンギン", res: "ペンギン"},
		{opt: NFKC, txt: "ﾍﾟﾝｷﾞﾝ", res: "ペンギン"},
		{opt: NFKD, txt: "ﾍﾟﾝｷﾞﾝ", res: "ペンギン"},
	}

	for _, tst := range testCase {
		buf := &bytes.Buffer{}
		if _, err := io.Copy(buf, Reader(strings.NewReader(tst.txt), tst.opt)); err != nil {
			t.Errorf("Do(%v) is \"%+v\", want nil.", tst.txt, err)
		}
		if buf.String() != tst.res {
			t.Errorf("Do(%v)  = \"%U\", want \"%U\".", tst.txt, []rune(buf.String()), []rune(tst.res))
		}
	}
}

func TestString(t *testing.T) {
	testCase := []struct {
		opt Option
		txt string
		res string
	}{
		{opt: Unknown, txt: "ペンギン", res: "ペンギン"},
		{opt: NFC, txt: "ペンギン", res: "ペンギン"},
		{opt: NFD, txt: "ペンギン", res: "ペンギン"},
		{opt: NFKC, txt: "ﾍﾟﾝｷﾞﾝ", res: "ペンギン"},
		{opt: NFKD, txt: "ﾍﾟﾝｷﾞﾝ", res: "ペンギン"},
	}

	for _, tst := range testCase {
		res := String(tst.txt, tst.opt)
		if res != tst.res {
			t.Errorf("Do(%v)  = \"%U\", want \"%U\".", tst.txt, []rune(res), []rune(tst.res))
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
