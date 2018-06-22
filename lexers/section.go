package lexers

import "regexp"

const (
	StartSectionPattern   = `@section()`
	EndSectionPattern     = `@stop`
	SectionContentPattern = `@section\((.*)\)(.*)@stop`
)

type Section struct{}

func (section *Section) Parse(context []byte) []byte {
	return context
}

func (section *Section) BlockContent(context []byte) (string, []byte) {
	scMatcher := regexp.MustCompile(SectionContentPattern)

	matches := scMatcher.FindSubmatch(context)

	return string(matches[1]), matches[2]
}
