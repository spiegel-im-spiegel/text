package conv

import (
	"io"

	"github.com/pkg/errors"
	"github.com/spiegel-im-spiegel/text/convert"
	"github.com/spiegel-im-spiegel/text/decode"
	"github.com/spiegel-im-spiegel/text/detect"
	"github.com/spiegel-im-spiegel/text/encode"
)

//Run return converted text
func Run(txt io.Reader, opt *Options) (io.Reader, error) {
	if opt.SrcEncoding() == opt.DstEncoding() {
		return txt, nil
	}
	if opt.DstEncoding() == detect.Unknown {
		return txt, errors.New("no character encoding of destination text")
	}
	var dst io.Reader
	var err error
	switch opt.SrcEncoding() {
	case detect.UTF8:
		dst, err = encode.FromUTF8To(opt.DstEncoding(), txt)
	case detect.Unknown:
		if opt.DstEncoding() == detect.UTF8 {
			dst, err = decode.ToUTF8ja(txt)
		} else {
			dst, err = convert.ToJa(opt.DstEncoding(), txt)
		}
	default:
		if opt.DstEncoding() == detect.UTF8 {
			dst, err = decode.ToUTF8From(opt.SrcEncoding(), txt)
		} else {
			dst, err = convert.FromTo(opt.SrcEncoding(), opt.DstEncoding(), txt)
		}
	}
	return dst, err
}
