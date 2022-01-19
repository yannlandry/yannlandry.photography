package content

import (
	"html/template"
	"time"

	"github.com/yannlandry/yannlandry.photography/util"
)

type BlogPost struct {
	Title       string    `yaml:"Title"`       // title to display in the page, supports markdown
	WindowTitle string    `yaml:"WindowTitle"` // if empty, use `title` with stripped HTML tags
	Slug        string    `yaml:"Slug"`
	Date        time.Time `yaml:"Date"`
	Image       string    `yaml:"Image"`
	Keywords    []string  `yaml:"Keywords"`
	Path        string    `yaml:"Path"`
	Content     string    // markdown-formatted content from `path`
	Summary     string    `yaml:"Summary"` // if empty, auto-generated from `Content`
}

type BlogContent struct {
	Template        *template.Template
	TemplatePost    *template.Template
	TemplateKeyword *template.Template
	Posts           []*BlogPost            // ordered list of posts
	Slugs           map[string]*BlogPost   // slug -> blog post for individual display
	Keywords        map[string][]*BlogPost // keyword -> ordered list of posts
}

func NewBlogContent() *BlogContent {
	return &BlogContent{
		Template:        nil,
		TemplatePost:    nil,
		TemplateKeyword: nil,
		Posts:           []*BlogPost{},
		Slugs:           map[string]*BlogPost{},
		Keywords:        map[string][]*BlogPost{},
	}
}

func (this *BlogContent) Load(path *util.Path) error {
	var err error

	this.Template, err = util.LoadTemplate(path.With("blog.html"))
	if err != nil {
		return err
	}

	this.TemplatePost, err = util.LoadTemplate(path.With("blogpost.html"))
	if err != nil {
		return err
	}

	err = util.LoadYAML(path.With("blog.yaml"), &this.Posts)
	if err != nil {
		return err
	}

	for _, post := range this.Posts {
		if err := loadContent(path, post); err != nil {
			return err
		}
		this.addToSlugs(post)
		this.addToKeywords(post)
	}

	return nil
}

func loadContent(path *util.Path, post *BlogPost) error {
	content, err := util.LoadFile(path.With(post.Path))
	if err != nil {
		return err
	}
	post.Content = string(content)
	return nil
}

func (this *BlogContent) addToSlugs(post *BlogPost) {
	this.Slugs[post.Slug] = post
}

func (this *BlogContent) addToKeywords(post *BlogPost) {
	for _, keyword := range post.Keywords {
		if _, ok := this.Keywords[keyword]; !ok {
			this.Keywords[keyword] = []*BlogPost{}
		}
		this.Keywords[keyword] = append(this.Keywords[keyword], post)
	}
}
