package newline

import (
	"bytes"
	"io"
	"strings"
)

//Reader returns converted text Reader
func Reader(txt io.Reader, opt Option) io.Reader {
	if opt == Unknown {
		return txt
	}
	str := new(bytes.Buffer)
	buf := new(bytes.Buffer)
	io.Copy(str, txt)
	strings.NewReplacer(
		CRLF.Code(), opt.Code(),
		LF.Code(), opt.Code(),
		CR.Code(), opt.Code(),
	).WriteString(buf, str.String())
	return buf
}

//String returns converted text string
func String(txt string, opt Option) string {
	if opt == Unknown {
		return txt
	}
	return strings.NewReplacer(
		CRLF.Code(), opt.Code(),
		LF.Code(), opt.Code(),
		CR.Code(), opt.Code(),
	).Replace(txt)
}
