package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yannlandry/yannlandry.photography/content"
)

func Page(response http.ResponseWriter, request *http.Request) {
	variables := mux.Vars(request)
	slug := variables["slug"]

	page, ok := content.Content.Pages.Pages[slug]
	if !ok {
		Error404(response, request)
		return
	}

	presenter := NewBasePresenter(page)
	presenter.WindowTitle = page.WindowTitle
	ExecuteTemplate(content.Content.Pages.Template, response, presenter)
}
