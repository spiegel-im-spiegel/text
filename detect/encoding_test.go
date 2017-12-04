package detect

import "testing"

func TestCharEncodingString(t *testing.T) {
	testCase := []struct {
		e CharEncoding
		s string
	}{
		{e: Unknown, s: "Unknown"},
		{e: UTF8, s: "UTF-8"},
		{e: ISO8859L1, s: "ISO-8859-1"},
		{e: ShiftJIS, s: "Shift_JIS"},
		{e: EUCJP, s: "EUC-JP"},
		{e: ISO2022JP, s: "ISO-2022-JP"},
		{e: EUCKR, s: "EUC-KR"},
		{e: GB18030, s: "GB-18030"},
		{e: Big5, s: "Big5"},
		{e: CharEncoding(9), s: "Unknown"},
	}

	for _, tst := range testCase {
		if tst.e.String() != tst.s {
			t.Errorf("Encoding(%d)  = \"%v\", want \"%v\".", int(tst.e), tst.e, tst.s)
		}
	}
}

func TestGetCharEncoding(t *testing.T) {
	testCase := []struct {
		e CharEncoding
		s string
	}{
		{e: Unknown, s: "foobar"},
		{e: UTF8, s: "UTF-8"},
		{e: ISO8859L1, s: "ISO-8859-1"},
		{e: ShiftJIS, s: "Shift_JIS"},
		{e: EUCJP, s: "EUC-JP"},
		{e: ISO2022JP, s: "ISO-2022-JP"},
		{e: EUCKR, s: "EUC-KR"},
		{e: GB18030, s: "GB-18030"},
		{e: Big5, s: "Big5"},
	}

	for _, tst := range testCase {
		e := typeofEncoding(tst.s)
		if e != tst.e {
			t.Errorf("typeofEncoding(%v)  = \"%v\", want \"%v\".", tst.s, e, tst.e)
		}
	}
}
