package newline

import "strings"

//Option is type of newline
type Option int

const (
	Unknown Option = iota //Unknown is unknown type
	LF                    //LF is LF only
	CR                    //CR is CR only
	CRLF                  //CRLF is CR+LF
)

var newlineMap = map[string]Option{
	"lf":   LF,
	"cr":   CR,
	"crlf": CRLF,
}

var newlineCodeMap = map[Option]string{
	LF:   "\n",
	CR:   "\r",
	CRLF: "\r\n",
}

//TypeofNewline returns type of newline
func TypeofNewline(s string) Option {
	if e, ok := newlineMap[strings.ToLower(s)]; ok {
		return e
	}
	return Unknown
}

//Code returns character code
func (nl Option) Code() string {
	if c, ok := newlineCodeMap[nl]; ok {
		return c
	}
	return ""
}

func (nl Option) String() string {
	for key, value := range newlineMap {
		if value == nl {
			return key
		}
	}
	return "Unknown"
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
