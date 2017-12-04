package detect

import "github.com/pkg/errors"

var (
	//ErrNoImplement is error "no implementation"
	ErrNoImplement = errors.New("no implementation")
)

//CharEncoding is type of character encoding
type CharEncoding int

const (
	//Unknown is unknown character
	Unknown CharEncoding = iota
	//UTF8 is UTF-8
	UTF8
	//ISO8859L1 is ISO-8859-1 (Latin-1)
	ISO8859L1
	//ShiftJIS is Shift-JIS
	ShiftJIS
	//EUCJP is EUC-JP
	EUCJP
	//ISO2022JP is ISO-2022-JP
	ISO2022JP
	//EUCKR is EUC-KR
	EUCKR
	//GB18030 is GB-18030
	GB18030
	//Big5 is Big5
	Big5
)

var encodingMap = map[string]CharEncoding{
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
	if e, ok := encodingMap[s]; ok {
		return e
	}
	return Unknown

}

func (e CharEncoding) String() string {
	for key, value := range encodingMap {
		if value == e {
			return key
		}
	}
	return "Unknown"
}
