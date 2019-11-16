package normalize

import (
	"testing"

	"golang.org/x/text/unicode/norm"
)

func TestOptionString(t *testing.T) {
	testCase := []struct {
		e Option
		s string
	}{
		{e: Unknown, s: "Unknown"},
		{e: NFC, s: "nfc"},
		{e: NFD, s: "nfd"},
		{e: NFKC, s: "nfkc"},
		{e: NFKD, s: "nfkd"},
		{e: Option(5), s: "Unknown"},
	}

	for _, tst := range testCase {
		if tst.e.String() != tst.s {
			t.Errorf("Encoding(%d)  = \"%v\", want \"%v\".", int(tst.e), tst.e, tst.s)
		}
	}
}

func TestTypeofNormalize(t *testing.T) {
	testCase := []struct {
		e Option
		s string
	}{
		{e: Unknown, s: "foobar"},
		{e: NFC, s: "nfc"},
		{e: NFD, s: "nfd"},
		{e: NFKC, s: "nfkc"},
		{e: NFKD, s: "nfkd"},
	}

	for _, tst := range testCase {
		n := FormofNormalize(tst.s)
		if n != tst.e {
			t.Errorf("TypeofNormalize(%s)  = \"%v\", want \"%v\".", tst.s, n, tst.e)
		}
	}
}

func TestGetForm(t *testing.T) {
	testCase := []struct {
		e    Option
		form norm.Form
	}{
		{e: Unknown, form: norm.Form(-1)},
		{e: NFC, form: norm.NFC},
		{e: NFD, form: norm.NFD},
		{e: NFKC, form: norm.NFKC},
		{e: NFKD, form: norm.NFKD},
	}

	for _, tst := range testCase {
		form := tst.e.GetForm()
		if form != tst.form {
			t.Errorf("GetEncoding(%v)  = \"%v\", want \"%v\".", tst.e, form, tst.form)
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
