package lexers

import "regexp"

const EchoPattern = `{{(.*)}}`

var (
	echoMatcher  = regexp.MustCompile(EchoPattern)
	echoReplacer = []byte(`{{${1}}}`)
)

type Echo struct{}

func (e *Echo) Parse(context []byte) []byte {
	return echoMatcher.ReplaceAll(context, echoReplacer)
}
