package list

import (
	"sort"
	"strings"

	"github.com/spiegel-im-spiegel/text/detect"
	"github.com/spiegel-im-spiegel/text/newline"
	"github.com/spiegel-im-spiegel/text/normalize"
	"github.com/spiegel-im-spiegel/text/width"
)

var encodingMap = map[string]detect.CharEncoding{
	"utf8": detect.UTF8,
	"sjis": detect.ShiftJIS,
	"euc":  detect.EUCJP,
	"jis":  detect.ISO2022JP,
	//"euckr":   detect.EUCKR,
	//"gb18030": detect.GB18030,
	//"big5":    detect.Big5,
}

//TypeofEncoding returns character encoding from string
func TypeofEncoding(s string) detect.CharEncoding {
	if e, ok := encodingMap[strings.ToLower(s)]; ok {
		return e
	}
	for _, e := range encodingMap {
		if strings.EqualFold(strings.ToLower(s), strings.ToLower(e.String())) {
			return e
		}
	}
	return detect.Unknown
}

//AvailableEncodingList return available character encoding list
func AvailableEncodingList(sep string) string {
	var names []string
	for key := range encodingMap {
		names = append(names, key)
	}
	sort.Strings(names)
	return strings.Join(names, sep)

}

var newlineNames = []string{
	newline.LF.String(),
	newline.CR.String(),
	newline.CRLF.String(),
}

//AvailableNewlineOptionsList return available newline options list
func AvailableNewlineOptionsList(sep string) string {
	return strings.Join(newlineNames, sep)
}

var normNames = []string{
	normalize.NFC.String(),
	normalize.NFD.String(),
	normalize.NFKC.String(),
	normalize.NFKD.String(),
}

//NormOptionsList return normalization form list
func NormOptionsList(sep string) string {
	return strings.Join(normNames, sep)
}

var widthNames = []string{
	width.Fold.String(),
	width.Narrow.String(),
	width.Widen.String(),
}

//WidthOptionsList return normalization form list
func WidthOptionsList(sep string) string {
	return strings.Join(widthNames, sep)
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
