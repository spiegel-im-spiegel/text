package convert

import (
	"github.com/spiegel-im-spiegel/text/decode"
	"github.com/spiegel-im-spiegel/text/detect"
	"github.com/spiegel-im-spiegel/text/encode"
)

//FromTo returns converted text
func FromTo(from, to detect.CharEncoding, txt []byte) ([]byte, error) {
	utf8Txt, err := decode.ToUTF8From(from, txt)
	if err != nil {
		return nil, err
	}
	if to == detect.UTF8 || to == detect.ISO8859L1 {
		return utf8Txt, nil
	}
	return encode.FromUTF8To(to, utf8Txt)
}

//To returns converted text
func To(to detect.CharEncoding, txt []byte) ([]byte, error) {
	from := detect.EncodingBest(txt)
	if from == detect.Unknown {
		return nil, detect.ErrNoImplement
	}
	return FromTo(from, to, txt)
}

//ToJa returns converted text (japanese encoding only)
func ToJa(to detect.CharEncoding, txt []byte) ([]byte, error) {
	from := detect.EncodingJa(txt)
	if from == detect.Unknown {
		return nil, detect.ErrNoImplement
	}
	return FromTo(from, to, txt)
}
