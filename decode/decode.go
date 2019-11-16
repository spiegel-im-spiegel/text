package decode

import (
	"bytes"
	"io"

	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/text/detect"
	"github.com/spiegel-im-spiegel/text/ecode"
	"golang.org/x/text/encoding"
)

//ToUTF8From returns UTF-8 text from other encoding text
func ToUTF8From(e detect.CharEncoding, txt io.Reader) (io.Reader, error) {
	var enc encoding.Encoding
	switch e {
	case detect.UTF8, detect.ISO8859L1:
		return txt, nil
	default:
		enc = e.GetEncoding()
	}
	if enc == nil {
		return nil, errs.Wrap(ecode.ErrNoImplement, "")
	}
	return enc.NewDecoder().Reader(txt), nil
}

//ToUTF8 returns UTF-8 text from other encoding text (auto detection of encoding)
func ToUTF8(txt io.Reader) (io.Reader, error) {
	buf := &bytes.Buffer{}
	tee := io.TeeReader(txt, buf)
	e := detect.EncodingBest(tee)
	if e == detect.Unknown {
		return nil, errs.Wrap(ecode.ErrNoImplement, "")
	}
	return ToUTF8From(e, buf)
}

//ToUTF8ja returns UTF-8 text from other encoding text (auto detection of japanses encoding)
func ToUTF8ja(txt io.Reader) (io.Reader, error) {
	buf := &bytes.Buffer{}
	tee := io.TeeReader(txt, buf)
	e := detect.EncodingJa(tee)
	if e == detect.Unknown {
		return nil, errs.Wrap(ecode.ErrNoImplement, "")
	}
	return ToUTF8From(e, buf)
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
