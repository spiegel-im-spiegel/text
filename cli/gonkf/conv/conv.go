package conv

import (
	"io"

	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/text/convert"
	"github.com/spiegel-im-spiegel/text/decode"
	"github.com/spiegel-im-spiegel/text/detect"
	"github.com/spiegel-im-spiegel/text/ecode"
	"github.com/spiegel-im-spiegel/text/encode"
)

//Run return converted text
func Run(txt io.Reader, opt *Options) (io.Reader, error) {
	if opt.SrcEncoding() == opt.DstEncoding() {
		return txt, nil
	}
	if opt.DstEncoding() == detect.Unknown {
		return txt, errs.Wrap(ecode.ErrNoImplement, "no character encoding of destination text")
	}
	var dst io.Reader
	var err error
	switch opt.SrcEncoding() {
	case detect.UTF8:
		dst, err = encode.FromUTF8To(opt.DstEncoding(), txt)
	case detect.Unknown:
		if opt.DstEncoding() == detect.UTF8 {
			dst, err = decode.ToUTF8ja(txt)
		} else {
			dst, err = convert.ToJa(opt.DstEncoding(), txt)
		}
	default:
		if opt.DstEncoding() == detect.UTF8 {
			dst, err = decode.ToUTF8From(opt.SrcEncoding(), txt)
		} else {
			dst, err = convert.FromTo(opt.SrcEncoding(), opt.DstEncoding(), txt)
		}
	}
	return dst, err
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
