package width

import (
	"strings"

	wdth "golang.org/x/text/width"
)

//Option is form of width transformer
type Option int

const (
	//Unknown form of width transformer
	Unknown Option = iota
	//Fold form of width transformer
	Fold
	//Narrow form of width transformer
	Narrow
	//Widen form of width transformer
	Widen
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
