package encode

import (
	"bytes"
	"testing"

	"github.com/spiegel-im-spiegel/text/detect"
)

func TestFromUTF8To(t *testing.T) {
	testCase := []struct {
		e   detect.CharEncoding
		txt []byte
		res []byte
	}{
		{e: detect.ISO8859L1, res: []byte("Hello World"), txt: []byte("Hello World")},
		{e: detect.UTF8, res: []byte("こんにちは。世界の国から。"), txt: []byte("こんにちは。世界の国から。")},
		{e: detect.ShiftJIS, res: []byte{0x82, 0xb1, 0x82, 0xf1, 0x82, 0xc9, 0x82, 0xbf, 0x82, 0xcd, 0x81, 0x42, 0x90, 0xa2, 0x8a, 0x45, 0x82, 0xcc, 0x8d, 0x91, 0x82, 0xa9, 0x82, 0xe7, 0x81, 0x42}, txt: []byte("こんにちは。世界の国から。")},
		{e: detect.EUCJP, res: []byte{0xa4, 0xb3, 0xa4, 0xf3, 0xa4, 0xcb, 0xa4, 0xc1, 0xa4, 0xcf, 0xa1, 0xa3, 0xc0, 0xa4, 0xb3, 0xa6, 0xa4, 0xce, 0xb9, 0xf1, 0xa4, 0xab, 0xa4, 0xe9, 0xa1, 0xa3}, txt: []byte("こんにちは。世界の国から。")},
		{e: detect.ISO2022JP, res: []byte{0x1b, 0x24, 0x42, 0x24, 0x33, 0x24, 0x73, 0x24, 0x4b, 0x24, 0x41, 0x24, 0x4f, 0x40, 0x24, 0x33, 0x26, 0x1b, 0x28, 0x42, 0x0a}, txt: []byte("こんにちは世界\n")},
	}

	for _, tst := range testCase {
		res, err := FromUTF8To(tst.e, bytes.NewReader(tst.txt))
		if err != nil {
			t.Errorf("ToUTF8ja(%v) is \"%v\", want nil.", tst.txt, err)
		}
		buf := &bytes.Buffer{}
		if _, err := buf.ReadFrom(res); err != nil {
			t.Errorf("ToUTF8ja(%v) is \"%v\", want nil.", tst.txt, err)
		}
		if !bytes.Equal(buf.Bytes(), tst.res) {
			t.Errorf("ToUTF8ja(%v)  = \"%v\", want \"%v\".", tst.txt, buf, tst.res)
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
