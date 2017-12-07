package wdth

import (
	"io"

	"github.com/spiegel-im-spiegel/text/width"
)

//Run return converted text
func Run(txt io.Reader, form width.Option) io.Reader {
	if form == width.Unknown {
		return txt
	}
	return width.Do(txt, form)
}
