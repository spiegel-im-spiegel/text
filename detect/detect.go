package detect

import (
	"bytes"
	"io"

	"github.com/saintfish/chardet"
)

//Encoding returns character encoding of text
func Encoding(txt io.Reader, lang string) CharEncoding {
	buf := new(bytes.Buffer)
	io.Copy(buf, txt)
	all, err := chardet.NewTextDetector().DetectAll(buf.Bytes())
	if err != nil {
		return Unknown
	}
	if all == nil {
		return Unknown
	}
	if len(lang) == 0 {
		return typeofEncoding(all[0].Charset)
	}
	for _, res := range all {
		//fmt.Println(res.Charset, res.Language)
		e := typeofEncoding(res.Charset)
		if (res.Language == lang || len(res.Language) == 0) && e != Unknown {
			return e
		}
	}
	return typeofEncoding(all[0].Charset)
}

//EncodingBest returns character encoding of text (best selection)
func EncodingBest(txt io.Reader) CharEncoding {
	return Encoding(txt, "")
}

//EncodingJa returns character encoding of text (Japanese only)
func EncodingJa(txt io.Reader) CharEncoding {
	return Encoding(txt, "ja")
}
