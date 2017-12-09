package normalize

import (
	"io"
)

//Reader returns normalized text reader
func Reader(txt io.Reader, opt Option) io.Reader {
	if opt == Unknown {
		return txt
	}
	return opt.GetForm().Reader(txt)
}

//String returns normalized text string
func String(txt string, opt Option) string {
	if opt == Unknown {
		return txt
	}
	return opt.GetForm().String(txt)
}
