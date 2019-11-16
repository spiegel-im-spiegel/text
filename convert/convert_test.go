package convert

import (
	"bytes"
	"testing"

	"github.com/spiegel-im-spiegel/text/detect"
)

func TestToJa(t *testing.T) {
	testCase := []struct {
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
	}

	for _, tst := range testCase {
		res, err := ToJa(tst.to, bytes.NewReader(tst.txt))
		if err != nil {
			t.Errorf("ToJa(%v)  = \"%v\", want nil.", tst.txt, err)
		}
		buf := &bytes.Buffer{}
		if _, err := buf.ReadFrom(res); err != nil {
			t.Errorf("ToJa(%v) is \"%v\", want nil.", tst.txt, err)
		}
		if !bytes.Equal(buf.Bytes(), tst.res) {
			t.Errorf("ToJa(%v) = \"%v\", want \"%v\".", tst.txt, buf, tst.res)
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
