package normalize

import (
	"testing"

	"golang.org/x/text/unicode/norm"
)

func TestOptionString(t *testing.T) {
	testCase := []struct {
		e Option
		s string
	}{
		{e: Unknown, s: "Unknown"},
		{e: NFC, s: "nfc"},
		{e: NFD, s: "nfd"},
		{e: NFKC, s: "nfkc"},
		{e: NFKD, s: "nfkd"},
		{e: Option(5), s: "Unknown"},
	}

	for _, tst := range testCase {
		if tst.e.String() != tst.s {
			t.Errorf("Encoding(%d)  = \"%v\", want \"%v\".", int(tst.e), tst.e, tst.s)
		}
	}
}

func TestTypeofNormalize(t *testing.T) {
	testCase := []struct {
		e Option
		s string
	}{
		{e: Unknown, s: "foobar"},
		{e: NFC, s: "nfc"},
		{e: NFD, s: "nfd"},
		{e: NFKC, s: "nfkc"},
		{e: NFKD, s: "nfkd"},
	}

	for _, tst := range testCase {
		n := FormofNormalize(tst.s)
		if n != tst.e {
			t.Errorf("TypeofNormalize(%s)  = \"%v\", want \"%v\".", tst.s, n, tst.e)
		}
	}
}

func TestGetForm(t *testing.T) {
	testCase := []struct {
		e    Option
		form norm.Form
	}{
		{e: Unknown, form: norm.Form(-1)},
		{e: NFC, form: norm.NFC},
		{e: NFD, form: norm.NFD},
		{e: NFKC, form: norm.NFKC},
		{e: NFKD, form: norm.NFKD},
	}

	for _, tst := range testCase {
		form := tst.e.GetForm()
		if form != tst.form {
			t.Errorf("GetEncoding(%v)  = \"%v\", want \"%v\".", tst.e, form, tst.form)
		}
	}
}
