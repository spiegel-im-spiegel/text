package normalize

import (
	"strings"

	"golang.org/x/text/unicode/norm"
)

//Option is type of Unicode normalization
type Option int

const (
	Unknown Option = iota //Unknown type of Unicode normalization
	NFC                   //NFC type of Unicode normalization
	NFD                   //NFD type of Unicode normalization
	NFKC                  //NFKC type of Unicode normalization
	NFKD                  //NFKD type of Unicode normalization
)

var normMap = map[Option]norm.Form{
	NFC:  norm.NFC,
	NFD:  norm.NFD,
	NFKC: norm.NFKC,
	NFKD: norm.NFKD,
}

var normNamesMap = map[string]Option{
	"nfc":  NFC,
	"nfd":  NFD,
	"nfkc": NFKC,
	"nfkd": NFKD,
}

//FormofNormalize returns Option
func FormofNormalize(s string) Option {
	if n, ok := normNamesMap[strings.ToLower(s)]; ok {
		return n
	}
	return Unknown
}

func (n Option) String() string {
	for key, value := range normNamesMap {
		if value == n {
			return key
		}
	}
	return "Unknown"
}

//GetForm returns norm.Form instance
func (n Option) GetForm() norm.Form {
	if f, ok := normMap[n]; ok {
		return f
	}
	return norm.Form(-1)

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
