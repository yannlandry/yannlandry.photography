package util

import (
	"fmt"
	"html/template"
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
	tmp, err := template.ParseFiles(arguments...)
	if err != nil {
		return nil, fmt.Errorf("failed loading templates `%s`: %s", paths, err)
	}
	return tmp, nil
}
