package blade

import (
	"fmt"
	"io"
	"html/template"
)

type View struct {
	Name string
	Source []byte
	Data interface{}
}

func (view View) Render(writer io.Writer) error {
	engine, err := template.New("view").Parse(string(view.Source))

	if err != nil {
		return err
	}

	if view.Name == "" {
		return engine.Execute(writer, view.Data)
	}

	return engine.ExecuteTemplate(writer, view.Name, view.Data)
}

func (view View) Strings() string {
	return fmt.Sprintf("%s", view.Source)
}
