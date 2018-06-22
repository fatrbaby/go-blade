package lexers

import "regexp"

const EndRangePattern = `@endrange`

var (
	endRangeMatcher  = regexp.MustCompile(EndRangePattern)
	endRangeReplacer = []byte("{{end}}")
)

type EndRange struct{}

func (ei *EndRange) Parse(context []byte) []byte {
	return endRangeMatcher.ReplaceAll(context, endRangeReplacer)
}
