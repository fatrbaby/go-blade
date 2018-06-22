package blade

import (
	"crypto/sha1"
	"fmt"
	"github.com/fatrbaby/go-blade/lexers"
	"io/ioutil"
	"os"
	"github.com/fatrbaby/go-blade/fs"
)

var (
	hasher = sha1.New()
)

type Compiler struct {
	compiledFilePath string
	lexers           []Lexer
}

func (compiler *Compiler) Compile(file string) ([]byte, error) {
	if compiler.IsExpired(file) {
		bytes, _ := fs.Load(file)

		for _, lexer := range compiler.lexers {
			bytes = lexer.Parse(bytes)
		}

		// doing compile
		return bytes, compiler.WriteCompiled(compiler.CompiledPath(file), bytes)
	}

	compiled := compiler.CompiledPath(file)

	return fs.Load(compiled)
}

func (compiler *Compiler) CompiledPath(file string) string {
	hasher.Reset()
	hasher.Write([]byte(file))
	hashed := fmt.Sprintf("%x", hasher.Sum(nil))

	return compiler.compiledFilePath + string(os.PathSeparator) + hashed + ".blade.html"
}

func (compiler *Compiler) IsExpired(file string) bool {
	compiled := compiler.CompiledPath(file)

	has, _ := fs.Exists(compiled)

	if !has {
		return true
	}

	fileLastModified := fs.LastModified(file)
	compiledLastModified := fs.LastModified(compiled)

	return fileLastModified.After(compiledLastModified) || fileLastModified.Equal(compiledLastModified)
}

func (compiler *Compiler) WriteCompiled(filename string, contents []byte) error {
	return ioutil.WriteFile(filename, contents, 0644)
}

func (compiler *Compiler) applyLexers() {
	compiler.lexers = []Lexer{
		new(lexers.Echo),
		new(lexers.If),
		new(lexers.Else),
		new(lexers.EndIf),
	}
}

func NewCompiler() *Compiler {
	compiler := new(Compiler)
	compiler.applyLexers()

	return compiler
}
