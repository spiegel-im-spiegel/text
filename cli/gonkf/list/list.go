package list

import (
	"bytes"
	"fmt"
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
	sep := ""
	buf := new(bytes.Buffer)
	for k := range encodingMap {
		fmt.Fprintf(buf, "%s%s", sep, k)
		sep = " "
	}
	return buf.String()
}

var newlineList = []newline.Option{
	newline.LF,
	newline.CR,
	newline.CRLF,
}

//AvailableNewlineOptionsList return available newline options list
func AvailableNewlineOptionsList() string {
	sep := ""
	buf := new(bytes.Buffer)
	for nl := range newlineList {
		fmt.Fprintf(buf, "%s%v", sep, nl)
		sep = " "
	}
	return buf.String()
}
