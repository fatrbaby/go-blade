package blade

import (
	"crypto/sha1"
	"fmt"
	"os"
)

var (
	hasher = sha1.New()
)

type Compiler struct {
	compiledFilePath string
}

func (compiler *Compiler) Compile() (strint, error) {
	return "", nil
}

func (compiler *Compiler)CompiledPath(file string) string  {
	hasher.Reset()
	hasher.Write([]byte(file))
	hashed := fmt.Sprintf("%x", hasher.Sum(nil))

	return compiler.compiledFilePath + string(os.PathSeparator) + hashed + ".blade.html"
}

func (compiler *Compiler)IsExpired(file string) bool  {
	compiled := compiler.CompiledPath(file)

	has, _ := exists(compiled)

	if !has {
		return true
	}

	return lastModified(file).After(lastModified(compiled))
}
