package handler

import (
	"net/http"

	"github.com/yannlandry/yannlandry.photography/content"
)

type HomePresenter struct {
	Navigation *content.NavigationContent
}

func Home(response http.ResponseWriter, request *http.Request) {
	content.Content.Home.Template.Execute(response, HomePresenter {
		Navigation: content.Content.Navigation,
	})
}
