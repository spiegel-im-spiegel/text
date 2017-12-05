package normalize

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestDo(t *testing.T) {
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
		res := Do(bytes.NewBufferString(tst.txt), tst.opt)
		buf := new(bytes.Buffer)
		io.Copy(buf, res)
		if buf.String() != tst.res {
			t.Errorf("Do(%v)  = \"%v\", want \"%v\".", tst.txt, buf.Bytes(), []byte(tst.res))
		}
	}
}

func ExampleDo() {
	res := Do(bytes.NewBufferString("ﾍﾟﾝｷﾞﾝ"), NFKC)
	buf := new(bytes.Buffer)
	io.Copy(buf, res)
	fmt.Println(buf)
	// Output:
	// ペンギン
}
