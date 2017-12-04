package decode

import (
	"bytes"
	"io"

	"github.com/spiegel-im-spiegel/text/detect"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

//ToUTF8From returns UTF-8 text from other encoding text
func ToUTF8From(e detect.CharEncoding, txt []byte) ([]byte, error) {
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
		return nil, detect.ErrNoImplement
	}
	buf := new(bytes.Buffer)
	_, err := io.Copy(buf, transform.NewReader(bytes.NewReader(txt), decoder))
	return buf.Bytes(), err
}

//ToUTF8 returns UTF-8 text from other encoding text (auto detection of encoding)
func ToUTF8(txt []byte) ([]byte, error) {
	e := detect.EncodingBest(txt)
	if e == detect.Unknown {
		return nil, detect.ErrNoImplement
	}
	return ToUTF8From(e, txt)
}

//ToUTF8ja returns UTF-8 text from other encoding text (auto detection of japanses encoding)
func ToUTF8ja(txt []byte) ([]byte, error) {
	e := detect.EncodingJa(txt)
	if e == detect.Unknown {
		return nil, detect.ErrNoImplement
	}
	return ToUTF8From(e, txt)
}
