package width

import (
	"testing"
)

func TestOptionString(t *testing.T) {
	testCase := []struct {
		e Option
		s string
	}{
		{e: Unknown, s: "unknown"},
		{e: Fold, s: "fold"},
		{e: Narrow, s: "narrow"},
		{e: Widen, s: "widen"},
		{e: Option(4), s: "unknown"},
	}

	for _, tst := range testCase {
		if tst.e.String() != tst.s {
			t.Errorf("Option(%d)  = \"%v\", want \"%v\".", int(tst.e), tst.e, tst.s)
		}
	}
}

func TestGetOption(t *testing.T) {
	testCase := []struct {
		e Option
		s string
	}{
		{e: Unknown, s: "unknown"},
		{e: Fold, s: "fold"},
		{e: Narrow, s: "narrow"},
		{e: Widen, s: "widen"},
	}

	for _, tst := range testCase {
		e := FormofWidth(tst.s)
		if e != tst.e {
			t.Errorf("typeofEncoding(%v)  = \"%v\", want \"%v\".", tst.s, e, tst.e)
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
