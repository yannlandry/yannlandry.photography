package content

import (
	"html/template"

	"github.com/yannlandry/yannlandry.photography/util"
)

type HomeContent struct {
	Template *template.Template
	Images   []string
}

func NewHomeContent() *HomeContent {
	return &HomeContent {
		Template: nil,
	}
}

func (this *HomeContent) Load(path *util.Path) error {
	var err error

	// Load home page template
	this.Template, err = util.LoadTemplate(path.With("home.html"))
	if err != nil {
		return err
	}

	// Load home page slideshow images
	err = util.LoadYAML(path.With("home.yaml"), &this.Images)
	if err != nil {
		return err
	}

	return nil
}
