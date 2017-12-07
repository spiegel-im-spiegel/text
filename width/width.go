package width

import (
	"io"

	"golang.org/x/text/transform"
)

//Do returns transformed text
func Do(txt io.Reader, opt Option) io.Reader {
	t := opt.GetForm()
	if t == nil {
		return txt
	}
	return transform.NewReader(txt, t)
}
