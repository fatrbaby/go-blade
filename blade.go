// blade engine factory

package blade

import (
	"github.com/fatrbaby/go-blade/cache"
	"path"
)

const ViewExt = ".blade.html"

type Blade struct {
	Debug    bool
	Compiler *Compiler
	ViewPath string
	cache    cache.Driver
}

func (blade *Blade) View(view string, data interface{}) *View {
	filename := path.Join(blade.ViewPath, view) + ViewExt
	hashedKey := HashKey(filename)
	var err error

	compiled := blade.cache.Get(hashedKey)

	if compiled != "" {
		return build(compiled, data)
	}

	var bytes []byte
	bytes, err = blade.Compiler.Compile(filename)
	compiled = string(bytes)
	blade.cache.Set(hashedKey, compiled, 1000)

	if err != nil {
		panic(err)
	}

	return build(compiled, data)
}

func (blade *Blade) bootstrap() {
	blade.Compiler = NewCompiler()
}

func build(compiled string, data interface{}) *View {
	engine := &View{
		HTML: compiled,
		Data: data,
	}

	err := engine.Prepare()

	if err != nil {
		panic(err)
	}

	return engine
}
