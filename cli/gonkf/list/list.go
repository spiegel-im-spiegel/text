package list

import (
	"sort"
	"strings"

	"github.com/spiegel-im-spiegel/text/detect"
	"github.com/spiegel-im-spiegel/text/newline"
	"github.com/spiegel-im-spiegel/text/normalize"
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
		if strings.ToLower(s) == strings.ToLower(e.String()) {
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
