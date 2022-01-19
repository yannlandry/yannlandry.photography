package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yannlandry/yannlandry.photography/content"
)

func Page(response http.ResponseWriter, request *http.Request) {
	variables := mux.Vars(request)
	slug := variables["slug"]

	page, ok := content.Content.Pages.Pages[slug]
	if !ok {
		io.WriteString(response, fmt.Sprintf("404 Page: %s", slug))
		return
	}

	presenter := NewBasePresenter(page)
	presenter.WindowTitle = page.WindowTitle
	content.Content.Pages.Template.Execute(response, presenter)
}
