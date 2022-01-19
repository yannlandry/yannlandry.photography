package content

import (
	"github.com/yannlandry/yannlandry.photography/util"
)

type NavigationLink struct {
	Title string `yaml:"Title"`
	Link  string `yaml:"Link"`
}

type NavigationContent struct {
	Links []*NavigationLink
}

func NewNavigationContent() *NavigationContent {
	return &NavigationContent {}
}

func (this *NavigationContent) Load(path *util.Path) error {
	err := util.LoadYAML(path.With("navigation.yaml"), &this.Links)
	if err != nil {
		return err
	}

	// Prepend base URL
	for _, link := range this.Links {
		link.Link = util.BaseURL.With(link.Link)
	}

	return nil
}
