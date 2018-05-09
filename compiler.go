package blade

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	hasher = sha1.New()
)

type Compiler struct {
	compiledFilePath string
}

func (compiler *Compiler) Compile(file string) ([]byte, error) {
	if compiler.IsExpired(file) {
		bytes, _ := load(file)
		// doing compile
		return bytes, compiler.WriteCompiled(compiler.CompiledPath(file), bytes)
	}

	compiled := compiler.CompiledPath(file)

	return load(compiled)
}

func (compiler *Compiler) CompiledPath(file string) string {
	hasher.Reset()
	hasher.Write([]byte(file))
	hashed := fmt.Sprintf("%x", hasher.Sum(nil))

	return compiler.compiledFilePath + string(os.PathSeparator) + hashed + ".blade.html"
}

func (compiler *Compiler) IsExpired(file string) bool {
	compiled := compiler.CompiledPath(file)

	has, _ := exists(compiled)

	if !has {
		return true
	}

	fileLastModified := lastModified(file)
	compiledLastModified := lastModified(compiled)

	return fileLastModified.After(compiledLastModified) || fileLastModified.Equal(compiledLastModified)
}

func (compiler *Compiler)WriteCompiled(filename string, contents []byte) error {
	return ioutil.WriteFile(filename, contents, 0644)
}
