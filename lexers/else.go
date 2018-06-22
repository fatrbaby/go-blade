// else lexer
// use as if else
package lexers

import "regexp"

const ElsePattern = `@else`

var (
	elseMatcher  = regexp.MustCompile(ElsePattern)
	elseReplacer = []byte("{{else}}")
)

type Else struct{}

func (es *Else) Parse(context []byte) []byte {
	return elseMatcher.ReplaceAll(context, elseReplacer)
}
