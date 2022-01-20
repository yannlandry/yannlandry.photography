package handler

import (
	"fmt"
	"net/http"

	"github.com/yannlandry/yannlandry.photography/content"
)

type HomePresenter struct {
	Images []string
}

func Home(response http.ResponseWriter, request *http.Request) {
	presenter := NewBasePresenter(&HomePresenter{
		Images: content.Content.Home.Images,
	})
	presenter.WindowTitle = "yannlandry.photography"
	if err := content.Content.Home.Template.Execute(response, presenter); err != nil {
		fmt.Printf("Error executing template: %s\n", err)
	}
}
