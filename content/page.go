package content

type Page struct {
	WindowTitle string `yaml:"window_title"`
	Slug        string `yaml:"slug"`
	Path        string `yaml:"path"`
	Content     string
}

type PagesContent struct {
	Pages map[string]Page
}
