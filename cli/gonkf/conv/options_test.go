package conv

import (
	"testing"

	"github.com/spiegel-im-spiegel/text/detect"
	"github.com/spiegel-im-spiegel/text/newline"
)

func TestSetSrcEncoding(t *testing.T) {
	opt := Options{}
	opt.SetSrcEncoding(detect.UTF8)
	if opt.SrcEncoding() != detect.UTF8 {
		t.Errorf("SetSrcEncoding()  = \"%v\", want \"%v\".", opt.SrcEncoding(), detect.UTF8)
	}

}

func TestSetDstEncoding(t *testing.T) {
	opt := Options{}
	opt.SetDstEncoding(detect.UTF8)
	if opt.DstEncoding() != detect.UTF8 {
		t.Errorf("SetSrcEncoding()  = \"%v\", want \"%v\".", opt.DstEncoding(), detect.UTF8)
	}

}

func TestSetNewline(t *testing.T) {
	opt := Options{}
	opt.SetNewline(newline.CRLF)
	if opt.Newline() != newline.CRLF {
		t.Errorf("SetSrcEncoding()  = \"%v\", want \"%v\".", opt.Newline(), newline.CRLF)
	}

}
