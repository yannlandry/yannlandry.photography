package content

import (
	"html/template"

	"github.com/yannlandry/yannlandry.photography/util"
)

type ErrorContent struct {
	Template *template.Template
}

func NewErrorContent() *ErrorContent {
	return &ErrorContent{}
}

func (this *ErrorContent) Load(path *util.Path, builder *util.TemplateBuilder) error {
	var err error

	this.Template, err = builder.Load(path.With("error.html"))
	if err != nil {
		return err
	}

	return nil
}
