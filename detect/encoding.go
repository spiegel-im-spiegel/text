package detect

import (
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/encoding/unicode"
)

//CharEncoding is type of character encoding
type CharEncoding int

const (
	Unknown   CharEncoding = iota //Unknown is unknown character encoding
	UTF8                          //UTF8 is UTF-8 encoding
	ISO8859L1                     //ISO8859L1 is ISO-8859-1 (Latin-1) encoding
	ShiftJIS                      //ShiftJIS is Shift-JIS encoding
	EUCJP                         //EUCJP is EUC-JP encoding
	ISO2022JP                     //ISO2022JP is ISO-2022-JP encoding
	EUCKR                         //EUCKR is EUC-KR encoding
	GB18030                       //GB18030 is GB-18030 encoding
	Big5                          //Big5 is Big5 encoding
)

var encodingMap = map[CharEncoding]encoding.Encoding{
	UTF8:      unicode.UTF8,
	ISO8859L1: charmap.ISO8859_1,
	ShiftJIS:  japanese.ShiftJIS,
	EUCJP:     japanese.EUCJP,
	ISO2022JP: japanese.ISO2022JP,
	EUCKR:     korean.EUCKR,
	GB18030:   simplifiedchinese.GB18030,
	Big5:      traditionalchinese.Big5,
}

var encodingNameMap = map[string]CharEncoding{
	"UTF-8":       UTF8,
	"ISO-8859-1":  ISO8859L1,
	"Shift_JIS":   ShiftJIS,
	"EUC-JP":      EUCJP,
	"ISO-2022-JP": ISO2022JP,
	"EUC-KR":      EUCKR,
	"GB-18030":    GB18030,
	"Big5":        Big5,
}

func typeofEncoding(s string) CharEncoding {
	if e, ok := encodingNameMap[s]; ok {
		return e
	}
	return Unknown
}

func (e CharEncoding) String() string {
	for key, value := range encodingNameMap {
		if value == e {
			return key
		}
	}
	return "Unknown"
}

//GetEncoding returns Encoding instance
func (e CharEncoding) GetEncoding() encoding.Encoding {
	for key, value := range encodingMap {
		if key == e {
			return value
		}
	}
	return nil
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
