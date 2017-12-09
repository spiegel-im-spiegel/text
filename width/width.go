package width

import (
	"io"

	"golang.org/x/text/transform"
)

//Reader returns transformed text reader
func Reader(txt io.Reader, opt Option) io.Reader {
	t := opt.GetForm()
	if t == nil {
		return txt
	}
	return transform.NewReader(txt, t)
}

//String returns normalized text reader
func String(txt string, opt Option) string {
	t := opt.GetForm()
	if t == nil {
		return txt
	}
	return t.String(txt)
}
