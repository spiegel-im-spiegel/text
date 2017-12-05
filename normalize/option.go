package normalize

import (
	"strings"

	"golang.org/x/text/unicode/norm"
)

//Option is type of Unicode normalization
type Option int

const (
	//Unknown type of Unicode normalization
	Unknown Option = iota
	//NFC type of Unicode normalization
	NFC
	//NFD type of Unicode normalization
	NFD
	//NFKC type of Unicode normalization
	NFKC
	//NFKD type of Unicode normalization
	NFKD
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
