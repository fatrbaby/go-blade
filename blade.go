package blade

import (
	"path"
	"errors"
	"os"
	"io/ioutil"
)

type Blade struct {
	Compiler  *Compiler
	viewPath  string
	cachePath string
	compiledFilePath string
}

func (blade *Blade) View(view string, data interface{}) (string, error) {
	return blade.Compiler.Compile()
}

func (blade *Blade) bootstrap() {
	blade.compiledFilePath = path.Join(blade.cachePath, "views")
	blade.Compiler = new(Compiler)
}

func (blade *Blade)load(view string) ([]byte, error) {
	filename := path.Join(blade.viewPath, view)

	has, err := exists(filename)

	if err != nil {
		return []byte(""), err
	}

	if !has {
		return []byte(""), errors.New("file does not exists")
	}

	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		return []byte(""), err
	}

	bytes, err := ioutil.ReadAll(file)

	if err != nil {
		return []byte(""), err
	}

	return bytes, nil
}