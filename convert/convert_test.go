package convert

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/spiegel-im-spiegel/text/detect"
)

func TestToJa(t *testing.T) {
	testCase := []struct {
		from detect.CharEncoding
		to   detect.CharEncoding
		txt  []byte
		res  []byte
	}{
		{
			from: detect.ISO2022JP,
			to:   detect.ShiftJIS,
			txt:  []byte{0x1b, 0x24, 0x42, 0x24, 0x33, 0x24, 0x73, 0x24, 0x4b, 0x24, 0x41, 0x24, 0x4f, 0x40, 0x24, 0x33, 0x26, 0x1b, 0x28, 0x42},
			res:  []byte{0x82, 0xb1, 0x82, 0xf1, 0x82, 0xc9, 0x82, 0xbf, 0x82, 0xcd, 0x90, 0xa2, 0x8a, 0x45},
		},
		{
			from: detect.ISO2022JP,
			to:   detect.UTF8,
			txt:  []byte{0x1b, 0x24, 0x42, 0x24, 0x33, 0x24, 0x73, 0x24, 0x4b, 0x24, 0x41, 0x24, 0x4f, 0x40, 0x24, 0x33, 0x26, 0x1b, 0x28, 0x42},
			res:  []byte("こんにちは世界"),
		},
	}

	for _, tst := range testCase {
		res, err := ToJa(tst.to, bytes.NewReader(tst.txt))
		if err != nil {
			t.Errorf("ToJa(%v)  = \"%v\", want nil.", tst.txt, err)
		}
		buf := new(bytes.Buffer)
		io.Copy(buf, res)
		if bytes.Compare(buf.Bytes(), tst.res) != 0 {
			t.Errorf("ToJa(%v)  = \"%v\", want \"%v\".", tst.txt, buf, tst.res)
		}
	}
}

func ExampleFromTo() {
	jisText := []byte{0x1b, 0x24, 0x42, 0x24, 0x33, 0x24, 0x73, 0x24, 0x4b, 0x24, 0x41, 0x24, 0x4f, 0x40, 0x24, 0x33, 0x26, 0x1b, 0x28, 0x42}
	res, err := FromTo(detect.ISO2022JP, detect.UTF8, bytes.NewReader(jisText))
	if err != nil {
		fmt.Println(err)
		return
	}
	buf := new(bytes.Buffer)
	io.Copy(buf, res)
	fmt.Println(buf)
	// Output:
	// こんにちは世界
}
