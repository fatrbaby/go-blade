package lexers

import "regexp"

const IfPattern = `@if\((.*)\)`

var (
	ifMatcher  = regexp.MustCompile(IfPattern)
	ifReplacer = []byte(`{{if ${1}}}`)
)

type If struct{}

func (i *If) Parse(context []byte) []byte {
	return ifMatcher.ReplaceAll(context, ifReplacer)
}
