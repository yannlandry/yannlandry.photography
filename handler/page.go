package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func Page(response http.ResponseWriter, request *http.Request) {
	variables := mux.Vars(request)
	page := variables["page"]

	io.WriteString(response, fmt.Sprintf("Page: %s", page))
}
