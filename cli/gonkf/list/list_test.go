package list

import (
	"testing"

	"github.com/spiegel-im-spiegel/text/detect"
)

func TestGetEncoding(t *testing.T) {
	testCase := []struct {
		e detect.CharEncoding
		s string
	}{
		{e: detect.Unknown, s: "foobar"},
		{e: detect.UTF8, s: "UTF8"},
		{e: detect.ShiftJIS, s: "sjis"},
		{e: detect.EUCJP, s: "EUCJP"},
		{e: detect.ISO2022JP, s: "jis"},
		{e: detect.EUCKR, s: "EUCKR"},
		{e: detect.GB18030, s: "GB18030"},
		{e: detect.Big5, s: "Big5"},
	}

	for _, tst := range testCase {
		e := TypeofEncoding(tst.s)
		if e != tst.e {
			t.Errorf("TypeofEncoding(%v)  = \"%v\", want \"%v\".", tst.s, e, tst.e)
		}
	}
}
