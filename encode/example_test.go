package encode_test

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/spiegel-im-spiegel/text/detect"
	"github.com/spiegel-im-spiegel/text/encode"
)

func ExampleFromUTF8To() {
	utf8Text := "こんにちは，世界\n"
	res, err := encode.FromUTF8To(detect.ISO2022JP, strings.NewReader(utf8Text))
	if err != nil {
		fmt.Println(err)
		return
	}
	buf := &bytes.Buffer{}
	if _, err := io.Copy(buf, res); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(buf.Bytes())
	// Output:
	// [27 36 66 36 51 36 115 36 75 36 65 36 79 33 36 64 36 51 38 27 40 66 10]
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
