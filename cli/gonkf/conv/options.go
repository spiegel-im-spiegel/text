package conv

import (
	"github.com/spiegel-im-spiegel/text/detect"
	"github.com/spiegel-im-spiegel/text/newline"
)

//Options class
type Options struct {
	from detect.CharEncoding
	to   detect.CharEncoding
	nl   newline.Option
}

//SetSrcEncoding is setting character encoding of source text
func (o *Options) SetSrcEncoding(e detect.CharEncoding) {
	o.from = e
}

//SetDstEncoding is setting character encoding of destination text
func (o *Options) SetDstEncoding(e detect.CharEncoding) {
	o.to = e
}

//SetNewline is setting type of newline
func (o *Options) SetNewline(nl newline.Option) {
	o.nl = nl
}

//SrcEncoding returns character encoding of source text
func (o Options) SrcEncoding() detect.CharEncoding {
	return o.from
}

//DstEncoding returns character encoding of destination text
func (o Options) DstEncoding() detect.CharEncoding {
	return o.to
}

//Newline returns type of newline
func (o Options) Newline() newline.Option {
	return o.nl
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
