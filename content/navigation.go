package content

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type NavigationLink struct {
	Title string `yaml:"title"`
	Link  string `yaml:"link"`
}

type NavigationContent struct {
	Links []*NavigationLink
}

func NewNavigationContent() *NavigationContent {
	return &NavigationContent {}
}

func (this *NavigationContent) Load(path string) error {
	// Read YAML file
	path = filepath.Join(path, "navigation.yaml")
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error loading `%s`: %s", path, err)
	}

	// Parse YAML
	err = yaml.Unmarshal(content, &this.Links)
	if err != nil {
		return fmt.Errorf("error parsing YAML: %s", err)
	}

	return nil
}
