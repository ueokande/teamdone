package render

import (
	"html/template"
	"io"
)

type Renderer interface {
	Render(io.Writer, string, interface{}) error
}

type TemplateRenderer struct {
	Template *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}) error {
	return t.Template.ExecuteTemplate(w, name, data)
}
