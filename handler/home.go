package handler

import (
	"net/http"

	"github.com/yannlandry/yannlandry.photography/content"
)

type HomePresenter struct {
}

func Home(response http.ResponseWriter, request *http.Request) {
	presenter := NewBasePresenter(&HomePresenter{})
	ExecuteTemplate(content.Content.Home.Template, response, presenter)
}
