package lexers

import "regexp"

const RangePattern = `@range\((.*)\)`

var (
	rangeMatcher  = regexp.MustCompile(RangePattern)
	rangeReplacer = []byte(`{{range ${1}}}`)
)

type Range struct{}

func (r Range) Parse(context []byte) []byte {
	return rangeMatcher.ReplaceAll(context, rangeReplacer)
}
