package blade

import (
	"fmt"
	"io"
	"html/template"
)

type View struct {
	Template *template.Template
	Name string
	Source []byte
	Data interface{}
}

func (view View)Parse() error {
	var err error
	view.Template, err = template.New(view.Name).Parse(string(view.Source))

	return err
}

func (view View) Render(writer io.Writer) error {
	if view.Name == "" {
		return view.Template.Execute(writer, view.Data)
	}

	return view.Template.ExecuteTemplate(writer, view.Name, view.Data)
}

func (view View) Strings() string {
	return fmt.Sprintf("%s", view.Source)
}
