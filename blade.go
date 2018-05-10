package blade

import (
	"os"
	"path"
)

const ViewExt = ".blade.html"

type Blade struct {
	Compiler     *Compiler
	loadViewPath string
	cachePath    string
}

func (blade *Blade) View(view string, data interface{}) ([]byte, error) {
	filename := path.Join(blade.loadViewPath, view) + ViewExt

	return blade.Compiler.Compile(filename)
}

func (blade *Blade) bootstrap() {
	blade.Compiler = NewCompiler()

	if has, _ := exists(blade.cachePath); has == false {
		os.MkdirAll(blade.cachePath, 0777)
	}

	compiledPath := path.Join(blade.cachePath, "views")

	if has, _ := exists(compiledPath); has == false {
		os.MkdirAll(compiledPath, 0777)
	}

	blade.Compiler.compiledFilePath = compiledPath
}
