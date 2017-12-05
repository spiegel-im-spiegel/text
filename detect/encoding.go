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
	//Unknown is unknown character encoding
	Unknown CharEncoding = iota
	//UTF8 is UTF-8 encoding
	UTF8
	//ISO8859L1 is ISO-8859-1 (Latin-1) encoding
	ISO8859L1
	//ShiftJIS is Shift-JIS encoding
	ShiftJIS
	//EUCJP is EUC-JP encoding
	EUCJP
	//ISO2022JP is ISO-2022-JP encoding
	ISO2022JP
	//EUCKR is EUC-KR encoding
	EUCKR
	//GB18030 is GB-18030 encoding
	GB18030
	//Big5 is Big5 encoding
	Big5
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
