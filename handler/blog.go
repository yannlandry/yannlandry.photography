package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yannlandry/yannlandry.photography/content"
)

type BlogPresenter struct {
	Posts []*content.BlogPost
	Keywords map[string][]*content.BlogPost
	ActiveKeyword string
}

func Blog(response http.ResponseWriter, request *http.Request) {
	content.Content.Blog.Template.Execute(response, BlogPresenter {
		Posts: content.Content.Blog.Posts,
		Keywords: content.Content.Blog.Keywords,
	})
}

type BlogPostPresenter struct {
	Post *content.BlogPost
}

func BlogPost(response http.ResponseWriter, request *http.Request) {
	variables := mux.Vars(request)
	slug := variables["slug"]

	post, ok := content.Content.Blog.Slugs[slug]
	if !ok {
		io.WriteString(response, fmt.Sprintf("404 Blog Post: %s", slug))
		return
	}

	content.Content.Blog.TemplatePost.Execute(response, BlogPostPresenter {
		Post: post,
	})
}

func BlogKeyword(response http.ResponseWriter, request *http.Request) {
	variables := mux.Vars(request)
	keyword := variables["keyword"]

	posts, ok := content.Content.Blog.Keywords[keyword]
	if !ok {
		io.WriteString(response, fmt.Sprintf("404 Blog Keyword: %s", keyword))
		return
	}

	content.Content.Blog.Template.Execute(response, BlogPresenter {
		Posts: posts,
		Keywords: content.Content.Blog.Keywords,
		ActiveKeyword: keyword,
	})
}
