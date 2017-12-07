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
		{e: detect.UTF8, s: "UTF-8"},
		{e: detect.ShiftJIS, s: "sjis"},
		{e: detect.ShiftJIS, s: "Shift_jis"},
		{e: detect.EUCJP, s: "EUC"},
		{e: detect.EUCJP, s: "EUC-JP"},
		{e: detect.ISO2022JP, s: "jis"},
		{e: detect.ISO2022JP, s: "iso-2022-jp"},
		//{e: detect.EUCKR, s: "EUCKR"},
		//{e: detect.GB18030, s: "GB18030"},
		//{e: detect.Big5, s: "Big5"},
	}

	for _, tst := range testCase {
		e := TypeofEncoding(tst.s)
		if e != tst.e {
			t.Errorf("TypeofEncoding(%v)  = \"%v\", want \"%v\".", tst.s, e, tst.e)
		}
	}
}

func TestAvailableEncodingList(t *testing.T) {
	str := AvailableEncodingList("|")
	//ref := "big5|euc|euckr|gb18030|jis|sjis|utf8"
	ref := "euc|jis|sjis|utf8"
	if str != ref {
		t.Errorf("AvailableEncodingList()  = \"%v\", want \"%v\".", str, ref)
	}
}

func TestAvailableNewlineOptionsList(t *testing.T) {
	str := AvailableNewlineOptionsList("|")
	ref := "lf|cr|crlf"
	if str != ref {
		t.Errorf("AvailableNewlineOptionsList()  = \"%v\", want \"%v\".", str, ref)
	}
}

func TestNormOptionsList(t *testing.T) {
	str := NormOptionsList("|")
	ref := "nfc|nfd|nfkc|nfkd"
	if str != ref {
		t.Errorf("NormOptionsList()  = \"%v\", want \"%v\".", str, ref)
	}
}

func TestWidthOptionsList(t *testing.T) {
	str := WidthOptionsList("|")
	ref := "fold|narrow|widen"
	if str != ref {
		t.Errorf("WidthOptionsList()  = \"%v\", want \"%v\".", str, ref)
	}
}
