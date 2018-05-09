package blade

type Lexer interface {
	Parse(context []byte) []byte
}
