package util

import (
	"fmt"
	"html/template"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func LoadFile(path string) ([]byte, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed loading `%s`: %s", path, err)
	}
	return content, nil
}

func LoadYAML(path string, destination interface{}) error {
	content, err := LoadFile(path)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(content, destination); err != nil {
		return fmt.Errorf("failed parsing `%s`: %s", path, err)
	}
	return nil
}

func LoadTemplate(path string) (*template.Template, error) {
	tmp, err := template.ParseFiles(path)
	if err != nil {
		return nil, fmt.Errorf("failed loading template `%s`: %s", path, err)
	}
	return tmp, nil
}
