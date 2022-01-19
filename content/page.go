package content

import (
	"html/template"

	"github.com/yannlandry/yannlandry.photography/util"
)

type Page struct {
	WindowTitle string `yaml:"WindowTitle"`
	Slug        string `yaml:"Slug"`
	Path        string `yaml:"Path"`
	Content     template.HTML
}

type PagesContent struct {
	Template *template.Template
	Pages map[string]*Page
}

func NewPagesContent() *PagesContent {
	return &PagesContent{
		Template: nil,
		Pages: map[string]*Page{},
	}
}

func (this *PagesContent) Load(path *util.Path, builder *util.TemplateBuilder) error {
	path = util.NewPath(path.With("pages"))
	var err error

	this.Template, err = builder.Load(path.With("page.html"))
	if err != nil {
		return err
	}

	pages := []*Page{}
	err = util.LoadYAML(path.With("pages.yaml"), &pages)
	if err != nil {
		return err
	}
	for _, page := range pages {
		if err := page.LoadContent(path); err != nil {
			return err
		}
		this.Pages[page.Slug] = page
	}

	return nil
}

func (this *Page) LoadContent(path *util.Path) error {
	content, err := util.LoadFile(path.With(this.Path))
	if err != nil {
		return err
	}
	this.Content = template.HTML(util.Markdown.Render(content))
	return nil
}
