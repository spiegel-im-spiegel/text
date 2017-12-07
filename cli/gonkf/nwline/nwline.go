package nwline

import (
	"io"

	"github.com/spiegel-im-spiegel/text/newline"
)

//Run return converted text
func Run(txt io.Reader, form newline.Option) io.Reader {
	if form != newline.Unknown {
		return newline.Convert(txt, form)
	}
	return txt
}
