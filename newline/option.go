package newline

import "strings"

//Option is type of newline
type Option int

const (
	//Unknown is unknown type
	Unknown Option = iota
	//LF is LF only
	LF
	//CR is CR only
	CR
	//CRLF is CR+LF
	CRLF
)

var newlineMap = map[string]Option{
	"lf":   LF,
	"cr":   CR,
	"crlf": CRLF,
}

//TypeofNewline returns type of newline
func TypeofNewline(s string) Option {
	if e, ok := newlineMap[strings.ToLower(s)]; ok {
		return e
	}
	return Unknown
}

func (nl Option) String() string {
	for key, value := range newlineMap {
		if value == nl {
			return key
		}
	}
	return "Unknown"
}
