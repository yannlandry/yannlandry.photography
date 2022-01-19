package handler

import (
	"github.com/yannlandry/yannlandry.photography/content"
)

type BasePresenter struct {
	WindowTitle string
	Stylesheets []string
	Navigation *content.NavigationContent
	Content interface{}
}

func NewBasePresenter(presenter interface{}) *BasePresenter {
	return &BasePresenter{
		Navigation: content.Content.Navigation,
		Content: presenter,
	}
}
