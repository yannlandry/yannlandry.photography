package content

import (
	"fmt"
)

var (
	Content = NewWebsiteContent()
)

type WebsiteContent struct {
	Path       string
	Navigation *NavigationContent
	Home       *HomeContent
	Blog       *BlogContent
	Pages      *PagesContent
}

func NewWebsiteContent() *WebsiteContent {
	return &WebsiteContent {
		"",
		NewNavigationContent(),
		NewHomeContent(),
		nil,
		nil,
	}
}

func (this *WebsiteContent) Load(path string) error {
	this.Path = path

	if err := this.Navigation.Load(this.Path); err != nil {
		return fmt.Errorf("error loading navigation: %s", err)
	}

	if err := this.Home.Load(this.Path); err != nil {
		return fmt.Errorf("error loading home: %s", err)
	}

	return nil
}
