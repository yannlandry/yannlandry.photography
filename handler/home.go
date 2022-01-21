package handler

import (
	"net/http"

	"github.com/yannlandry/yannlandry.photography/content"
)

type HomePresenter struct {
}

func Home(response http.ResponseWriter, request *http.Request) {
	presenter := NewBasePresenter(&HomePresenter{})
	presenter.WindowTitle = "yannlandry.photography"
	ExecuteTemplate(content.Content.Home.Template, response, presenter)
}
