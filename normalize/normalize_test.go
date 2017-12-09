package normalize

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestReader(t *testing.T) {
	testCase := []struct {
		opt Option
		txt string
		res string
	}{
		{opt: Unknown, txt: "ペンギン", res: "ペンギン"},
		{opt: NFC, txt: "ペンギン", res: "ペンギン"},
		{opt: NFD, txt: "ペンギン", res: "ペンギン"},
		{opt: NFKC, txt: "ﾍﾟﾝｷﾞﾝ", res: "ペンギン"},
		{opt: NFKD, txt: "ﾍﾟﾝｷﾞﾝ", res: "ペンギン"},
	}

	for _, tst := range testCase {
		res := Reader(bytes.NewBufferString(tst.txt), tst.opt)
		buf := new(bytes.Buffer)
		io.Copy(buf, res)
		if buf.String() != tst.res {
			t.Errorf("Do(%v)  = \"%U\", want \"%U\".", tst.txt, []rune(buf.String()), []rune(tst.res))
		}
	}
}

func TestString(t *testing.T) {
	testCase := []struct {
		opt Option
		txt string
		res string
	}{
		{opt: Unknown, txt: "ペンギン", res: "ペンギン"},
		{opt: NFC, txt: "ペンギン", res: "ペンギン"},
		{opt: NFD, txt: "ペンギン", res: "ペンギン"},
		{opt: NFKC, txt: "ﾍﾟﾝｷﾞﾝ", res: "ペンギン"},
		{opt: NFKD, txt: "ﾍﾟﾝｷﾞﾝ", res: "ペンギン"},
	}

	for _, tst := range testCase {
		res := String(tst.txt, tst.opt)
		if res != tst.res {
			t.Errorf("Do(%v)  = \"%U\", want \"%U\".", tst.txt, []rune(res), []rune(tst.res))
		}
	}
}

func ExampleReader() {
	res := Reader(bytes.NewBufferString("ﾍﾟﾝｷﾞﾝ"), NFKC)
	buf := new(bytes.Buffer)
	io.Copy(buf, res)
	fmt.Println(buf)
	// Output:
	// ペンギン
}

func ExampleString() {
	res := String("ﾍﾟﾝｷﾞﾝ", NFKC)
	fmt.Println(res)
	// Output:
	// ペンギン
}
