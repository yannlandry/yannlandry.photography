package handler

import (
	"net/http"

	"github.com/yannlandry/yannlandry.photography/content"
)

type HomePresenter struct {
	Images []string
}

func Home(response http.ResponseWriter, request *http.Request) {
	content.Content.Home.Template.Execute(response, NewBasePresenter(&HomePresenter{
		Images: content.Content.Home.Images,
	}))
}
