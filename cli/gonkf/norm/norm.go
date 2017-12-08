package norm

import (
	"io"

	"github.com/spiegel-im-spiegel/text/normalize"
)

//Run returns normalized text
func Run(txt io.Reader, form normalize.Option) io.Reader {
	return normalize.Reader(txt, form)
}
