package lexers

import "regexp"

const EndIfPattern = `@endif`

var (
	endIfMatcher  = regexp.MustCompile(EndIfPattern)
	endIfReplacer = []byte("{{end}}")
)

type EndIf struct{}

func (ei *EndIf) Parse(context []byte) []byte {
	return endIfMatcher.ReplaceAll(context, endIfReplacer)
}
