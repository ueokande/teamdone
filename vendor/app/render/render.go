package render

import (
	"html/template"
	"io"
)

var DefaultRenderer Renderer

type Renderer interface {
	Render(io.Writer, string, interface{}) error
}

func InitTemplateRenderer(t *template.Template) {
	DefaultRenderer = &TemplateRenderer{t}
}

func Render(w io.Writer, name string, data interface{}) error {
	return DefaultRenderer.Render(w, name, data)
}

type TemplateRenderer struct {
	template *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}) error {
	return t.template.ExecuteTemplate(w, name, data)
}
