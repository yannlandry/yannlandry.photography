package content

import (
	"time"
)

type BlogPost struct {
	Title       string    `yalm:"title"`        // title to display in the page, supports markdown
	WindowTitle string    `yaml:"window_title"` // if empty, use `title` with stripped HTML tags
	Slug        string    `yaml:"slug"`
	Date        time.Time `yaml:"date"`
	Image       string    `yaml:"image"`
	Keywords    []string  `yaml:"keywords"`
	Path        string    `yaml:"path"`
	Content     string    // markdown-formatted content from `path`
}

type BlogContent struct {
	Posts    []*BlogPost            // ordered list of posts
	Slugs    map[string]*BlogPost   // slug -> blog post for individual display
	Keywords map[string][]*BlogPost // keyword -> ordered list of posts
}
