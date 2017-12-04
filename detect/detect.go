package detect

import (
	"github.com/saintfish/chardet"
)

//Encoding returns CharEncoding of text
func Encoding(txt []byte, lang string) CharEncoding {
	//if utf8.Valid(txt) {
	//	return UTF8
	//}
	all, err := chardet.NewTextDetector().DetectAll(txt)
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

//EncodingBest returns CharEncoding of text (best selection)
func EncodingBest(txt []byte) CharEncoding {
	return Encoding(txt, "")
}

//EncodingJa returns CharEncoding of text (Japanese only)
func EncodingJa(txt []byte) CharEncoding {
	return Encoding(txt, "ja")
}
