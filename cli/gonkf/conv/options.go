package conv

import (
	"github.com/spiegel-im-spiegel/text/detect"
	"github.com/spiegel-im-spiegel/text/newline"
)

//Options class
type Options struct {
	from detect.CharEncoding
	to   detect.CharEncoding
	nl   newline.Option
}

//SetSrcEncoding is setting character encoding of source text
func (o *Options) SetSrcEncoding(e detect.CharEncoding) {
	o.from = e
}

//SetDstEncoding is setting character encoding of destination text
func (o *Options) SetDstEncoding(e detect.CharEncoding) {
	o.to = e
}

//SetNewline is setting type of newline
func (o *Options) SetNewline(nl newline.Option) {
	o.nl = nl
}

//SrcEncoding returns character encoding of source text
func (o Options) SrcEncoding() detect.CharEncoding {
	return o.from
}

//DstEncoding returns character encoding of destination text
func (o Options) DstEncoding() detect.CharEncoding {
	return o.to
}

//Newline returns type of newline
func (o Options) Newline() newline.Option {
	return o.nl
}
