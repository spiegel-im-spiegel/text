package decode

import (
	"bytes"
	"io"

	"github.com/spiegel-im-spiegel/text"
	"github.com/spiegel-im-spiegel/text/detect"
	"golang.org/x/text/encoding"
)

//ToUTF8From returns UTF-8 text from other encoding text
func ToUTF8From(e detect.CharEncoding, txt io.Reader) (io.Reader, error) {
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
	return enc.NewDecoder().Reader(txt), nil
}

//ToUTF8 returns UTF-8 text from other encoding text (auto detection of encoding)
func ToUTF8(txt io.Reader) (io.Reader, error) {
	buf := new(bytes.Buffer)
	tee := io.TeeReader(txt, buf)
	e := detect.EncodingBest(tee)
	if e == detect.Unknown {
		return nil, text.ErrNoImplement
	}
	return ToUTF8From(e, buf)
}

//ToUTF8ja returns UTF-8 text from other encoding text (auto detection of japanses encoding)
func ToUTF8ja(txt io.Reader) (io.Reader, error) {
	buf := new(bytes.Buffer)
	tee := io.TeeReader(txt, buf)
	e := detect.EncodingJa(tee)
	if e == detect.Unknown {
		return nil, text.ErrNoImplement
	}
	return ToUTF8From(e, buf)
}
