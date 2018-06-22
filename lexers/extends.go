package lexers

import "regexp"

const ExtendsPattern = `@extends\((.*)\)`
const YieldPattern = `@yield\((.*)\)`

var (
	extendsMatcher = regexp.MustCompile(ExtendsPattern)
)

type Extends struct {
	parents [][]byte
}

func (ex *Extends) Parse(context []byte) []byte {
	return context
}
