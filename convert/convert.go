package convert

import (
	"bytes"
	"io"

	"github.com/spiegel-im-spiegel/text"
	"github.com/spiegel-im-spiegel/text/decode"
	"github.com/spiegel-im-spiegel/text/detect"
	"github.com/spiegel-im-spiegel/text/encode"
)

//FromTo returns converted text
func FromTo(from, to detect.CharEncoding, txt io.Reader) (io.Reader, error) {
	if from == to {
		return txt, nil
	}
	utf8Txt, err := decode.ToUTF8From(from, txt)
	if err != nil {
		return nil, err
	}
	if to == detect.UTF8 || to == detect.ISO8859L1 {
		return utf8Txt, nil
	}
	return encode.FromUTF8To(to, utf8Txt)
}

//To returns converted text (auto detection of encoding)
func To(to detect.CharEncoding, txt io.Reader) (io.Reader, error) {
	buf := new(bytes.Buffer)
	tee := io.TeeReader(txt, buf)
	from := detect.EncodingBest(tee)
	if from == detect.Unknown {
		return nil, text.ErrNoImplement
	}
	return FromTo(from, to, buf)
}

//ToJa returns converted text (auto detection of japanese encoding)
func ToJa(to detect.CharEncoding, txt io.Reader) (io.Reader, error) {
	buf := new(bytes.Buffer)
	tee := io.TeeReader(txt, buf)
	from := detect.EncodingJa(tee)
	if from == detect.Unknown {
		return nil, text.ErrNoImplement
	}
	return FromTo(from, to, buf)
}
