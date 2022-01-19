package handler

import (
	"net/http"

	"github.com/yannlandry/yannlandry.photography/content"
)

type HomePresenter struct {
	Navigation *content.NavigationContent
	Images     []string
}

func Home(response http.ResponseWriter, request *http.Request) {
	content.Content.Home.Template.Execute(response, HomePresenter {
		Navigation: content.Content.Navigation,
		Images:     content.Content.Home.Images,
	})
}
