package blade

import (
	"github.com/fatrbaby/go-blade/fs"
	"github.com/fatrbaby/go-blade/lexers"
)

type Compiler struct {
	lexers []Lexer
}

func (compiler *Compiler) Compile(file string) ([]byte, error) {
	bytes, err := fs.Load(file)

	if err != nil {
		return []byte(""), err
	}

	for _, lexer := range compiler.lexers {
		bytes = lexer.Parse(bytes)
	}

	return bytes, nil
}

func (compiler *Compiler) applyLexers() {
	compiler.lexers = []Lexer{
		new(lexers.Echo),
		new(lexers.If),
		new(lexers.Else),
		new(lexers.EndIf),
		new(lexers.Range),
		new(lexers.EndRange),
	}
}

func NewCompiler() *Compiler {
	compiler := new(Compiler)
	compiler.applyLexers()

	return compiler
}
