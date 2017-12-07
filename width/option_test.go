package width

import (
	"testing"
)

func TestOptionString(t *testing.T) {
	testCase := []struct {
		e Option
		s string
	}{
		{e: Unknown, s: "unknown"},
		{e: Fold, s: "fold"},
		{e: Narrow, s: "narrow"},
		{e: Widen, s: "widen"},
		{e: Option(4), s: "unknown"},
	}

	for _, tst := range testCase {
		if tst.e.String() != tst.s {
			t.Errorf("Option(%d)  = \"%v\", want \"%v\".", int(tst.e), tst.e, tst.s)
		}
	}
}

func TestGetOption(t *testing.T) {
	testCase := []struct {
		e Option
		s string
	}{
		{e: Unknown, s: "unknown"},
		{e: Fold, s: "fold"},
		{e: Narrow, s: "narrow"},
		{e: Widen, s: "widen"},
	}

	for _, tst := range testCase {
		e := FormofWidth(tst.s)
		if e != tst.e {
			t.Errorf("typeofEncoding(%v)  = \"%v\", want \"%v\".", tst.s, e, tst.e)
		}
	}
}
