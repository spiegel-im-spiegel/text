package encode

import (
	"bytes"
	"io"

	"github.com/spiegel-im-spiegel/text"
	"github.com/spiegel-im-spiegel/text/detect"
	"golang.org/x/text/encoding"
)

//FromUTF8To returns encoded text from UTF-8 text
func FromUTF8To(e detect.CharEncoding, txt io.Reader) (io.Reader, error) {
	var enc encoding.Encoding
	switch e {
	case detect.UTF8, detect.ISO8859L1:
		return txt, nil
	default:
		enc = e.GetEncoding()
	}
	if enc == nil {
		return nil, text.ErrNoImplement
	}
	buf := new(bytes.Buffer)
	_, err := io.Copy(enc.NewEncoder().Writer(buf), txt)
	return buf, err
}
