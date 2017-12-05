package list

import (
	"sort"
	"strings"

	"github.com/spiegel-im-spiegel/text/detect"
	"github.com/spiegel-im-spiegel/text/newline"
)

var encodingMap = map[string]detect.CharEncoding{
	"utf8":    detect.UTF8,
	"sjis":    detect.ShiftJIS,
	"eucjp":   detect.EUCJP,
	"jis":     detect.ISO2022JP,
	"euckr":   detect.EUCKR,
	"gb18030": detect.GB18030,
	"big5":    detect.Big5,
}

//TypeofEncoding returns character encoding from string
func TypeofEncoding(s string) detect.CharEncoding {
	if e, ok := encodingMap[strings.ToLower(s)]; ok {
		return e
	}
	return detect.Unknown
}

//AvailableEncodingList return available character encoding list
func AvailableEncodingList() string {
	var names []string
	for key := range encodingMap {
		names = append(names, key)
	}
	sort.Strings(names)
	return strings.Join(names, " ")

}

var newlineNames = []string{
	newline.LF.String(),
	newline.CR.String(),
	newline.CRLF.String(),
}

//AvailableNewlineOptionsList return available newline options list
func AvailableNewlineOptionsList() string {
	return strings.Join(newlineNames, " ")
}
