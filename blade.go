// blade engine factory

package blade

import (
	"github.com/fatrbaby/go-blade/fs"
	"os"
	"path"
)

const ViewExt = ".blade.html"

type Blade struct {
	Compiler     *Compiler
	loadViewPath string
	cachePath    string
}

func (blade *Blade) View(view string, data interface{}) *View {
	filename := path.Join(blade.loadViewPath, view) + ViewExt

	bytes, err := blade.Compiler.Compile(filename)

	if err != nil {
		panic(err)
	}

	engine := &View{
		HTML: string(bytes),
		Data: data,
	}

	err = engine.Prepare()

	if err != nil {
		panic(err)
	}

	return engine
}

func (blade *Blade) bootstrap() {
	blade.Compiler = NewCompiler()

	if has, _ := fs.Exists(blade.cachePath); has == false {
		os.MkdirAll(blade.cachePath, 0777)
	}

	compiledPath := path.Join(blade.cachePath, "views")

	if has, _ := fs.Exists(compiledPath); has == false {
		os.MkdirAll(compiledPath, 0777)
	}

	blade.Compiler.compiledFilePath = compiledPath
}
