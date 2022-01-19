package handler

import (
	"github.com/yannlandry/yannlandry.photography/content"
	"github.com/yannlandry/yannlandry.photography/util"
)

type BasePresenter struct {
	WindowTitle string
	Stylesheets []string
	Home        string
	Navigation  *content.NavigationContent
	Content     interface{}
}

func NewBasePresenter(presenter interface{}) *BasePresenter {
	return &BasePresenter{
		Home:       util.BaseURL.With(""),
		Navigation: content.Content.Navigation,
		Content:    presenter,
	}
}
