package blade

import (
	"fmt"
	"html/template"
	"io"
)

type View struct {
	Template *template.Template
	Name     string
	HTML     string
	Data     interface{}
}

func (view *View) Prepare() error {
	var err error
	view.Template, err = template.New("").Parse(view.HTML)

	return err
}

func (view *View) Render(writer io.Writer) error {
	if view.Name == "" {
		return view.Template.Execute(writer, view.Data)
	}

	return view.Template.ExecuteTemplate(writer, view.Name, view.Data)
}

func (view *View) Strings() string {
	return fmt.Sprintf("%s", view.HTML)
}
