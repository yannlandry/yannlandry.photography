package content

import (
	"fmt"

	"github.com/yannlandry/yannlandry.photography/util"
)

var (
	Content = NewWebsiteContent()
)

type WebsiteContent struct {
	Path       *util.Path
	Builder    *util.TemplateBuilder
	Navigation *NavigationContent
	Home       *HomeContent
	Blog       *BlogContent
	Pages      *PagesContent
}

func NewWebsiteContent() *WebsiteContent {
	return &WebsiteContent{
		Path:       nil,
		Builder:    nil,
		Navigation: NewNavigationContent(),
		Home:       NewHomeContent(),
		Blog:       NewBlogContent(),
		Pages:      nil,
	}
}

func (this *WebsiteContent) Load(path string) error {
	this.Path = util.NewPath(path)

	this.Builder = util.NewTemplateBuilder(this.Path.With("base.html"))

	if err := this.Navigation.Load(this.Path); err != nil {
		return fmt.Errorf("failed loading navigation: %s", err)
	}

	if err := this.Home.Load(this.Path, this.Builder); err != nil {
		return fmt.Errorf("failed loading home: %s", err)
	}

	if err := this.Blog.Load(this.Path, this.Builder); err != nil {
		return fmt.Errorf("failed loading blog: %s", err)
	}

	return nil
}
