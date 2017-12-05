package decode

import (
	"bytes"
	"io"

	"github.com/spiegel-im-spiegel/text"
	"github.com/spiegel-im-spiegel/text/detect"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

//ToUTF8From returns UTF-8 text from other encoding text
func ToUTF8From(e detect.CharEncoding, txt io.Reader) (io.Reader, error) {
	var decoder transform.Transformer
	switch e {
	case detect.UTF8, detect.ISO8859L1:
		return txt, nil
	case detect.ShiftJIS:
		decoder = japanese.ShiftJIS.NewDecoder()
	case detect.EUCJP:
		decoder = japanese.EUCJP.NewDecoder()
	case detect.ISO2022JP:
		decoder = japanese.ISO2022JP.NewDecoder()
	case detect.EUCKR:
		decoder = korean.EUCKR.NewDecoder()
	case detect.GB18030:
		decoder = simplifiedchinese.GB18030.NewDecoder()
	case detect.Big5:
		decoder = traditionalchinese.Big5.NewDecoder()
	default:
		return nil, text.ErrNoImplement
	}
	buf := new(bytes.Buffer)
	_, err := io.Copy(buf, transform.NewReader(txt, decoder))
	return buf, err
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
