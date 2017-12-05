package normalize

import "io"

//Do returns normalized text
func Do(txt io.Reader, opt Option) io.Reader {
	if opt == Unknown {
		return txt
	}
	return opt.GetForm().Reader(txt)
}
