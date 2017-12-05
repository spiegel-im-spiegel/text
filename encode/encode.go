package encode

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

//FromUTF8To returns encoded text from UTF-8 text
func FromUTF8To(e detect.CharEncoding, txt io.Reader) (io.Reader, error) {
	var encoder transform.Transformer
	switch e {
	case detect.UTF8, detect.ISO8859L1:
		return txt, nil
	case detect.ShiftJIS:
		encoder = japanese.ShiftJIS.NewEncoder()
	case detect.EUCJP:
		encoder = japanese.EUCJP.NewEncoder()
	case detect.ISO2022JP:
		encoder = japanese.ISO2022JP.NewEncoder()
	case detect.EUCKR:
		encoder = korean.EUCKR.NewEncoder()
	case detect.GB18030:
		encoder = simplifiedchinese.GB18030.NewEncoder()
	case detect.Big5:
		encoder = traditionalchinese.Big5.NewEncoder()
	default:
		return nil, text.ErrNoImplement
	}
	buf := new(bytes.Buffer)
	_, err := io.Copy(transform.NewWriter(buf, encoder), txt)
	return buf, err
}
