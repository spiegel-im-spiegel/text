package guess

import (
	"io"

	"github.com/spiegel-im-spiegel/text/detect"
)

//Run returna character encoding of text
func Run(txt io.Reader) detect.CharEncoding {
	return detect.EncodingJa(txt)
}
