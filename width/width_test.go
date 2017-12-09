package width

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestWidthReader(t *testing.T) {
	testCase := []struct {
		opt Option
		txt string
		res string
	}{
		{opt: Unknown, txt: "abｦ￦￮￥Ａ", res: "abｦ￦￮￥Ａ"},
		{opt: Fold, txt: "abｦ￦￮￥Ａ", res: "abヲ₩○¥A"},
		{opt: Narrow, txt: "abｦ￦￮￥Ａ", res: "abｦ₩￮¥A"},
		{opt: Widen, txt: "abｦ￦￮￥Ａ", res: "ａｂヲ￦○￥Ａ"},
	}

	for _, tst := range testCase {
		res := Reader(bytes.NewBufferString(tst.txt), tst.opt)
		buf := new(bytes.Buffer)
		io.Copy(buf, res)
		if buf.String() != tst.res {
			t.Errorf("Do(%v)  = \"%v\", want \"%v\".", tst.txt, buf.String(), tst.res)
		}
	}
}

func TestWidthString(t *testing.T) {
	testCase := []struct {
		opt Option
		txt string
		res string
	}{
		{opt: Unknown, txt: "abｦ￦￮￥Ａ", res: "abｦ￦￮￥Ａ"},
		{opt: Fold, txt: "abｦ￦￮￥Ａ", res: "abヲ₩○¥A"},
		{opt: Narrow, txt: "abｦ￦￮￥Ａ", res: "abｦ₩￮¥A"},
		{opt: Widen, txt: "abｦ￦￮￥Ａ", res: "ａｂヲ￦○￥Ａ"},
	}

	for _, tst := range testCase {
		res := String(tst.txt, tst.opt)
		if res != tst.res {
			t.Errorf("Do(%v)  = \"%v\", want \"%v\".", tst.txt, res, tst.res)
		}
	}
}

func ExampleReader() {
	res := Reader(bytes.NewBufferString("abｦ￦￮￥Ａ"), Fold)
	buf := new(bytes.Buffer)
	io.Copy(buf, res)
	fmt.Println(buf)
	// Output:
	// abヲ₩○¥A
}

func ExampleString() {
	res := String("abｦ￦￮￥Ａ", Fold)
	fmt.Println(res)
	// Output:
	// abヲ₩○¥A
}
