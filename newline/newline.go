package newline

import (
	"bytes"
	"io"
	"strings"
)

//Reader returns converted text Reader
func Reader(txt io.Reader, opt Option) io.Reader {
	if opt == Unknown {
		return txt
	}
	str := &bytes.Buffer{}
	if _, err := str.ReadFrom(txt); err != nil {
		return txt
	}
	buf := &bytes.Buffer{}
	if _, err := strings.NewReplacer(
		CRLF.Code(), opt.Code(),
		LF.Code(), opt.Code(),
		CR.Code(), opt.Code(),
	).WriteString(buf, str.String()); err != nil {
		return txt
	}
	return buf
}

//String returns converted text string
func String(txt string, opt Option) string {
	if opt == Unknown {
		return txt
	}
	return strings.NewReplacer(
		CRLF.Code(), opt.Code(),
		LF.Code(), opt.Code(),
		CR.Code(), opt.Code(),
	).Replace(txt)
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
