package content

import (
	"fmt"
	"html/template"
	"path/filepath"
)

type HomeContent struct {
	Template *template.Template
}

func NewHomeContent() *HomeContent {
	return &HomeContent {
		nil,
	}
}

func (this *HomeContent) Load(path string) error {
	path = filepath.Join(path, "home.html")

	var err error
	this.Template, err = template.ParseFiles(path)
	if err != nil {
		return fmt.Errorf("failed loading the template `%s`: %s", path, err)
	}

	return nil
}
