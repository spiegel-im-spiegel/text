package newline

import (
	"bytes"
	"io"
	"regexp"
)

var regxNewline = regexp.MustCompile(`\r\n|\r|\n`)

//Reader returns converted text Reader
func Reader(txt io.Reader, opt Option) io.Reader {
	rep := []byte{}
	switch opt {
	case LF:
		rep = []byte("\n")
	case CR:
		rep = []byte("\r")
	case CRLF:
		rep = []byte("\r\n")
	default:
		return txt
	}
	buf := new(bytes.Buffer)
	io.Copy(buf, txt)
	return bytes.NewReader(regxNewline.Copy().ReplaceAll(buf.Bytes(), rep))
}
