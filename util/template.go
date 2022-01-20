package util

import (
	"fmt"
	"html/template"
	"path"
)

type TemplateBuilder struct {
	base string
}

func NewTemplateBuilder(base string) *TemplateBuilder {
	return &TemplateBuilder{
		base: base,
	}
}

func (this *TemplateBuilder) Load(paths ...string) (*template.Template, error) {
	arguments := []string{this.base}
	arguments = append(arguments, paths...)

	functions := template.FuncMap{
		"baseURL": func(extension string) string {
			return BaseURL.With(extension)
		},
		"staticURL": func(extension string) string {
			return StaticURL.With(extension)
		},
	}

	tmp, err := template.New(path.Base(this.base)).Funcs(functions).ParseFiles(arguments...)
	if err != nil {
		return nil, fmt.Errorf("failed loading templates `%s`: %s", paths, err)
	}

	return tmp, nil
}
