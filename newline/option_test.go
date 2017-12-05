package newline

import "testing"

func TestOptionString(t *testing.T) {
	testCase := []struct {
		nl Option
		s  string
	}{
		{nl: Unknown, s: "Unknown"},
		{nl: LF, s: "lf"},
		{nl: CR, s: "cr"},
		{nl: CRLF, s: "crlf"},
		{nl: Option(4), s: "Unknown"},
	}

	for _, tst := range testCase {
		if tst.nl.String() != tst.s {
			t.Errorf("Option(%d)  = \"%v\", want \"%v\".", int(tst.nl), tst.nl, tst.s)
		}
	}
}

func TestTypeofNewline(t *testing.T) {
	testCase := []struct {
		nl Option
		s  string
	}{
		{nl: Unknown, s: "foobar"},
		{nl: LF, s: "lf"},
		{nl: CR, s: "cr"},
		{nl: CRLF, s: "crlf"},
	}

	for _, tst := range testCase {
		nl := TypeofNewline(tst.s)
		if nl != tst.nl {
			t.Errorf("TypeofNewline(%v)  = \"%v\", want \"%v\".", tst.s, nl, tst.nl)
		}
	}
}
