package newline

import "testing"

func TestOptionString(t *testing.T) {
	testCase := []struct {
		nl Option
		s  string
	}{
		{nl: Unknown, s: "Unknown"},
		{nl: LF, s: "lf"},
		{nl: CR, s: "cr"},
		{nl: CRLF, s: "crlf"},
		{nl: Option(4), s: "Unknown"},
	}

	for _, tst := range testCase {
		if tst.nl.String() != tst.s {
			t.Errorf("Option(%d)  = \"%v\", want \"%v\".", int(tst.nl), tst.nl, tst.s)
		}
	}
}

func TestTypeofNewline(t *testing.T) {
	testCase := []struct {
		nl Option
		s  string
	}{
		{nl: Unknown, s: "foobar"},
		{nl: LF, s: "lf"},
		{nl: CR, s: "cr"},
		{nl: CRLF, s: "crlf"},
	}

	for _, tst := range testCase {
		nl := TypeofNewline(tst.s)
		if nl != tst.nl {
			t.Errorf("TypeofNewline(%v)  = \"%v\", want \"%v\".", tst.s, nl, tst.nl)
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
