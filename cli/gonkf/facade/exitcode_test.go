package facade

import (
	"testing"
)

type ecodeTestCase struct { //Test case for ExitCode
	ec  ExitCode
	v   int
	str string
}

var ecodeTests = []ecodeTestCase{ //Test cases for ExitCode
	{ExitNormal, 0, "normal end"},
	{ExitAbnormal, 1, "abnormal end"},
	{ExitCode(2), 2, "unknown"},
}

func TestExitCode(t *testing.T) {
	for _, testCase := range ecodeTests {
		if testCase.ec.Int() != testCase.v {
			t.Errorf("ExitCode.Int()  = %v, want %v.", testCase.ec.Int(), testCase.v)
		}
		if testCase.ec.String() != testCase.str {
			t.Errorf("ExitCode.String()  = %v, want %v.", testCase.ec.String(), testCase.str)
		}
	}
}
