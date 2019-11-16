package detect

import (
	"testing"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/encoding/unicode"
)

func TestCharEncodingString(t *testing.T) {
	testCase := []struct {
		e CharEncoding
		s string
	}{
		{e: Unknown, s: "Unknown"},
		{e: UTF8, s: "UTF-8"},
		{e: ISO8859L1, s: "ISO-8859-1"},
		{e: ShiftJIS, s: "Shift_JIS"},
		{e: EUCJP, s: "EUC-JP"},
		{e: ISO2022JP, s: "ISO-2022-JP"},
		{e: EUCKR, s: "EUC-KR"},
		{e: GB18030, s: "GB-18030"},
		{e: Big5, s: "Big5"},
		{e: CharEncoding(9), s: "Unknown"},
	}

	for _, tst := range testCase {
		if tst.e.String() != tst.s {
			t.Errorf("Encoding(%d)  = \"%v\", want \"%v\".", int(tst.e), tst.e, tst.s)
		}
	}
}

func TestGetCharEncoding(t *testing.T) {
	testCase := []struct {
		e CharEncoding
		s string
	}{
		{e: Unknown, s: "foobar"},
		{e: UTF8, s: "UTF-8"},
		{e: ISO8859L1, s: "ISO-8859-1"},
		{e: ShiftJIS, s: "Shift_JIS"},
		{e: EUCJP, s: "EUC-JP"},
		{e: ISO2022JP, s: "ISO-2022-JP"},
		{e: EUCKR, s: "EUC-KR"},
		{e: GB18030, s: "GB-18030"},
		{e: Big5, s: "Big5"},
	}

	for _, tst := range testCase {
		e := typeofEncoding(tst.s)
		if e != tst.e {
			t.Errorf("typeofEncoding(%v)  = \"%v\", want \"%v\".", tst.s, e, tst.e)
		}
	}
}

func TestGetEncoding(t *testing.T) {
	testCase := []struct {
		e   CharEncoding
		enc encoding.Encoding
	}{
		{e: Unknown, enc: nil},
		{e: UTF8, enc: unicode.UTF8},
		{e: ISO8859L1, enc: charmap.ISO8859_1},
		{e: ShiftJIS, enc: japanese.ShiftJIS},
		{e: EUCJP, enc: japanese.EUCJP},
		{e: ISO2022JP, enc: japanese.ISO2022JP},
		{e: EUCKR, enc: korean.EUCKR},
		{e: GB18030, enc: simplifiedchinese.GB18030},
		{e: Big5, enc: traditionalchinese.Big5},
	}

	for _, tst := range testCase {
		enc := tst.e.GetEncoding()
		if enc != tst.enc {
			t.Errorf("GetEncoding(%v)  = \"%v\", want \"%v\".", tst.e, enc, tst.enc)
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
