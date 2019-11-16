package convert

import (
	"bytes"
	"io"

	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/text/decode"
	"github.com/spiegel-im-spiegel/text/detect"
	"github.com/spiegel-im-spiegel/text/ecode"
	"github.com/spiegel-im-spiegel/text/encode"
)

//FromTo returns converted text
func FromTo(from, to detect.CharEncoding, txt io.Reader) (io.Reader, error) {
	if from == to {
		return txt, nil
	}
	utf8Txt, err := decode.ToUTF8From(from, txt)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}
	if to == detect.UTF8 || to == detect.ISO8859L1 {
		return utf8Txt, nil
	}
	return encode.FromUTF8To(to, utf8Txt)
}

//To returns converted text (auto detection of encoding)
func To(to detect.CharEncoding, txt io.Reader) (io.Reader, error) {
	buf := &bytes.Buffer{}
	tee := io.TeeReader(txt, buf)
	from := detect.EncodingBest(tee)
	if from == detect.Unknown {
		return nil, errs.Wrap(ecode.ErrNoImplement, "")
	}
	return FromTo(from, to, buf)
}

//ToJa returns converted text (auto detection of japanese encoding)
func ToJa(to detect.CharEncoding, txt io.Reader) (io.Reader, error) {
	buf := &bytes.Buffer{}
	tee := io.TeeReader(txt, buf)
	from := detect.EncodingJa(tee)
	if from == detect.Unknown {
		return nil, errs.Wrap(ecode.ErrNoImplement, "")
	}
	return FromTo(from, to, buf)
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
