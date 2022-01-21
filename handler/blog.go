package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yannlandry/yannlandry.photography/content"
)

type BlogPresenter struct {
	Posts         []*content.BlogPost
	Keywords      map[string][]*content.BlogPost
	ActiveKeyword string
}

func Blog(response http.ResponseWriter, request *http.Request) {
	presenter := NewBasePresenter(&BlogPresenter{
		Posts:    content.Content.Blog.Posts,
		Keywords: content.Content.Blog.Keywords,
	})
	presenter.WindowTitle = "Blog"
	ExecuteTemplate(content.Content.Blog.Template, response, presenter)
}

func BlogPost(response http.ResponseWriter, request *http.Request) {
	variables := mux.Vars(request)
	slug := variables["slug"]

	post, ok := content.Content.Blog.Slugs[slug]
	if !ok {
		Error404(response, request)
		return
	}

	presenter := NewBasePresenter(post)
	presenter.WindowTitle = post.WindowTitle
	ExecuteTemplate(content.Content.Blog.TemplatePost, response, presenter)
}

func BlogKeyword(response http.ResponseWriter, request *http.Request) {
	variables := mux.Vars(request)
	keyword := variables["keyword"]

	posts, ok := content.Content.Blog.Keywords[keyword]
	if !ok {
		Error404(response, request)
		return
	}

	presenter := NewBasePresenter(&BlogPresenter{
		Posts:         posts,
		Keywords:      content.Content.Blog.Keywords,
		ActiveKeyword: keyword,
	})
	presenter.WindowTitle = "Blog: " + keyword
	ExecuteTemplate(content.Content.Blog.Template, response, presenter)
}
