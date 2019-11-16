package width

import (
	"strings"

	wdth "golang.org/x/text/width"
)

//Option is form of width transformer
type Option int

const (
	Unknown Option = iota //Unknown form of width transformer
	Fold                  //Fold form of width transformer
	Narrow                //Narrow form of width transformer
	Widen                 //Widen form of width transformer
)

var widthNameMap = map[string]Option{
	"fold":   Fold,
	"narrow": Narrow,
	"widen":  Widen,
}

var widthMap = map[Option]wdth.Transformer{
	Fold:   wdth.Fold,
	Narrow: wdth.Narrow,
	Widen:  wdth.Widen,
}

//FormofWidth returns form of width transformer
func FormofWidth(s string) Option {
	if w, ok := widthNameMap[strings.ToLower(s)]; ok {
		return w
	}
	return Unknown
}

func (w Option) String() string {
	for key, value := range widthNameMap {
		if value == w {
			return key
		}
	}
	return "unknown"
}

//GetForm returns transform.Transformer instance
func (w Option) GetForm() *wdth.Transformer {
	if f, ok := widthMap[w]; ok {
		return &f
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
