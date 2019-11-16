package conv

import (
	"testing"

	"github.com/spiegel-im-spiegel/text/detect"
	"github.com/spiegel-im-spiegel/text/newline"
)

func TestSetSrcEncoding(t *testing.T) {
	opt := Options{}
	opt.SetSrcEncoding(detect.UTF8)
	if opt.SrcEncoding() != detect.UTF8 {
		t.Errorf("SetSrcEncoding()  = \"%v\", want \"%v\".", opt.SrcEncoding(), detect.UTF8)
	}

}

func TestSetDstEncoding(t *testing.T) {
	opt := Options{}
	opt.SetDstEncoding(detect.UTF8)
	if opt.DstEncoding() != detect.UTF8 {
		t.Errorf("SetSrcEncoding()  = \"%v\", want \"%v\".", opt.DstEncoding(), detect.UTF8)
	}

}

func TestSetNewline(t *testing.T) {
	opt := Options{}
	opt.SetNewline(newline.CRLF)
	if opt.Newline() != newline.CRLF {
		t.Errorf("SetSrcEncoding()  = \"%v\", want \"%v\".", opt.Newline(), newline.CRLF)
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
