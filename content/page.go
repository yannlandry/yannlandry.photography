package content

type Page struct {
	WindowTitle string `yaml:"WindowTitle"`
	Slug        string `yaml:"Slug"`
	Path        string `yaml:"Path"`
	Content     string
}

type PagesContent struct {
	Pages map[string]Page
}
