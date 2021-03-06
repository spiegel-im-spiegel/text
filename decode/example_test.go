package decode_test

import (
	"bytes"
	"fmt"

	"github.com/spiegel-im-spiegel/text/decode"
)

func ExampleToUTF8() {
	jisText := []byte{0x1b, 0x24, 0x42, 0x24, 0x33, 0x24, 0x73, 0x24, 0x4b, 0x24, 0x41, 0x24, 0x4f, 0x40, 0x24, 0x33, 0x26, 0x1b, 0x28, 0x42}
	res, err := decode.ToUTF8(bytes.NewReader(jisText))
	if err != nil {
		fmt.Println(err)
		return
	}
	buf := &bytes.Buffer{}
	if _, err := buf.ReadFrom(res); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(buf)
	// Output:
	// こんにちは世界
}

func ExampleToUTF8ja() {
	jisText := []byte{0x1b, 0x24, 0x42, 0x24, 0x33, 0x24, 0x73, 0x24, 0x4b, 0x24, 0x41, 0x24, 0x4f, 0x40, 0x24, 0x33, 0x26, 0x1b, 0x28, 0x42}
	res, err := decode.ToUTF8ja(bytes.NewReader(jisText))
	if err != nil {
		fmt.Println(err)
		return
	}
	buf := &bytes.Buffer{}
	if _, err := buf.ReadFrom(res); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(buf)
	// Output:
	// こんにちは世界
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
